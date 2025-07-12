package user

import "github.com/siti-nabila/backend-siti-nabila/internal/domain"

type userHandler struct {
	userService domain.UserService
}

func NewUserHandler(userService domain.UserService) *userHandler {
	return &userHandler{
		userService: userService,
	}
}
