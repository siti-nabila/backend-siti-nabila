package domain

import (
	"github.com/gofiber/fiber/v2"
	"github.com/siti-nabila/backend-siti-nabila/internal/models"
)

type (
	MerchantProduct struct {
		UserId           int
		UserEmail        string
		ProductListingId int
		ProductId        int
		ProductName      string
		ProductPrice     int
		ProductStock     int
	}

	CustomerProduct struct {
		MerchantId         int
		UserId             int
		UserEmail          string
		ProductListingId   int
		ProductName        string
		ProductQty         int
		ProductPriceAmount int
		ProductCost        int
		Ongkir             int
		DiscountAmount     int
		PaidAmount         int
	}

	MerchantProductWithBuyer struct {
		ProductListingId   int
		ProductId          int
		ProductName        string
		ProductPrice       int
		ProductStock       int
		UserEmail          string
		ProductQty         int
		ProductPriceAmount int
		ProductCost        int
		Ongkir             int
		DiscountAmount     int
		PaidAmount         int
	}

	Buyer struct {
	}
	ProductRepository interface {
		GetProductMerchantByUserId(userId int) ([]MerchantProduct, error)
		GetProductCustomerByUserId(userId int) ([]CustomerProduct, error)
		GetProducts() ([]MerchantProduct, error)
		GetProductByListingId(listingId int) (result MerchantProduct, err error)
		GetProductWithBuyer(userId int) (results []MerchantProductWithBuyer, err error)

		AddProductMerchant(MerchantProduct) error
		AddCustomerItem(request CustomerProduct) error
	}
	ProductService interface {
		// merchant
		GetMerchantListingProducts(userId int) (models.ListingProductMerchantResponse, error)
		AddMerchantListingProduct(models.InsertMerchantProductRequest) (models.ListingProductMerchantResponse, error)
		GetProductWithBuyer(userId int) (results models.ListingProductMerchantWithBuyer, err error)

		// customer
		GetAllProduct() (models.ListingProductCustomerResponse, error)
		GetHistoryItems(userId int) (models.ListingProductCustomerResponse, error)
		BuyProduct(models.InsertCustomerProductRequest) (models.ListingProductCustomerResponse, error)
	}

	ProductHandler interface {
		// merchant
		AddNewProduct(*fiber.Ctx) error
		GetListingProducts(*fiber.Ctx) error
		GetListingProductsWithBuyer(*fiber.Ctx) error
		// customer
		GetAllProducts(*fiber.Ctx) error
		GetPurchasedItemHistory(*fiber.Ctx) error
		BuyProduct(*fiber.Ctx) error
	}
)
