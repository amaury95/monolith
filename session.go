package monolith

import (
	"github.com/golang-jwt/jwt/v4"
)

//go:generate mockgen --destination=./mocks/session.go --package=mocks github.com/amaury95/monolith ISession
type ISession interface {
	Claims(accessToken string) (*jwt.RegisteredClaims, error)
}

//go:generate mockgen --destination=./mocks/authorizer.go --package=mocks github.com/amaury95/monolith IAuthorizer
type IAuthorizer interface {
	Authorize(actor interface{}, action interface{}, resource interface{}) error
}
