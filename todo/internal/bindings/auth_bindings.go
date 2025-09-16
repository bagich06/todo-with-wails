package bindings

import (
	"todo/internal/models"
	"todo/internal/services"
)

type AuthBinding struct {
	userService *services.UserService
}

func NewAuthBinding(userService *services.UserService) *AuthBinding {
	return &AuthBinding{userService: userService}
}

func (a *AuthBinding) Login(email, password string) (*models.AuthResponse, error) {
	return a.userService.Login(email, password)
}

func (a *AuthBinding) Register(username, email, password string) (*models.AuthResponse, error) {
	return a.userService.Register(username, email, password)
}
