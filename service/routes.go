package user

import (
	"net/http"

	"github.com/Albert-tru/ecom/types"
	"github.com/gorilla/mux"
)

type handler struct {
	store types.UserStore
}

func NewHandler(store types.UserStore) *handler {
	return &handler{
		store: store,
	}
}

func (h *handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/login", h.handleLogin).Methods("POST")
	router.HandleFunc("/register", h.handleRegister).Methods("POST")
}

func (h *handler) handleLogin(w http.ResponseWriter, r *http.Request) {

}

func (h *handler) handleRegister(w http.ResponseWriter, r *http.Request) {

}
