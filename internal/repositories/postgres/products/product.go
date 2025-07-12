package products

import (
	"github.com/gofiber/fiber/v2/log"
	"github.com/siti-nabila/backend-siti-nabila/internal/domain"
)

// fungsi yang dibuat untuk menambahkan produk yang dijual merchant
func (p *productRepository) AddProductMerchant(request domain.MerchantProduct) error {
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
		return err
	}
	stmt, err := tx.Prepare(query)
	defer stmt.Close()
	if err != nil {
		log.Error(err, err.Error())
		return err
	}

	_, err = stmt.Exec(request.ProductName, request.UserId, request.ProductPrice, request.ProductStock)
	if err != nil {
		tx.Rollback()
		log.Error(err, err.Error())
		return err
	}

	if err = tx.Commit(); err != nil {
		log.Error(err)
		return err
	}

	return err
}

// fungsi untuk menampilkan produk yang dijual oleh merchant yang sedang login
func (p *productRepository) GetProductMerchantByUserId(userId int) (results []domain.MerchantProduct, err error) {
	query := `
	SELECT 
		ml.pivot_id,
		p.product_id,
		p.product_name,
		ml.price,
		ml.stock
	FROM pivot_merchant_listing ml
	JOIN products p
		ON ml.product_id = p.product_id
	WHERE ml.user_id = $1
	ORDER BY ml.pivot_id DESC
	`

	stmt, err := p.db.Prepare(query)
	defer stmt.Close()

	if err != nil {
		log.Error(err, err.Error())
		return results, err
	}

	rows, err := stmt.Query(userId)
	if err != nil {
		log.Error(err)
		return results, err
	}

	for rows.Next() {
		var result domain.MerchantProduct
		err = rows.Scan(
			&result.ProductListingId,
			&result.ProductId,
			&result.ProductName,
			&result.ProductPrice,
			&result.ProductStock,
		)
		if err != nil {
			log.Error(err)
			return results, err
		}
		results = append(results, result)
	}

	return results, err
}

// menampilkan history pembelian user
func (p *productRepository) GetProductCustomerByUserId(userId int) (results []domain.CustomerProduct, err error) {
	query := `
	SELECT 
		ml.pivot_id,
		ml.user_id as merchant_id,
		u.user_email,
		p.product_name,
		ci.price as product_cost,
		ci.qty,
		ci.ongkir,
		ci.discount_amount,
		ci.paid_amount
	FROM pivot_customer_items ci
	JOIN pivot_merchant_listing ml
		ON ci.merchant_listing_id = ml.pivot_id
	JOIN products p
		ON ml.product_id = p.product_id
	JOIN users u
		ON ml.user_id = u.user_id
	WHERE ci.user_id = $1
	ORDER BY ci.pivot_id DESC
	`
	stmt, err := p.db.Prepare(query)
	defer stmt.Close()

	if err != nil {
		log.Error(err, err.Error())
		return results, err
	}

	rows, err := stmt.Query(userId)
	if err != nil {
		log.Error(err)
		return results, err
	}

	for rows.Next() {
		var result domain.CustomerProduct
		err = rows.Scan(
			&result.ProductListingId,
			&result.MerchantId,
			&result.UserEmail,
			&result.ProductName,
			&result.ProductCost,
			&result.ProductQty,
			&result.Ongkir,
			&result.DiscountAmount,
			&result.PaidAmount,
		)
		if err != nil {
			log.Error(err)
			return results, err
		}
		results = append(results, result)
	}

	return results, err
}

// fungsi digunakan untuk customer yang membeli product dari merchant, dan mengurangi stock product si merchant sesuai dengan quantity
func (p *productRepository) AddCustomerItem(request domain.CustomerProduct) (err error) {

	query := `
		WITH p AS (
				SELECT 
					ml.pivot_id,
					p.product_id,
					p.product_name,
					ml.price,
					ml.stock
				FROM pivot_merchant_listing ml
				JOIN products p
					ON ml.product_id = p.product_id
				WHERE ml.pivot_id = $1
			)
		INSERT INTO pivot_customer_items(user_id, merchant_listing_id, qty, price, ongkir, discount_amount, paid_amount)
		SELECT $2, p.pivot_id, $3, $4, $5, $6, $7 FROM p;
	`
	tx, err := p.db.Begin()
	if err != nil {
		log.Error(err)
		return err
	}
	stmt, err := tx.Prepare(query)
	defer stmt.Close()
	if err != nil {
		log.Error(err, err.Error())
		return err
	}

	_, err = stmt.Exec(
		request.ProductListingId,
		request.UserId,
		request.ProductQty,
		request.ProductPriceAmount,
		request.Ongkir,
		request.DiscountAmount,
		request.PaidAmount,
	)
	if err != nil {
		tx.Rollback()
		log.Error(err, err.Error())
		return err
	}

	if err = tx.Commit(); err != nil {
		log.Error(err)
		return err
	}

	return err
}

