package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"strconv"

	"github.com/saiddis/echo_service/server"
)

func main() {
	domain := os.Getenv("DOMAIN")
	log.Printf("domain: %s", domain)

	port, err := strconv.Atoi(os.Getenv("PORT"))
	log.Printf("port: %d", port)
	if err != nil {
		log.Printf("error converting env var to int: %v", err)
		port = 8080
	}

	s := server.New(
		server.WithPort(port),
		server.WithDomain(domain),
	)

	ctx, cancel := context.WithCancel(context.Background())
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() { <-c; cancel() }()

	if err := s.Open(); err != nil {
		log.Fatalf("error openning connection with the server: %v", err)
	}

	log.Printf("running: url=%q", s.URL())

	<-ctx.Done()

	if err := s.Close(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
