package category

import (
	"PPOB_BACKEND/businesses/category"
	"PPOB_BACKEND/businesses/producttypes"
	recProductType "PPOB_BACKEND/drivers/postgresql/producttypes"
	"time"

	"gorm.io/gorm"
)

type Category struct {
	ID          uint                         `json:"id" gorm:"size:100;primaryKey"`
	Name        string                       `json:"name"`
	ProductType []recProductType.ProductType `json:"product_type" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;`
	CreatedAt   time.Time                    `json:"created_at"`
	UpdatedAt   time.Time                    `json:"updated_at"`
	DeletedAt   gorm.DeletedAt               `json:"deleted_at"`
}

func FromDomain(domain *category.Domain) *Category {
	var productTypeData []recProductType.ProductType
	productTypeFromDomain := domain.ProductType

	for _, productType := range productTypeFromDomain {
		productTypeData = append(productTypeData, recProductType.ProductType{
			ID:         productType.ID,
			Name:       productType.Name,
			Image:      productType.Image,
			CategoryID: productType.CategoryID,
			CreatedAt:  productType.CreatedAt,
			UpdatedAt:  productType.UpdateAt,
			DeletedAt:  productType.DeletedAt,
		})
	}

	return &Category{
		ID:          domain.ID,
		Name:        domain.Name,
		ProductType: productTypeData,
		CreatedAt:   domain.CreatedAt,
		UpdatedAt:   domain.UpdatedAt,
		DeletedAt:   domain.DeeletedAt,
	}
}

func (recCat *Category) ToDomain() category.Domain {

	var productTypeFromDomain []producttypes.Domain
	for _, productType := range recCat.ProductType {
		productTypeFromDomain = append(productTypeFromDomain, productType.ToDomain())
	}

	return category.Domain{
		ID:          recCat.ID,
		Name:        recCat.Name,
		ProductType: productTypeFromDomain,
		CreatedAt:   recCat.CreatedAt,
		UpdatedAt:   recCat.UpdatedAt,
		DeeletedAt:  recCat.DeletedAt,
	}
}
