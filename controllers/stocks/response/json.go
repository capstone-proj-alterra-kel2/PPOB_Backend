package response

import (
	"PPOB_BACKEND/businesses/stocks"
	resproduct "PPOB_BACKEND/controllers/products/response"
)

type Stock struct {
	ID       uint                 `json:"id" gorm:"primaryKey"`
	Products []resproduct.Product `json:"products"`
	Quantity int                  `json:"quantity"`
}

func FromDomain(domain stocks.Domain) Stock {
	var productsData []resproduct.Product
	productFromDomain := domain.Products

	if len(productFromDomain) != 0 {
		for _, product := range productFromDomain {
			productsData = append(productsData, resproduct.Product{
				ID:                    product.ID,
				Name:                  product.Name,
				Category:              product.Category,
				Description:           product.Description,
				Price:                 product.Price,
				ProviderID:            product.ProviderID,
				StockID:               product.StockID,
				TotalPurchased:        product.TotalPurchased,
				AdditionalInformation: product.AdditionalInformation,
			})
		}
	}

	return Stock{
		ID:       domain.ID,
		Products: productsData,
		Quantity: domain.Quantity,
	}
}
