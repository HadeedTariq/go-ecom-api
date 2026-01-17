package user

import (
	"fmt"
	"net/http"

	"github.com/HadeedTariq/go-ecom-api/types"
	"github.com/HadeedTariq/go-ecom-api/utils"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
)

type Handler struct {
	store types.UserStore
}

func NewHandler(store types.UserStore) *Handler {
	return &Handler{
		store: store,
	}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	subRouter := router.PathPrefix("/auth").Subrouter()
	subRouter.HandleFunc("/register", h.RegsiterUser).Methods("POST")
}

// ~ so now the first thing I am gonna do is create the schema for the user
func (h *Handler) RegsiterUser(w http.ResponseWriter, r *http.Request) {
	var payload types.RegisterUserPayload
	if err := utils.ParseJson(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	if err := utils.Validate.Struct(payload); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid payload: %v", errors))
		return
	}

	// ~ so after validation the next step is to check that the user exist with in the database or not

}
