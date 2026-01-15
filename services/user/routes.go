package user

import (
	"net/http"

	"github.com/HadeedTariq/go-ecom-api/types"
	"github.com/gorilla/mux"
)

type Handler struct{}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	subRouter := router.PathPrefix("/auth").Subrouter()
	subRouter.HandleFunc("/register", h.RegsiterUser).Methods("POST")
}

// ~ so now the first thing I am gonna do is create the schema for the user
func (h *Handler) RegsiterUser(w http.ResponseWriter, r *http.Request) {
	// ~ so over there for registering the user certain type of validation and the data type definition is required

	var payload *types.RegisterUserPayload

	// ~ so over there using the validator package I have to validate that out
}
