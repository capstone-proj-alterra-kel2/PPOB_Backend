package response

import (
	"PPOB_BACKEND/businesses/category"
	"PPOB_BACKEND/businesses/producttypes"
	resProductType "PPOB_BACKEND/controllers/producttypes/response"
)

type Category struct {
	ID          uint                         `json:"id"`
	Name        string                       `json:"name"`
	ProductType []resProductType.ProductType `json:"product_type"`
}

func FromDomain(domain category.Domain) Category {
	var productTypeData []resProductType.ProductType
	productTypeFromDomain := domain.ProductType

	for _, productType := range productTypeFromDomain {
		productTypeData = append(productTypeData, resProductType.FromDomain(producttypes.Domain{
			ID:         productType.ID,
			Name:       productType.Name,
			Image:      productType.Image,
			CategoryID: productType.CategoryID,
			CreatedAt:  productType.CreatedAt,
			UpdateAt:   productType.UpdateAt,
			DeletedAt:  productType.DeletedAt,
		}))
	}

	return Category{
		ID:          domain.ID,
		Name:        domain.Name,
		ProductType: productTypeData,
	}
}
