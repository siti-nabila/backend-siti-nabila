package models

type (
	InsertMerchantProductRequest struct {
		ProductName  string `json:"product_name"`
		ProductPrice int    `json:"product_price"`
		ProductStock int    `json:"product_stock"`
		UserId       int    `json:"-"`
	}
	InsertCustomerProductRequest struct {
		ProductListingId int `json:"product_listing_id"`
		ProductQty       int `json:"product_quantity"`
		UserId           int `json:"-"`
	}

	ListingProductCustomerResponse struct {
		ErrorMessage string            `json:"error,omitempty"`
		Products     []ProductCustomer `json:"listing_product"`
	}

	ListingProductMerchantResponse struct {
		ErrorMessage string    `json:"error,omitempty"`
		Products     []Product `json:"listing_product"`
	}
	ListingProductMerchantWithBuyer struct {
		ErrorMessage   string             `json:"error,omitempty"`
		CustomerDetail []ProductWithBuyer `json:"listing_product"`
	}
	ProductWithBuyer struct {
		ProductListingId int            `json:"product_listing_id"`
		ProductName      string         `json:"product_name"`
		ProductPrice     int            `json:"product_price"`
		ProductStock     int            `json:"product_stock"`
		BuyerDetails     []ProductBuyer `json:"customers"`
	}
	ProductBuyer struct {
		BuyerEmail             string `json:"customer_email"`
		ProductQty             int    `json:"product_quantity,omitempty"`
		ProductSubTotalAmount  int    `json:"product_sub_total"`
		ProductOngkir          int    `json:"product_ongkir"`
		ProductDiscountAmount  int    `json:"product_discount_amount"`
		ProdcutTotalPaidAmount int    `json:"product_total_paid"`
	}
	Product struct {
		ProductListingId int    `json:"product_listing_id"`
		ProductName      string `json:"product_name"`
		ProductPrice     int    `json:"product_price"`
		ProductStock     int    `json:"product_stock"`
		BuyerEmail       string `json:"customer_email,omitempty"`
	}

	ProductCustomer struct {
		ProductListingId       int    `json:"product_listing_id"`
		ProductName            string `json:"product_name"`
		ProductPrice           int    `json:"product_price"`
		ProductQty             int    `json:"product_quantity,omitempty"`
		ProductStock           int    `json:"product_stock,omitempty"`
		ProductOngkir          int    `json:"product_ongkir"`
		ProductDiscountAmount  int    `json:"product_discount_amount"`
		ProdcutTotalPaidAmount int    `json:"product_total_paid"`
		MerchantEmail          string `json:"merchant_email"`
	}
)
