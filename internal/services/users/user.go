package users

import "github.com/siti-nabila/backend-siti-nabila/internal/domain"

func (u *userService) GetUserByUserId(userId int) (result domain.User, err error) {
	return u.UserRepo.GetUserByUserId(userId)
}
