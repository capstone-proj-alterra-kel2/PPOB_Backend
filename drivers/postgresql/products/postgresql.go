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

func (pr *productRepository) GetAll() []products.Domain {
	var prod []Product

	pr.conn.Find(&prod)

	productDomain := []products.Domain{}

	for _, products := range prod {
		productDomain = append(productDomain, products.ToDomain())
	}

	return productDomain
}

func (pr *productRepository) Create(productDomain *products.Domain) products.Domain {
	prod := FromDomain(productDomain)
	prod.TotalPurchased = 0

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
			StockID:               productDomain.StockID,
			AdditionalInformation: productDomain.AdditionalInformation,
		},
	)

	return prod.ToDomain()
}

func (pr *productRepository) Delete(product_id int) products.Domain {
	var prod Product

	pr.conn.Unscoped().Where("id = ?", product_id).Delete(&prod)

	return prod.ToDomain()
}
