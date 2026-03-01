package endpoint

import (
	"context"

	"github.com/go-kit/kit/endpoint"

	"wsu-senior-project/service"
)

// Endpoints holds all Go kit endpoints for the service
type Endpoints struct {
	SayHelloEndpoint   endpoint.Endpoint
	CreateUserEndpoint endpoint.Endpoint
	GetUsersEndpoint   endpoint.Endpoint
}

// MakeEndpoints creates and returns all endpoints for the service
func MakeEndpoints(svc service.Service) Endpoints {
	return Endpoints{
		SayHelloEndpoint:   makeSayHelloEndpoint(svc),
		CreateUserEndpoint: makeCreateUserEndpoint(svc),
		GetUsersEndpoint:   makeGetUsersEndpoint(svc),
	}
}

// HelloRequest holds the request parameters for the SayHello endpoint
type HelloRequest struct {
	Name string `json:"name"`
}

// HelloResponse holds the response from the SayHello endpoint
type HelloResponse struct {
	Message string `json:"message"`
}

// CreateUserRequest holds the request parameters for the CreateUser endpoint
type CreateUserRequest struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

// CreateUserResponse holds the response from the CreateUser endpoint
type CreateUserResponse struct {
	User  *service.User `json:"user,omitempty"`
	Error string        `json:"error,omitempty"`
}

// GetUsersResponse holds the response from the GetUsers endpoint
type GetUsersResponse struct {
	Users []service.User `json:"users"`
	Error string         `json:"error,omitempty"`
}

func makeSayHelloEndpoint(svc service.Service) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(HelloRequest)
		msg := svc.SayHello(req.Name)
		return HelloResponse{Message: msg}, nil
	}
}

func makeCreateUserEndpoint(svc service.Service) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(CreateUserRequest)
		user, err := svc.CreateUser(req.Name, req.Email)
		if err != nil {
			return CreateUserResponse{Error: err.Error()}, nil
		}
		return CreateUserResponse{User: user}, nil
	}
}

func makeGetUsersEndpoint(svc service.Service) endpoint.Endpoint {
	return func(_ context.Context, _ interface{}) (interface{}, error) {
		users, err := svc.GetUsers()
		if err != nil {
			return GetUsersResponse{Error: err.Error()}, nil
		}
		return GetUsersResponse{Users: users}, nil
	}
}
