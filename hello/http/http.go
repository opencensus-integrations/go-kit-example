package http

import "context"
import "encoding/json"

import "net/http"

import httptransport "github.com/go-kit/kit/transport/http"
import "github.com/opencensus-integrations/go-kit-example/hello/endpoints"

func NewHTTPHandler(endpoints endpoints.Endpoints) http.Handler {
	m := http.NewServeMux()
	m.Handle("/hello", httptransport.NewServer(endpoints.Hello, DecodeHelloRequest, EncodeHelloResponse))
	return m
}
func DecodeHelloRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req endpoints.HelloRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}
func EncodeHelloResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}
