package products

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/siti-nabila/backend-siti-nabila/internal/models"
)

func (p *productHandler) GetAllProducts(c *fiber.Ctx) error {
	var (
		response models.ListingProductCustomerResponse
	)

	res, err := p.productService.GetAllProduct()
	if err != nil {
		log.Error(err)
		response.ErrorMessage = "error when getting product data"
		return c.Status(fiber.StatusUnauthorized).JSON(response)
	}
	response = res
	return c.Status(fiber.StatusOK).JSON(response)

}

func (p *productHandler) GetPurchasedItemHistory(c *fiber.Ctx) error {
	var (
		response models.ListingProductCustomerResponse
	)
	userID := c.Locals("user_id").(int)

	res, err := p.productService.GetHistoryItems(userID)
	if err != nil {
		log.Error(err)
		response.ErrorMessage = "error when getting history data"
		return c.Status(fiber.StatusUnauthorized).JSON(response)
	}
	response = res
	return c.Status(fiber.StatusOK).JSON(response)

}

func (p *productHandler) BuyProduct(c *fiber.Ctx) error {
	var (
		request  models.InsertCustomerProductRequest
		response models.ListingProductCustomerResponse
	)
	userID := c.Locals("user_id").(int)
	if err := c.BodyParser(&request); err != nil {
		response.ErrorMessage = err.Error()
		return c.Status(fiber.StatusConflict).JSON(response)
	}
	request.UserId = userID
	res, err := p.productService.BuyProduct(request)
	if err != nil {
		log.Error(err)
		response.ErrorMessage = "error when buying product"
		return c.Status(fiber.StatusUnauthorized).JSON(response)
	}
	response = res

	return c.Status(fiber.StatusOK).JSON(response)
}
