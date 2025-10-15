package service

import (
	"context"

	authv1 "github.com/versegeek/verse-proto-go/gen/auth/v1"
)

var _ IService = (*service)(nil)

type (
	IService interface {
		Resource() Resource
	}

	service struct {
		clientSDK authv1.ClientAPIClient
		resource  Resource
	}

	Resource struct{}
)

func (s *service) Resource() Resource {
	return s.resource
}

func New() IService {
	return &service{}
}

func (s *service) Test() {
	client := &authv1.Client{
		Id: "",
	}
	req := &authv1.GetClientByIDRequest{Id: client.Id}
	_, err := s.clientSDK.GetClientByID(context.Background(), req)
	if err != nil {
		panic(err)
	}

}
