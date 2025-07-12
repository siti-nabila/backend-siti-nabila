package db

import (
	"database/sql"
	"fmt"

	"github.com/gofiber/fiber/v2/log"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
)

func Open() (db *sql.DB, err error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		viper.GetString("DB_HOST"),
		viper.GetString("DB_USER"),
		viper.GetString("DB_PASS"),
		viper.GetString("DB_NAME"),
		viper.GetString("DB_PORT"),
		viper.GetString("DB_SSL_MODE"),
	)
	db, err = sql.Open("postgres", dsn)
	if err != nil {
		log.Error("error when opening postgres : ", err)
		return db, err
	}
	return db, err
}
