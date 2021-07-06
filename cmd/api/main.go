package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"
)

const (
	shutdownTimeout = 10 * time.Second
	port            = 8080
)

func main() {
	log.Printf("Starting server at port: %v", port)

	api := NewAPIServer(port)
	err := api.ListenAndServe()
	if err != nil {
		log.Fatalf("http server error: %s", err)
	}
}

type APIServer struct {
	server *http.Server
}

func NewAPIServer(port int) *APIServer {
	s := &APIServer{
		server: &http.Server{
			Addr: ":" + strconv.Itoa(port),
		},
	}

	s.mount()

	return s
}

func (s *APIServer) ListenAndServe() error {
	return s.server.ListenAndServe()
}

func (s *APIServer) Shutdown() {
	ctx, cancel := context.WithTimeout(context.Background(), shutdownTimeout)
	defer cancel()

	if err := s.server.Shutdown(ctx); err != nil {
		fmt.Println("error shutting down http server gracefully")

		return
	}

	fmt.Println("server shutdown gracefully")
}

func (s *APIServer) mount() {
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "yoyoyoyoyo!")
	})
}
