package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/siti-nabila/backend-siti-nabila/internal/inject"
	"github.com/spf13/viper"
)

func main() {
	Init()
	app := fiber.New()
	app.Use(cors.New())
	inject.Inject(app)
	app.Listen(viper.GetString("APP_PORT"))
}

func Init() {
	viper.SetConfigFile(".env")
	viper.SetConfigType("env")
	viper.AddConfigPath("../../")
	// Read the config file
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// Config file not found, handle error or proceed without it
			log.Println("Config file .env not found, using environment variables if set.")
		} else {
			log.Fatalf("Fatal error reading config file: %s \n", err)
		}
	}

}
