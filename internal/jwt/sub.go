package jwt

import (
	"fmt"
	"time"

	"github.com/lestrrat-go/jwx/jwt"
)

type Subject struct {
	UID     string `json:"uid"`
	Expires int64  `json:"exp"`
}

func (s *Subject) updateExpiry(exp int64) {
	s.Expires = time.Now().Unix() + exp
}

func getTokenSubject(token jwt.Token) (*Subject, error) {
	const (
		uidParamName = "uid"
	)

	uid, ok := token.Get(uidParamName)
	if !ok {
		return nil, fmt.Errorf("no %d param", uidParamName)
	}

	uidString, ok := uid.(string)
	if !ok {
		return nil, fmt.Errorf("invalid type of %s: %t", uidParamName, uid)
	}

	return &Subject{
		UID:     uidString,
		Expires: token.Expiration().Unix(),
	}, nil
}
