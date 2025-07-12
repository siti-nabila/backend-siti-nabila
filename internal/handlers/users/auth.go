package user

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/siti-nabila/backend-siti-nabila/internal/models"
	"github.com/siti-nabila/backend-siti-nabila/internal/util"
)

func (u *userHandler) Register(c *fiber.Ctx) error {
	var (
		request  models.RegisterRequest
		response models.AuthResponse
	)

	if err := c.BodyParser(&request); err != nil {
		response.ErrorMessage = err.Error()
		return c.Status(fiber.StatusConflict).JSON(response)
	}

	res, err := u.userService.Register(request)
	if err != nil {
		response.ErrorMessage = err.Error()
		return c.Status(fiber.StatusUnauthorized).JSON(response)
	}

	token, err := util.GenerateJWTToken(res.UserId)
	if err != nil {
		log.Error(err)
		response.ErrorMessage = "failed to generate token"
		return c.Status(fiber.StatusUnauthorized).JSON(response)
	}

	response.Token = token
	return c.Status(fiber.StatusOK).JSON(response)
}

func (u *userHandler) Login(c *fiber.Ctx) error {
	var (
		request  models.LoginReqeust
		response models.AuthResponse
	)
	if err := c.BodyParser(&request); err != nil {
		response.ErrorMessage = err.Error()
		return c.Status(fiber.StatusConflict).JSON(response)
	}

	res, err := u.userService.Login(request)
	if err != nil {
		response.ErrorMessage = err.Error()
		return c.Status(fiber.StatusUnauthorized).JSON(response)
	}
	fmt.Println("[login] res : ", res)

	token, err := util.GenerateJWTToken(res.UserId)
	if err != nil {
		log.Error(err)
		response.ErrorMessage = "failed to generate token"
		return c.Status(fiber.StatusUnauthorized).JSON(response)
	}

	response.Token = token
	return c.Status(fiber.StatusOK).JSON(response)

}
