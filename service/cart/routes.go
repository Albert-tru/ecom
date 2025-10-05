package cart

import (
	"net/http"

	"github.com/Albert-tru/ecom/service/auth" // ✅ 添加这行
	"github.com/Albert-tru/ecom/types"
	"github.com/Albert-tru/ecom/utils"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
)

type Handler struct {
	store        types.OrderStore
	productStore types.ProductStore
	userStore    types.UserStore
}

func NewHandler(store types.OrderStore, productStore types.ProductStore, userStore types.UserStore) *Handler {
	return &Handler{
		store:        store,
		productStore: productStore,
		userStore:    userStore,
	}
}
func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/cart/checkout", auth.WithJWTAuth(h.handleCheckout, h.userStore)).Methods("POST")
}

func (h *Handler) handleCheckout(w http.ResponseWriter, r *http.Request) {
	userID := auth.GetUserIDFromContext(r.Context())
	var cart types.CartCheckoutPayload
	if err := utils.ParseJson(r, &cart); err != nil {
		utils.WriteJson(w, http.StatusBadRequest, map[string]string{"error": "Invalid request payload"})
		return
	}

	if err := utils.Validate.Struct(cart); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteJson(w, http.StatusBadRequest, map[string]interface{}{"error": "Validation failed", "details": errors.Error()})
		return
	}

	//获取产品
	productIDs, err := getCartItemsIDs(cart.Items)
	if err != nil {
		utils.WriteJson(w, http.StatusBadRequest, map[string]string{"error": "Invalid cart items"})
		return
	}

	ps, err := h.productStore.GetProductByIDs(productIDs)
	if err != nil {
		utils.WriteJson(w, http.StatusInternalServerError, map[string]string{"error": "Failed to retrieve products"})
		return
	}

	orderID, totalPrice, err := h.CreateOrder(ps, cart.Items, userID)
	if err != nil {
		utils.WriteJson(w, http.StatusInternalServerError, map[string]string{"error": "Failed to create order"})
		return
	}

	utils.WriteJson(w, http.StatusOK, map[string]interface{}{
		"status":     "success",
		"orderId":    orderID,
		"totalPrice": totalPrice,
	})

}
