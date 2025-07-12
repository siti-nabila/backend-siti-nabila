package products

import (
	"encoding/json"

	"github.com/gofiber/fiber/v2/log"
	"github.com/siti-nabila/backend-siti-nabila/internal/domain"
	"github.com/siti-nabila/backend-siti-nabila/internal/models"
)

func (p *productService) GetAllProduct() (result models.ListingProductCustomerResponse, err error) {
	result.Products = make([]models.ProductCustomer, 0)
	res, err := p.ProductRepo.GetProducts()
	if err != nil {
		log.Error(err)
		return result, err
	}
	for _, v := range res {
		data := models.ProductCustomer{
			ProductListingId: v.ProductListingId,
			ProductName:      v.ProductName,
			ProductPrice:     v.ProductPrice,
			ProductStock:     v.ProductStock,
			MerchantEmail:    v.UserEmail,
		}
		result.Products = append(result.Products, data)
	}
	return result, err
}

func (p *productService) GetHistoryItems(userId int) (result models.ListingProductCustomerResponse, err error) {
	result.Products = make([]models.ProductCustomer, 0)

	res, err := p.ProductRepo.GetProductCustomerByUserId(userId)
	if err != nil {
		log.Error(err)
		return result, err
	}
	for _, v := range res {
		data := models.ProductCustomer{
			ProductListingId:       v.ProductListingId,
			ProductName:            v.ProductName,
			MerchantEmail:          v.UserEmail,
			ProductPrice:           v.ProductCost,
			ProductQty:             v.ProductQty,
			ProductOngkir:          v.Ongkir,
			ProductDiscountAmount:  v.DiscountAmount,
			ProdcutTotalPaidAmount: v.PaidAmount,
		}
		result.Products = append(result.Products, data)
	}

	return result, err
}

func (p *productService) BuyProduct(request models.InsertCustomerProductRequest) (result models.ListingProductCustomerResponse, err error) {
	result.Products = make([]models.ProductCustomer, 0)
	var setting domain.PaymentRule
	discAmt := 0
	sett, err := p.SettingRepo.GetSettingByKey("payment_rule")
	if err != nil {
		log.Error(err)
		return result, err
	}
	if err = json.Unmarshal([]byte(sett.SettingValue), &setting); err != nil {
		log.Error(err)
		return result, err
	}

	product, err := p.ProductRepo.GetProductByListingId(request.ProductListingId)
	if err != nil {
		log.Error(err)
		return result, err
	}

	totalPaid := request.ProductQty * product.ProductPrice
	if totalPaid > setting.MinPurchaseDisc {
		discAmt = int(float64(totalPaid) * setting.DiscountValue)
		if ttl := totalPaid - discAmt; ttl < 0 {
			totalPaid = 0
		}
		totalPaid -= discAmt

	}
	if totalPaid > setting.MinPurchaseFreeOngkir {
		setting.Ongkir = 0
	}
	totalPaid += setting.Ongkir

	insReq := domain.CustomerProduct{
		ProductListingId:   request.ProductListingId,
		UserId:             request.UserId,
		ProductQty:         request.ProductQty,
		DiscountAmount:     discAmt,
		Ongkir:             setting.Ongkir,
		PaidAmount:         totalPaid,
		ProductPriceAmount: product.ProductPrice * request.ProductQty,
	}

	err = p.ProductRepo.AddCustomerItem(insReq)
	if err != nil {
		log.Error(err)
		return result, err
	}

	result, err = p.GetHistoryItems(request.UserId)
	if err != nil {
		log.Error(err)
		return result, err
	}

	return result, err
}
