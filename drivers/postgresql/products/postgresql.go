package products

import (
	"PPOB_BACKEND/businesses/products"
	"strings"

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

func (pr *productRepository) GetAll(Page int, Size int, Sort string, Search string) (*gorm.DB, []products.Domain) {
	var prod []Product
	var sort string
	var search string
	var model *gorm.DB

	if strings.HasPrefix(Sort, "-") {
		sort = Sort[1:] + " DESC"
	} else {
		sort = Sort[0:] + " ASC"
	}

	model = pr.conn.Order(sort).Find(&prod).Model(&prod)

	if Search != "" {
		search = "%" + Search + "%"
		model = pr.conn.Order(sort).Model(&prod).Where("name LIKE ?", search)
	}

	pr.conn.Offset(Page).Limit(Size).Order(sort).Find(&prod)
	productDomain := []products.Domain{}
	for _, product := range prod {
		productDomain = append(productDomain, product.ToDomain())
	}

	return model, productDomain
}

func (pr *productRepository) Create(productDomain *products.Domain) products.Domain {
	prod := FromDomain(productDomain)

	pr.conn.Create(&prod)
	return prod.ToDomain()
}

func (pr *productRepository) GetOne(product_id int) products.Domain {
	var prod Product

	pr.conn.First(&prod, product_id)

	return prod.ToDomain()
}

func (pr *productRepository) UpdateData(productDomain *products.UpdateDataDomain, product_id int) (products.Domain, error) {
	prod := FromUpdatedDomain(productDomain)

	err := pr.conn.First(&prod, product_id).Error

	if err != nil {
		return prod.ToDomain(), err
	}

	pr.conn.Model(&prod).Where("id = ?", product_id).Updates(
		Product{
			PriceStatus:    productDomain.PriceStatus,
			IsPromoActive:  productDomain.IsPromoActive,
			Discount:       productDomain.Discount,
			PromoStartDate: productDomain.PromoStartDate,
			PromoEndDate:   productDomain.PromoEndDate,
		},
	)

	return prod.ToDomain(), nil
}

func (pr *productRepository) UpdatePromo(productDomain *products.Domain) products.Domain {
	prod := FromDomain(productDomain)

	pr.conn.Model(&prod).Where("id = ?", productDomain.ID).Updates(
		Product{
			ID:             productDomain.ID,
			Name:           productDomain.Name,
			Description:    productDomain.Description,
			Price:          productDomain.Price,
			ProviderID:     productDomain.ProviderID,
			Status:         productDomain.Status,
			IsAvailable:    productDomain.IsAvailable,
			PriceStatus:    productDomain.PriceStatus,
			IsPromoActive:  productDomain.IsPromoActive,
			Discount:       productDomain.Discount,
			PromoStartDate: productDomain.PromoStartDate,
			PromoEndDate:   productDomain.PromoEndDate,
		},
	)

	return prod.ToDomain()
}

func (pr *productRepository) UpdateStockStatus(productDomain *products.UpdateStockStatusDomain, product_id int) products.Domain {
	prod := FromUpdatedStockStatusDomain(productDomain)

	pr.conn.Model(&prod).Where("id = ?", product_id).Updates(
		Product{
			Status:         productDomain.Status,
			TotalPurchased: productDomain.TotalPurchased,
			Stock:          productDomain.Stock,
			IsAvailable:    productDomain.IsAvailable,
		})
	return prod.ToDomain()
}

func (pr *productRepository) Delete(product_id int) (products.Domain, error) {
	var prod Product

	err := pr.conn.First(&prod, product_id).Error

	if err != nil {
		return prod.ToDomain(), err
	}

	pr.conn.Unscoped().Where("id = ?", product_id).Delete(&prod)

	return prod.ToDomain(), nil
}
