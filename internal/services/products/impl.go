package products

import "github.com/siti-nabila/backend-siti-nabila/internal/domain"

type productService struct {
	ProductRepo domain.ProductRepository
	SettingRepo domain.SettingRepository
}

func NewProductService(productRepo domain.ProductRepository, settingRepo domain.SettingRepository) domain.ProductService {
	return &productService{
		ProductRepo: productRepo,
		SettingRepo: settingRepo,
	}
}
