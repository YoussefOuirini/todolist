package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/youssefouirini/todolist/cmd/controller"
	"github.com/youssefouirini/todolist/cmd/storage"
)

const (
	shutdownTimeout = 10 * time.Second
	port            = 8080
)

func main() {
	log.Printf("Starting server at port: %v", port)

	db, err := gorm.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		panic(err)
	}

	api := NewAPIServer(port, db, storage.NewToDoRepository())
	err = api.ListenAndServe()
	if err != nil {
		log.Fatalf("http server error: %s", err)
	}
}

type APIServer struct {
	server *http.Server
	repo   storage.ToDoRepository
	db     *gorm.DB
}

func NewAPIServer(port int, db *gorm.DB, repo storage.ToDoRepository) *APIServer {
	s := &APIServer{
		server: &http.Server{
			Addr: ":" + strconv.Itoa(port),
		},
		db:   db,
		repo: repo,
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
	controller.NewController(s.server, s.db, s.repo)

	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "yoyoyoyoyo!")
	})
}
