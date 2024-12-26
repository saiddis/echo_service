package server

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"path"
	"strings"
	"time"

	"golang.org/x/crypto/acme/autocert"
)

type Server struct {
	ln     net.Listener
	server *http.Server
	router *http.ServeMux
	addr   string
	domain string
}

// New returns a new instance of Server with functional options as arguments.
func New(opts ...opt) *Server {
	s := &Server{
		server: &http.Server{},
		router: http.NewServeMux(),
		addr:   ":8080",
	}
	var options options

	for _, opt := range opts {
		opt(&options)
	}

	if options.domain != "" {
		s.domain = options.domain
	}
	if options.port != "" {
		s.addr = options.port
	}
	// if options.handler != nil {
	// 	s.router = options.handler
	// }

	s.server.Handler = http.HandlerFunc(s.serveHTTP)

	s.router.HandleFunc("/", handleRequests)

	return s
}

func (s *Server) Open() (err error) {
	if s.domain != "" && s.domain != "localhost" {
		s.ln = autocert.NewListener(s.domain)
	} else {
		if s.ln, err = net.Listen("tcp", s.addr); err != nil {
			return err
		}
	}

	go s.server.Serve(s.ln)
	return nil
}

func (s *Server) Close() error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	return s.server.Shutdown(ctx)
}

func (s *Server) Port() int {
	if s.ln == nil {
		return 0
	}
	return s.ln.Addr().(*net.TCPAddr).Port
}

func (s *Server) URL() string {
	scheme, port := s.Scheme(), s.Port()

	domain := "localhost"
	if s.domain != "" {
		domain = s.domain
	}

	if (scheme == "http" && port == 80) || (scheme == "https" && port == 443) {
		return fmt.Sprintf("%s:%s", scheme, domain)
	}

	return fmt.Sprintf("%s://%s:%d", scheme, domain, port)
}

func (s *Server) UseTLS() bool {
	return s.domain != ""
}

func (s *Server) Scheme() string {
	if s.UseTLS() {
		return "https"
	}

	return "http"
}

func (s *Server) serveHTTP(w http.ResponseWriter, r *http.Request) {
	log.Printf("hello world")
	ext := path.Ext(r.URL.Path)
	r.URL.Path = strings.TrimSuffix(r.URL.Path, ext)
	s.router.ServeHTTP(w, r)
}
