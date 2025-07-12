package products

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/siti-nabila/backend-siti-nabila/internal/models"
)

func (p *productHandler) AddNewProduct(c *fiber.Ctx) error {
	var (
		request  models.InsertMerchantProductRequest
		response models.ListingProductMerchantResponse
	)
	userID := c.Locals("user_id").(int)
	log.Info("user id : ", userID)
	// panic("")
	if err := c.BodyParser(&request); err != nil {
		response.ErrorMessage = err.Error()
		return c.Status(fiber.StatusConflict).JSON(response)
	}

	request.UserId = userID
	res, err := p.productService.AddMerchantListingProduct(request)
	if err != nil {
		log.Error(err)
		response.ErrorMessage = "error when adding product"
		return c.Status(fiber.StatusUnauthorized).JSON(response)
	}
	response = res
	return c.Status(fiber.StatusOK).JSON(response)

}

func (p *productHandler) GetListingProducts(c *fiber.Ctx) error {
	var (
		response models.ListingProductMerchantResponse
	)
	userID := c.Locals("user_id").(int)
	res, err := p.productService.GetMerchantListingProducts(userID)
	if err != nil {
		log.Error(err)
		response.ErrorMessage = "error when getting product data"
		return c.Status(fiber.StatusUnauthorized).JSON(response)
	}
	response = res
	return c.Status(fiber.StatusOK).JSON(response)

}

func (p *productHandler) GetListingProductsWithBuyer(c *fiber.Ctx) error {
	var (
		response models.ListingProductMerchantWithBuyer
	)
	userID := c.Locals("user_id").(int)
	res, err := p.productService.GetProductWithBuyer(userID)
	if err != nil {
		log.Error(err)
		response.ErrorMessage = "error when getting product data"
		return c.Status(fiber.StatusUnauthorized).JSON(response)
	}
	response = res
	return c.Status(fiber.StatusOK).JSON(response)

}
