package users

import (
	"errors"

	"github.com/gofiber/fiber/v2/log"
	"github.com/siti-nabila/backend-siti-nabila/internal/domain"
	"github.com/siti-nabila/backend-siti-nabila/internal/models"
)

func (u *userService) Register(request models.RegisterRequest) (result domain.User, err error) {

	if request.RoleId == nil {
		def := 2
		request.RoleId = &def
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
	res, err := u.UserRepo.Login(request)
	if err != nil {
		log.Error(err, err.Error())
		return result, err
	}
	if CheckPasswordHash(request.Password, res.UserPassword) {
		return res, err
	}

	return result, errors.New("something went wrong")
}
