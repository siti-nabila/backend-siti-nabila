package domain

type (
	Setting struct {
		SettingKey   string `json:"setting_key"`
		SettingValue string `json:"setting_value"`
	}

	PaymentRule struct {
		Ongkir                int     `json:"ongkir"`
		MinPurchaseFreeOngkir int     `json:"min_purchase_free_ongkir"`
		MinPurchaseDisc       int     `json:"min_purchase_discount"`
		DiscountValue         float64 `json:"discount"`
	}

	SettingRepository interface {
		GetSettingByKey(settingKey string) (Setting, error)
	}
)
