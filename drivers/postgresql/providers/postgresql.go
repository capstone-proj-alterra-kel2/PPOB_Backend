package providers

import (
	"PPOB_BACKEND/businesses/providers"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type providerRepository struct {
	conn *gorm.DB
}

func NewPostgreSQLRepository(conn *gorm.DB) providers.Repository {
	return &providerRepository{
		conn: conn,
	}
}

func (pr *providerRepository) GetAll(product_type_id int) ([]providers.Domain, error) {
	var providersData []Provider
	providerDomain := []providers.Domain{}

	err := pr.conn.First(&providersData, "product_type_id = ?", product_type_id).Error

	if err != nil {
		return providerDomain, err
	}

	pr.conn.Preload(clause.Associations).Find(&providersData).Where("product_type_id = ?", product_type_id)

	for _, provider := range providersData {
		providerDomain = append(providerDomain, provider.ToDomain())
	}

	return providerDomain, nil
}

func (pr *providerRepository) Create(providerDomain *providers.Domain, product_type_id int) (providers.Domain, bool) {
	providerData := FromDomain(providerDomain)

	checkName := providerData.Name

	if checkProvider := pr.conn.First(&providerData, "name = ? AND product_type_id = ?", checkName, product_type_id).Error; checkProvider != gorm.ErrRecordNotFound {
		return providerData.ToDomain(), true
	}

	pr.conn.Create(&providerData)

	return providerData.ToDomain(), false
}

func (pr *providerRepository) GetOne(provider_id int, product_type_id int) (providers.Domain, bool, bool) {
	var providerData Provider
	var isProviderFound bool
	var isProductTypeFound bool

	if checkProductType := pr.conn.First(&providerData, "product_type_id = ?", product_type_id).Error; checkProductType != gorm.ErrRecordNotFound {
		isProductTypeFound = true
	}

	if checkProvider := pr.conn.First(&providerData, "id = ?", provider_id).Error; checkProvider != gorm.ErrRecordNotFound {
		isProviderFound = true
	}

	if isProductTypeFound && isProviderFound {
		pr.conn.Preload("Products").Find(&providerData).Where("id = ? AND product_type_id = ?", provider_id, product_type_id)

		return providerData.ToDomain(), true, true
	}

	if !isProductTypeFound {
		return providerData.ToDomain(), false, true
	}

	if !isProviderFound {
		return providerData.ToDomain(), true, false
	}

	return providerData.ToDomain(), false, false

}

func (pr *providerRepository) GetByPhone(provider string, product_type_id int) providers.Domain {
	var providerData Provider

	pr.conn.Preload("Products", func(db *gorm.DB) *gorm.DB {
		return db.Order("products.price")
	}).First(&providerData).Where("product_type_id = ? AND name = ?", product_type_id, provider)
	return providerData.ToDomain()
}

func (pr *providerRepository) Update(providerDomain *providers.Domain, provider_id int) providers.Domain {
	providerData := FromDomain(providerDomain)

	pr.conn.Model(&providerData).Where("id = ?", provider_id).Updates(
		Provider{
			Name:  providerDomain.Name,
			Image: providerDomain.Image,
		},
	)

	return providerData.ToDomain()
}

func (pr *providerRepository) UpdateCheck(providerDomain *providers.ProviderDomain, provider_id int) providers.Domain {
	updatedProviderData := FromDomainUpdate(providerDomain)

	pr.conn.Session(&gorm.Session{FullSaveAssociations: true}).Updates(&updatedProviderData)

	return updatedProviderData.ToDomain()
}

func (pr *providerRepository) Delete(provider_id int) providers.Domain {
	var providerData Provider

	pr.conn.Unscoped().Where("id = ?", provider_id).Delete(&providerData)
	return providerData.ToDomain()
}
