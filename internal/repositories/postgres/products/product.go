package products

import (
	"github.com/gofiber/fiber/v2/log"
	"github.com/siti-nabila/backend-siti-nabila/internal/domain"
)

func (p *productRepository) AddProductMerchant(request domain.MerchantProduct) (results []domain.MerchantProduct, err error) {

	query := `
		WITH p AS (
			INSERT INTO products(product_name)
			VALUES ($1)
			RETURNING product_id
		)
	INSERT INTO pivot_merchant_listing(user_id, product_id, price, stock)
	SELECT $2, p.product_id, $3, $4 FROM p;
	`

	tx, err := p.db.Begin()
	if err != nil {
		log.Error(err)
		return results, err
	}
	stmt, err := tx.Prepare(query)
	defer stmt.Close()
	if err != nil {
		log.Error(err, err.Error())
		return results, err
	}

	_, err = stmt.Exec(request.ProductName, request.UserId, request.ProductPrice, request.ProductStock)
	if err != nil {
		tx.Rollback()
		log.Error(err, err.Error())
		return results, err
	}

	if err = tx.Commit(); err != nil {
		log.Error(err)
		return results, err
	}

	return results, err
}
