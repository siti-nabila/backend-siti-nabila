package products

import (
	"database/sql"

	"github.com/siti-nabila/backend-siti-nabila/internal/domain"
)

type productRepository struct {
	db *sql.DB
}

func NewProductPostgresRepository(db *sql.DB) domain.ProductRepository {
	return &productRepository{
		db: db,
	}
}
