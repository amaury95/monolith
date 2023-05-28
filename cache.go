package monolith

import (
	"context"
	"time"

	"google.golang.org/protobuf/reflect/protoreflect"
)

//go:generate mockgen --destination=./mocks/cache.go --package=mocks github.com/amaury95/monolith ICache
type ICache interface {
	Get(ctx context.Context, key string, out protoreflect.ProtoMessage) error
	Set(ctx context.Context, key string, val protoreflect.ProtoMessage, expiration time.Duration) error
}
