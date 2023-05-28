package session

import (
	"context"
	"fmt"

	"github.com/golang-jwt/jwt/v4"
)

type IConfig interface {
	TokenVerifierPublicKey(context.Context) (string, error)
}

type session struct {
	cnf IConfig
	ctx context.Context
}

func NewSession(ctx context.Context, cnf IConfig) *session {
	return &session{ctx: ctx, cnf: cnf}
}

func (s *session) Claims(accessToken string) (*jwt.RegisteredClaims, error) {
	publicKey, err := s.cnf.TokenVerifierPublicKey(s.ctx)
	if err != nil {
		return nil, fmt.Errorf("TokenVerifierPublicKey: %w", err)
	}

	token, err := jwt.ParseWithClaims(accessToken, &jwt.RegisteredClaims{}, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodRSA)
		if !ok {
			return nil, fmt.Errorf("unexpected token signing method")
		}
		return jwt.ParseRSAPublicKeyFromPEM([]byte(publicKey))
	})
	if err != nil {
		return nil, fmt.Errorf("error parsing token: %v", err)
	}
	if !token.Valid || token.Claims.Valid() != nil {
		return nil, fmt.Errorf("invalid token")
	}
	claims, ok := token.Claims.(*jwt.RegisteredClaims)
	if !ok {
		return nil, fmt.Errorf("invalid token claims")
	}
	return claims, nil
}
