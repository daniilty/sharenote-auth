package jwt

import (
	"encoding/json"

	"github.com/lestrrat-go/jwx/jws"
)

func (m *ManagerImpl) Generate(sub *Subject) (string, error) {
	sub.updateExpiry(m.tokenExp)

	bb, err := json.Marshal(sub)
	if err != nil {
		return "", err
	}

	accessToken, err := jws.Sign(bb, m.alg, m.privateKey)

	return string(accessToken), err
}
