package users

import (
	"github.com/siti-nabila/backend-siti-nabila/internal/domain"
	"golang.org/x/crypto/bcrypt"
)

type userService struct {
	UserRepo domain.UserRepository
}

func NewUserService(userRepo domain.UserRepository) domain.UserService {
	return &userService{
		UserRepo: userRepo,
	}
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(reqPassword, dbPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(dbPassword), []byte(reqPassword))
	return err == nil
}
