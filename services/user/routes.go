package user

import (
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/phildehovre/go-complete-backend/services/auth"
	"github.com/phildehovre/go-complete-backend/types"
	"github.com/phildehovre/go-complete-backend/utils"
)

type Handler struct {
	store *Store
}

func NewHandler(store *Store) *Handler {
	return &Handler{store: store}
}

func (h *Handler) RegisterRoutes(router *http.ServeMux) {
	router.HandleFunc("/home", h.handleHome)
	router.HandleFunc("/login", h.handleLogin)
	router.HandleFunc("/register", h.handleRegister)

}

func (h *Handler) handleHome(w http.ResponseWriter, r *http.Request) {
	var payload interface{}
	utils.ParseJSON(r, &payload)
	utils.WriteJSON(w, http.StatusOK, payload)
}

func (h *Handler) handleLogin(w http.ResponseWriter, r *http.Request) {

	if r.Method != "POST" {
		utils.WriteJSON(w, http.StatusMethodNotAllowed, fmt.Errorf("method is not allowed: %s", r.Method))
		return
	}
	var payload types.LoginUserPayload
	if err := utils.ParseJSON(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	// validate payload
	if err := utils.Validate.Struct(payload); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid payload: %v", errors))
		return
	}

	// fetch user for pw comparison
	user, err := h.store.GetUserByEmail(payload.Email)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}
	//  compare passwords
	if err := auth.ComparePasswords(user.Password, []byte(payload.Password)); err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("incorrect password or email"))
		return
	}

	// create jwt

	utils.WriteJSON(w, http.StatusOK, "login route: user logged in, no JWT yet")
}

func (h *Handler) handleRegister(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		utils.WriteJSON(w, http.StatusMethodNotAllowed, fmt.Errorf("method is not allowed: %s", r.Method))
		return
	}

	var payload types.RegisterUserPayload
	err := utils.ParseJSON(r, &payload)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	// validate payload
	if err := utils.Validate.Struct(payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	// check if user exists already
	user, err := h.store.GetUserByEmail(payload.Email)
	if user != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("user already exists: %s", payload.Email))
		return
	}

	hashedPassword, err := auth.HashPassword([]byte(payload.Password))
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, fmt.Errorf("user creation failed %v", err))
		return
	}

	newUser := &types.User{
		Firstname: payload.Firstname,
		Lastname:  payload.Lastname,
		Email:     payload.Email,
		Password:  hashedPassword,
	}

	err = h.store.CreateUser(newUser)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, fmt.Errorf("user creation failed %v", err))
		return
	}

	utils.WriteJSON(w, http.StatusOK, nil)

}
