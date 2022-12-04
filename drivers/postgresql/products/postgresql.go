package products

import (
	"PPOB_BACKEND/businesses/products"

	"gorm.io/gorm"
)

type productRepository struct {
	conn *gorm.DB
}

func NewPostgreSQLRepository(conn *gorm.DB) products.Repository {
	return &productRepository{
		conn: conn,
	}
}

func (pr *productRepository) GetAll() *gorm.DB {
	var prod []Product

	model := pr.conn.Order("price desc").Find(&prod).Model(&prod)

	return model
}

func (pr *productRepository) Create(productDomain *products.Domain) products.Domain {
	prod := FromDomain(productDomain)
	prod.TotalPurchased = 0
	prod.IsPromoActive = false

	pr.conn.Create(&prod)
	return prod.ToDomain()
}

func (pr *productRepository) GetOne(product_id int) products.Domain {
	var prod Product

	pr.conn.First(&prod, product_id)

	return prod.ToDomain()
}

func (pr *productRepository) Update(productDomain *products.Domain, product_id int) products.Domain {
	prod := FromDomain(productDomain)

	pr.conn.Model(&prod).Where("id = ?", product_id).Updates(
		Product{
			Name:                  productDomain.Name,
			Category:              productDomain.Category,
			Description:           productDomain.Description,
			Price:                 productDomain.Price,
			ProviderID:            productDomain.ProviderID,
			Status:                productDomain.Status,
			AdditionalInformation: productDomain.AdditionalInformation,
			IsAvailable:           productDomain.IsAvailable,
			IsPromo:               productDomain.IsPromo,
			IsPromoActive:         prod.IsPromoActive,
			Discount:              productDomain.Discount,
			PromoStartDate:        productDomain.PromoStartDate,
			PromoEndDate:          productDomain.PromoEndDate,
		},
	)

	return prod.ToDomain()
}

func (pr *productRepository) Delete(product_id int) products.Domain {
	var prod Product

	pr.conn.Unscoped().Where("id = ?", product_id).Delete(&prod)

	return prod.ToDomain()
}
