package server

import (
	"net/http"
	"strconv"
)

type options struct {
	port    string
	domain  string
	handler http.Handler
}

type opt func(opts *options)

// WithPort is a builder function to set a specific port to a server.
// Default is "8080".
func WithPort(port int) opt {
	return func(opts *options) {
		opts.port = ":" + strconv.Itoa(port)
	}
}

// WithDomain is a builder function to set a specific domain to a server.
// Default is "localhost".
func WithDomain(domain string) opt {
	return func(opts *options) {
		opts.domain = domain
	}
}

// WithHandler is a builder function to set a specific handler for routing to a server.
// Default is "http.ServeMux".
func WithHandler(handler http.Handler) opt {
	return func(opts *options) {
		opts.handler = handler
	}
}
