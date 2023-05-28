package monolith

import (
	"context"

	"github.com/jmoiron/sqlx"
)

//go:generate mockgen --destination=./mocks/connection.go --package=mocks github.com/amaury95/monolith IConnection
type IConnection interface {
	GetConnection(ctx context.Context) (*sqlx.DB, bool)
}
