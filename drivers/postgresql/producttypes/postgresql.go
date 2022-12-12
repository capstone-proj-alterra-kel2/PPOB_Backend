package producttypes

import (
	"PPOB_BACKEND/businesses/producttypes"

	"gorm.io/gorm"
)

type productTypeRepository struct {
	conn *gorm.DB
}

func NewPostgreSQLRepository(conn *gorm.DB) producttypes.Repository {
	return &productTypeRepository{
		conn: conn,
	}
}

func (ptr *productTypeRepository) GetAll() []producttypes.Domain {
	var prodtypes []ProductType

	ptr.conn.Preload("Providers").Find(&prodtypes)

	productTypeDomain := []producttypes.Domain{}

	for _, productype := range prodtypes {
		productTypeDomain = append(productTypeDomain, productype.ToDomain())
	}
	return productTypeDomain
}

func (ptr *productTypeRepository) Create(productTypeDomain *producttypes.Domain) producttypes.Domain {
	prodtype := FromDomain(productTypeDomain)

	ptr.conn.Create(&prodtype)

	return prodtype.ToDomain()
}
func (ptr *productTypeRepository) GetOne(product_type_id int) producttypes.Domain {
	var prodtype ProductType

	ptr.conn.Preload("Providers").First(&prodtype, product_type_id)
	return prodtype.ToDomain()
}

func (ptr *productTypeRepository) Update(productTypeDomain *producttypes.Domain, product_type_id int) (producttypes.Domain, error) {
	prodtype := FromDomain(productTypeDomain)

	err := ptr.conn.First(&prodtype, product_type_id).Error

	if err != nil {
		return prodtype.ToDomain(), err
	}

	ptr.conn.Model(&prodtype).Where("id = ?", product_type_id).Updates(
		ProductType{
			Name:  productTypeDomain.Name,
			Image: productTypeDomain.Image,
		},
	)

	return prodtype.ToDomain(), nil
}

func (ptr *productTypeRepository) Delete(product_type_id int) (producttypes.Domain, error) {
	var prodtype ProductType

	err := ptr.conn.First(&prodtype, product_type_id).Error

	if err != nil {
		return prodtype.ToDomain(), err
	}

	ptr.conn.Unscoped().Where("id = ?", product_type_id).Delete(&prodtype)
	return prodtype.ToDomain(), nil
}
