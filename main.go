package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"

	"github.com/saiddis/echo_service/server"
)

func main() {
	s := server.New(
		server.WithPort(443),
		server.WithDomain("echoServer"),
	)

	ctx, cancel := context.WithCancel(context.Background())
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() { <-c; cancel() }()

	if err := s.Open(); err != nil {
		log.Fatalf("error openning connection with the server: %v", err)
	}

	log.Printf("running: url=%s", s.URL())

	<-ctx.Done()

	if err := s.Close(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
