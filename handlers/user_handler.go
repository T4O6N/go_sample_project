package handlers

import (
	"encoding/json"
	"net/http"
	"sample-project/internal/entities"
	"sample-project/internal/usecases/user"
)

type UserHandler struct {
	UserUseCase *user.UserUseCase
}

func NewUserHandler(userUseCase *user.UserUseCase) *UserHandler {
	return &UserHandler{UserUseCase: userUseCase}
}

// CreateUser godoc
// @Summary Create a new user
// @Description Create a new user with the input payload
// @Tags users
// @Accept  json
// @Produce  json
// @Param user body models.User true "Create user"
// @Success 201 {object} models.User
// @Router /users [post]
func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var user entities.User
	json.NewDecoder(r.Body).Decode(&user)
	err := h.UserUseCase.CreateUser(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}
