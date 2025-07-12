package users

import (
	"github.com/gofiber/fiber/v2/log"
	"github.com/siti-nabila/backend-siti-nabila/internal/domain"
	"github.com/siti-nabila/backend-siti-nabila/internal/models"
)

func (u *userService) Register(request models.RegisterRequest) (result domain.User, err error) {
	if request.RoleId == nil {
		*request.RoleId = 2
	}
	pass, err := HashPassword(request.Password)
	if err != nil {
		log.Error(err, err.Error())
		return result, err
	}
	request.Password = pass

	result, err = u.UserRepo.Register(request)
	if err != nil {
		log.Error(err, err.Error())
		return result, err
	}
	return result, err
}

func (u *userService) Login(request models.LoginReqeust) (result domain.User, err error) {
	hashedPass, err := HashPassword(request.Password)
	if err != nil {
		log.Error(err, err.Error())
		return result, err
	}
	res, err := u.UserRepo.Login(request)
	if err != nil {
		log.Error(err, err.Error())
		return result, err
	}

	if CheckPasswordHash(res.UserPassword, hashedPass) {
		return res, err
	}

	return result, err
}
