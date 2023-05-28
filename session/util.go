package session

import (
	"context"
	"fmt"

	"google.golang.org/grpc/metadata"
)

func GetAccessToken(ctx context.Context) (token string, err error) {
	data, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return token, fmt.Errorf("metadata not present in context")
	}
	tokens := data.Get("Authorization")
	if len(tokens) < 1 {
		return token, fmt.Errorf("token not found")
	}
	return tokens[0], nil
}
