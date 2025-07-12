package domain

import (
	"github.com/gofiber/fiber/v2"
	"github.com/siti-nabila/backend-siti-nabila/internal/models"
)

type (
	User struct {
		UserId       int
		UserEmail    string
		UserPassword string
		RoleId       int
		RoleName     string
	}
)

type UserRepository interface {
	Register(request models.RegisterRequest) (User, error)
	Login(request models.LoginReqeust) (User, error)
	GetUserByUserId(userId int) (User, error)
}

type UserService interface {
	Register(request models.RegisterRequest) (User, error)
	Login(request models.LoginReqeust) (User, error)
	GetUserByUserId(userId int) (User, error)
}

type UserHandler interface {
	Register(*fiber.Ctx) error
	Login(*fiber.Ctx) error
}
