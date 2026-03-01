package endpoint

import (
	"context"

	"github.com/go-kit/kit/endpoint"

	"wsu-senior-project/service"
)

// Endpoints holds all Go kit endpoints for the Hello service
type Endpoints struct {
	SayHelloEndpoint endpoint.Endpoint
}

// MakeEndpoints creates and returns all endpoints for the Hello service
func MakeEndpoints(svc service.HelloService) Endpoints {
	return Endpoints{
		SayHelloEndpoint: makeSayHelloEndpoint(svc),
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

func makeSayHelloEndpoint(svc service.HelloService) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(HelloRequest)
		msg := svc.SayHello(req.Name)
		return HelloResponse{Message: msg}, nil
	}
}
