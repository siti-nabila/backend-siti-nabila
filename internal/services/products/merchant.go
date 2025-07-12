package products

import (
	"github.com/gofiber/fiber/v2/log"
	"github.com/siti-nabila/backend-siti-nabila/internal/domain"
	"github.com/siti-nabila/backend-siti-nabila/internal/models"
)

func (p *productService) GetMerchantListingProducts(userId int) (results models.ListingProductMerchantResponse, err error) {
	results.Products = make([]models.Product, 0)

	res, err := p.ProductRepo.GetProductMerchantByUserId(userId)
	if err != nil {
		log.Error(err)
		return results, err
	}

	for _, v := range res {
		data := models.Product{
			ProductListingId: v.ProductListingId,
			ProductName:      v.ProductName,
			ProductPrice:     v.ProductPrice,
			ProductStock:     v.ProductStock,
		}
		results.Products = append(results.Products, data)
	}
	return results, err
}

func (p *productService) AddMerchantListingProduct(request models.InsertMerchantProductRequest) (result models.ListingProductMerchantResponse, err error) {

	insReq := domain.MerchantProduct{
		UserId:       request.UserId,
		ProductName:  request.ProductName,
		ProductPrice: request.ProductPrice,
		ProductStock: request.ProductStock,
	}
	err = p.ProductRepo.AddProductMerchant(insReq)
	if err != nil {
		log.Error(err)
		return result, err
	}
	result, err = p.GetMerchantListingProducts(request.UserId)
	if err != nil {
		log.Error(err)
		return result, err
	}

	return result, err
}

func (p *productService) GetProductWithBuyer(userId int) (results models.ListingProductMerchantWithBuyer, err error) {

	grouped := make(map[models.Product][]models.ProductBuyer, 0)

	datas, err := p.ProductRepo.GetProductWithBuyer(userId)
	if err != nil {
		log.Error(err)
		return results, err
	}

	for _, v := range datas {
		prod := models.Product{
			ProductListingId: v.ProductListingId,
			ProductName:      v.ProductName,
			ProductPrice:     v.ProductPrice,
			ProductStock:     v.ProductStock,
		}
		buyer := models.ProductBuyer{
			BuyerEmail:             v.UserEmail,
			ProductQty:             v.ProductQty,
			ProductOngkir:          v.Ongkir,
			ProductDiscountAmount:  v.DiscountAmount,
			ProdcutTotalPaidAmount: v.PaidAmount,
			ProductSubTotalAmount:  v.ProductCost,
		}
		grouped[prod] = append(grouped[prod], buyer)
	}

	for k, v := range grouped {
		result := models.ProductWithBuyer{}
		result.ProductListingId = k.ProductListingId
		result.ProductName = k.ProductName
		result.ProductPrice = k.ProductPrice
		result.ProductStock = k.ProductStock
		result.BuyerDetails = v
		results.CustomerDetail = append(results.CustomerDetail, result)
	}

	return results, err
}
