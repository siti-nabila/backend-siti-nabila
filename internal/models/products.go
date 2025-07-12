package models

type (
	InsertProductRequest struct {
		ProductName  string `json:"product_name"`
		ProductPrice int    `json:"product_price"`
		ProductStock int    `json:"product_stock"`
	}

	ListingProductMerchant struct {
		Products []Product `json:"listing_product"`
	}
	Product struct {
		ProductName  string `json:"product_name"`
		ProductPrice int    `json:"product_price"`
		ProductStock int    `json:"product_stock"`
	}
)
