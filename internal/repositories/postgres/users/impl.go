package users

import (
	"database/sql"

	"github.com/siti-nabila/backend-siti-nabila/internal/domain"
)

type userRepository struct {
	db *sql.DB
}

func NewUserPostgresRepository(db *sql.DB) domain.UserRepository {
	return &userRepository{
		db: db,
	}
}
