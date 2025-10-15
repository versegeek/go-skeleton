package handler

import (
	"context"

	authv1 "github.com/versegeek/verse-proto-go/gen/auth/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (h *handler) CreateClient(ctx context.Context, req *authv1.CreateClientRequest) (*authv1.CreateClientResponse, error) {
	return nil, status.Error(codes.Unimplemented, "method CreateClient not implemented")
}
