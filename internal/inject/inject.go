package inject

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/siti-nabila/backend-siti-nabila/internal/db"
	userHandler "github.com/siti-nabila/backend-siti-nabila/internal/handlers/users"
	userRepo "github.com/siti-nabila/backend-siti-nabila/internal/repositories/postgres/users"
	userService "github.com/siti-nabila/backend-siti-nabila/internal/services/users"
)

func Inject(server *fiber.App) {
	db, err := db.Open()
	if err != nil {
		log.Error(err, err.Error())
		return
	}

	userRepo := userRepo.NewUserPostgresRepository(db)
	userService := userService.NewUserService(userRepo)
	userHandler := userHandler.NewUserHandler(userService)

	// middleware
	// authCustomer := middlewares.RoleAuthorization(userService, 2)
	// authMerchant := middlewares.RoleAuthorization(userService, 1)

	// auth routes, semua role bisa akses
	server.Post("/register", userHandler.Register)
	server.Post("/login", userHandler.Login)

}
