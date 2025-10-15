package handler

import (
	authServerv1 "github.com/versegeek/go-skeleton/api/oauth2"
	"github.com/versegeek/go-skeleton/internal/service"
	authv1 "github.com/versegeek/verse-proto-go/gen/auth/v1"
)

// var _ Handler = (*handler)(nil)

type (
	Handler interface {
		OAuth2Handler
		VersionHandler
		// authv1.ClientAPIServer
		// authServerv1.ServerInterface
	}

	handler struct {
		service service.IService
		authv1.UnimplementedClientAPIServer
		authServerv1.ServerInterfaceWrapper
	}
)

func New(service service.IService) Handler {
	return &handler{
		service: service,
	}
}
