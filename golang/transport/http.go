package transport

import (
	"context"
	"encoding/json"
	"net/http"

	httptransport "github.com/go-kit/kit/transport/http"

	"wsu-senior-project/endpoint"
)

// NewHTTPHandler returns an HTTP handler for the service
func NewHTTPHandler(endpoints endpoint.Endpoints) http.Handler {
	mux := http.NewServeMux()

	mux.Handle("/hello", httptransport.NewServer(
		endpoints.SayHelloEndpoint,
		decodeHelloRequest,
		encodeResponse,
	))

	mux.Handle("/users", httptransport.NewServer(
		endpoints.CreateUserEndpoint,
		decodeCreateUserRequest,
		encodeResponse,
	))

	return mux
}

func decodeHelloRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request endpoint.HelloRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return endpoint.HelloRequest{}, nil // Return empty request on decode error
	}
	return request, nil
}

func decodeCreateUserRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request endpoint.CreateUserRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return endpoint.CreateUserRequest{}, nil
	}
	return request, nil
}

func encodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(response)
}
