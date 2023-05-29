package monolith

import (
	"context"

	"github.com/jmoiron/sqlx"
)

type IConnection[T any] interface {
	GetConnection(ctx context.Context) (*T, bool)
}

//go:generate mockgen --destination=./mocks/connection.go --package=mocks github.com/amaury95/monolith Connection
type Connection = IConnection[sqlx.DB]