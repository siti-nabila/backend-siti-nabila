package products

import "github.com/siti-nabila/backend-siti-nabila/internal/domain"

type productHandler struct {
	productService domain.ProductService
}

func NewProductHandler(productService domain.ProductService) domain.ProductHandler {
	return &productHandler{
		productService: productService,
	}
}