// fungsi digunakan hanya untuk customer untuk melihat product yang dijual dari berbagai seller
func (p *productRepository) GetProducts() (results []domain.MerchantProduct, err error) {
	query := `
	SELECT
		ml.pivot_id,
		ml.user_id,
		u.user_email,
		p.product_id,
		p.product_name,
		ml.price,
		ml.stock
	FROM pivot_merchant_listing ml
	JOIN products p
		ON ml.product_id = p.product_id
	JOIN users u
		ON ml.user_id = u.user_id
	ORDER BY ml.user_id
	`
	stmt, err := p.db.Prepare(query)
	defer stmt.Close()

	if err != nil {
		log.Error(err, err.Error())
		return results, err
	}

	rows, err := stmt.Query()
	if err != nil {
		log.Error(err)
		return results, err
	}

	for rows.Next() {
		var result domain.MerchantProduct
		err = rows.Scan(
			&result.ProductListingId,
			&result.UserId,
			&result.UserEmail,
			&result.ProductId,
			&result.ProductName,
			&result.ProductPrice,
			&result.ProductStock,
		)
		if err != nil {
			log.Error(err)
			return results, err
		}
		results = append(results, result)
	}

	return results, err
}

func (p *productRepository) GetProductByListingId(listingId int) (result domain.MerchantProduct, err error) {
	query := `
	SELECT 
		ml.pivot_id,
		p.product_id,
		p.product_name,
		ml.price,
		ml.stock
	FROM pivot_merchant_listing ml
	JOIN products p
		ON ml.product_id = p.product_id
	WHERE ml.pivot_id = $1
	`
	stmt, err := p.db.Prepare(query)
	defer stmt.Close()

	if err != nil {
		log.Error(err, err.Error())
		return result, err
	}

	err = stmt.QueryRow(listingId).Scan(
		&result.ProductListingId,
		&result.ProductId,
		&result.ProductName,
		&result.ProductPrice,
		&result.ProductStock,
	)
	if err != nil {
		log.Error(err)
		return result, err
	}
	return result, err
}

func (p *productRepository) GetProductWithBuyer(userId int) (results []domain.MerchantProductWithBuyer, err error) {
	query := `
	with merchant_listing AS ( 
		SELECT 
			ml.pivot_id,
			p.product_id,
			p.product_name,
			ml.price,
			ml.stock
		FROM pivot_merchant_listing ml
		JOIN products p
			ON ml.product_id = p.product_id
		WHERE ml.user_id = $1
		ORDER BY ml.pivot_id DESC)
	SELECT 
		ml.pivot_id as listing_id,
		ml.product_name,
		ml.price,
		ml.stock,
			ci.price as product_cost,
			ci.qty,
			ci.ongkir,
			ci.discount_amount,
			ci.paid_amount,
		u.user_email
	FROM merchant_listing ml
	JOIN pivot_customer_items ci
		ON ml.pivot_id = ci.merchant_listing_id
	JOIN users u
		ON ci.user_id = u.user_id
	order by ml.product_id desc
	`
	stmt, err := p.db.Prepare(query)
	defer stmt.Close()

	if err != nil {
		log.Error(err, err.Error())
		return results, err
	}

	rows, err := stmt.Query(userId)
	if err != nil {
		log.Error(err)
		return results, err
	}

	for rows.Next() {
		var result domain.MerchantProductWithBuyer
		err = rows.Scan(
			&result.ProductListingId,
			&result.ProductName,
			&result.ProductPrice,
			&result.ProductStock,
			&result.ProductCost,
			&result.ProductQty,
			&result.Ongkir,
			&result.DiscountAmount,
			&result.PaidAmount,
			&result.UserEmail,
		)
		if err != nil {
			log.Error(err)
			return nil, err
		}
		results = append(results, result)
	}
	return results, err
}
