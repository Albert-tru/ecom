package user

import (
	"fmt"
	"net/http"

	"github.com/Albert-tru/ecom/service/auth"
	"github.com/Albert-tru/ecom/types"
	"github.com/Albert-tru/ecom/utils"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
)

type Handler struct {
	store types.UserStore
}

func NewHandler(store types.UserStore) *Handler {
	return &Handler{store: store}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/login", h.handleLogin).Methods("POST")
	router.HandleFunc("/register", h.handleRegister).Methods("POST")
}

func (h *Handler) handleLogin(w http.ResponseWriter, r *http.Request) {

}

// 处理用户注册
func (h *Handler) handleRegister(w http.ResponseWriter, r *http.Request) {
	//1. 获取json数据
	var payload types.RegisterUserPayload
	//解码过程中发生错误
	if err := utils.ParseJson(r, &payload); err != nil {
		//进行错误处理
		utils.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	//验证数据
	if err := utils.Validate.Struct(&payload); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("validation error: %v", errors.Error()).Error())
		return
	}

	//2. 检查user是否存在
	_, err := h.store.GetUserByEmail(payload.Email)
	if err == nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("user with email %s already exists", payload.Email).Error())
		return
	}

	hashedPassword, err := auth.HashPassword(payload.Password)

	//3. 不存在，创建user
	err = h.store.CreateUser(&types.User{
		Firstname: payload.Firstname,
		Lastname:  payload.Lastname,
		Email:     payload.Email,
		Password:  hashedPassword,
	})

	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.WriteJson(w, http.StatusCreated, nil)
}
