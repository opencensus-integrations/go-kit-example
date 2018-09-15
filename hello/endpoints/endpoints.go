package endpoints

import "context"

import "github.com/go-kit/kit/endpoint"

import "github.com/opencensus-integrations/go-kit-example/hello/service"

type HelloRequest struct {
	FirstName string
	LastName  string
}
type HelloResponse struct {
	Greeting string
	Err      error
}

func MakeHelloEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(HelloRequest)
		greeting, err := s.Hello(ctx, req.FirstName, req.LastName)
		return HelloResponse{Greeting: greeting, Err: err}, nil
	}
}

type Endpoints struct {
	Hello endpoint.Endpoint
}
