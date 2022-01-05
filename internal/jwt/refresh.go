package jwt

import (
	"github.com/lestrrat-go/jwx/jwk"
	"github.com/lestrrat-go/jwx/jwt"
)

func (m *ManagerImpl) Refresh(accessToken string) (string, error) {
	jwks := jwk.NewSet()
	jwks.Add(m.publicKey)

	token, err := jwt.Parse([]byte(accessToken), jwt.WithKeySet(jwks))
	if err != nil {
		return "", err
	}

	sub, err := getTokenSubject(token)
	if err != nil {
		return "", err
	}

	return m.Generate(sub)
}
