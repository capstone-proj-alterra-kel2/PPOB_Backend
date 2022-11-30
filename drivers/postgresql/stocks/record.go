package stocks

import (
	productdomain "PPOB_BACKEND/businesses/products"
	"PPOB_BACKEND/businesses/stocks"
	"PPOB_BACKEND/drivers/postgresql/products"
	"time"

	"gorm.io/gorm"
)

type Stock struct {
	ID        uint               `json:"id" gorm:"size:100;primaryKey"`
	Quantity  int                `json:"quantity"`
	Products  []products.Product `json:"products" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	CreatedAt time.Time          `json:"created_at"`
	UpdatedAt time.Time          `json:"updated_at"`
	DeletedAt gorm.DeletedAt     `json:"deleted_at"`
}

func FromDomain(domain *stocks.Domain) *Stock {
	var productData []products.Product
	productFromDomain := domain.Products

	for _, product := range productFromDomain {
		productData = append(productData, products.Product{
			ID:                    product.ID,
			Name:                  product.Name,
			Category:              product.Category,
			Description:           product.Description,
			Price:                 product.Price,
			ProviderID:            product.ProviderID,
			StockID:               product.StockID,
			TotalPurchased:        product.TotalPurchased,
			AdditionalInformation: product.AdditionalInformation,
			CreatedAt:             product.CreatedAt,
			UpdatedAt:             product.UpdatedAt,
			DeletedAt:             product.DeletedAt,
		})
	}

	return &Stock{
		ID:        domain.ID,
		Quantity:  domain.Quantity,
		Products:  productData,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdateAt,
		DeletedAt: domain.DeletedAt,
	}
}

func (recStock *Stock) ToDomain() stocks.Domain {
	var productsFromDomain []productdomain.Domain
	for _, products := range recStock.Products {
		productsFromDomain = append(productsFromDomain, products.ToDomain())
	}

	return stocks.Domain{
		ID:        recStock.ID,
		Quantity:  recStock.Quantity,
		Products:  productsFromDomain,
		CreatedAt: recStock.CreatedAt,
		UpdateAt:  recStock.UpdatedAt,
		DeletedAt: recStock.DeletedAt,
	}
}
