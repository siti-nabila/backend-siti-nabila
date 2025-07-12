package domain

type (
	MerchantProduct struct {
		UserId       int
		ProductId    int
		ProductName  string
		ProductPrice int
		ProductStock int
	}

	CustomerProduct struct {
		UserId       int
		ProductId    int
		ProductName  string
		ProductPrice int
		ProductQty   int
	}
	ProductRepository interface {
		// GetProductByUserId(userId int, viewType string)
		AddProductMerchant(MerchantProduct) ([]MerchantProduct, error)
	}
	// ProductService interface {
	// 	GetProductByUserId(userId int, viewType string)
	// 	AddProductMerchant(MerchantProduct) ([]MerchantProduct, error)
	// }
)
