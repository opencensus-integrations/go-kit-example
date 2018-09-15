// Package kitoc is the base for an example of a Go kit service instrumented
// with OpenCensus.
package kitoc

import "context"

// Service provides our OpenCensus instrumented example service
type Service interface {
	Hello(ctx context.Context, firstName string, lastName string) (greeting string, err error)
}
