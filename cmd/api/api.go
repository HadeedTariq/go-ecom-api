package api

import (
	"database/sql"
	"net/http"

	"github.com/gorilla/mux"
)

type ApiServer struct {
	Addr string
	Db   *sql.DB
}

func NewApiServer(addr string, db *sql.DB) *ApiServer {
	return &ApiServer{
		Addr: addr,
		Db:   db,
	}
}

func (server *ApiServer) Run() error {
	router := mux.NewRouter()

	subRouter := router.PathPrefix("/api/v1").Subrouter()

	return http.ListenAndServe(server.Addr, subRouter)
}
