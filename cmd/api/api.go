package api

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/phildehovre/go-complete-backend/services/user"
)

type APIServer struct {
	addr string
	db   *sql.DB
}

func NewAPIServer(addr string, db *sql.DB) *APIServer {
	return &APIServer{addr: addr, db: db}
}

func (s *APIServer) Run() {
	router := http.NewServeMux()

	userStore := user.NewStore(s.db)
	userHandler := user.NewHandler(userStore)
	userHandler.RegisterRoutes(router)

	log.Fatal(http.ListenAndServe(s.addr, router))
}
