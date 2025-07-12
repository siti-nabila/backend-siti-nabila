package inject

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/siti-nabila/backend-siti-nabila/internal/db"
	productHandler "github.com/siti-nabila/backend-siti-nabila/internal/handlers/products"
	userHandler "github.com/siti-nabila/backend-siti-nabila/internal/handlers/users"
	"github.com/siti-nabila/backend-siti-nabila/internal/middlewares"
	productRepo "github.com/siti-nabila/backend-siti-nabila/internal/repositories/postgres/products"
	settingRepo "github.com/siti-nabila/backend-siti-nabila/internal/repositories/postgres/settings"
	userRepo "github.com/siti-nabila/backend-siti-nabila/internal/repositories/postgres/users"
	productService "github.com/siti-nabila/backend-siti-nabila/internal/services/products"
	userService "github.com/siti-nabila/backend-siti-nabila/internal/services/users"
)

func Inject(server *fiber.App) {
	db, err := db.Open()
	if err != nil {
		log.Error(err, err.Error())
		return
	}

	// repositories
	userRepo := userRepo.NewUserPostgresRepository(db)
	productRepo := productRepo.NewProductPostgresRepository(db)
	settingRepo := settingRepo.NewSettingPostgresRepository(db)
	// services
	userService := userService.NewUserService(userRepo)
	productService := productService.NewProductService(productRepo, settingRepo)

	// handlers
	userHandler := userHandler.NewUserHandler(userService)
	productHandler := productHandler.NewProductHandler(productService)

	// middleware
	authCustomer := middlewares.RoleAuthorization(userService, 2)
	authMerchant := middlewares.RoleAuthorization(userService, 1)

	// auth routes, semua role bisa akses
	server.Post("/register", userHandler.Register)
	server.Post("/login", userHandler.Login)

	// product-customer only
	server.Get("/products", authCustomer, productHandler.GetAllProducts)
	server.Get("/products/history", authCustomer, productHandler.GetPurchasedItemHistory)
	server.Post("/checkout", authCustomer, productHandler.BuyProduct)

	// product-merchant only
	server.Get("/product/listing", authMerchant, productHandler.GetListingProducts)
	server.Get("/product/listing-with-buyer", authMerchant, productHandler.GetListingProductsWithBuyer)

	server.Post("/product/add", authMerchant, productHandler.AddNewProduct)
}
