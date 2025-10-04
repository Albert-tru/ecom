package product

import (
	"net/http"

	"github.com/Albert-tru/ecom/types"
	"github.com/Albert-tru/ecom/utils"
	"github.com/gorilla/mux"
)

type Handler struct {
	store types.ProductStore
}

func NewHandler(store types.ProductStore) *Handler {
	return &Handler{
		store: store,
	}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/products", h.handleGetProducts).Methods("GET")
}

func (h *Handler) handleGetProducts(w http.ResponseWriter, r *http.Request) {
	ps, err := h.store.GetProducts()
	if err != nil {
		http.Error(w, "Failed to get products", http.StatusInternalServerError)
		return
	}

	utils.WriteJson(w, http.StatusOK, ps)
}
