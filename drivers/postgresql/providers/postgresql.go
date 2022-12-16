package providers

import (
	"PPOB_BACKEND/businesses/providers"

	"gorm.io/gorm"
)

type providerRepository struct {
	conn *gorm.DB
}

func NewPostgreSQLRepository(conn *gorm.DB) providers.Repository {
	return &providerRepository{
		conn: conn,
	}
}

func (pr *providerRepository) GetAll(product_type_id int) ([]providers.Domain, bool) {
	var providersData []Provider
	providerDomain := []providers.Domain{}

	var productTypeID int

	pr.conn.Raw("SELECT pt.id FROM providers INNER JOIN product_types AS pt ON pt.id = ?", product_type_id).Scan(&productTypeID)

	if productTypeID == 0 {
		return providerDomain, false
	}

	pr.conn.Where("product_type_id = ?", product_type_id).Find(&providersData)

	for _, provider := range providersData {
		providerDomain = append(providerDomain, provider.ToDomain())
	}

	return providerDomain, true
}

func (pr *providerRepository) Create(providerDomain *providers.Domain, product_type_id int) (providers.Domain, bool, bool) {
	providerData := FromDomain(providerDomain)

	checkName := providerData.Name

	var productTypeID int
	var isProductTypeFound bool
	var isNameDuplicated bool

	pr.conn.Raw("SELECT pt.id FROM providers RIGHT JOIN product_types AS pt ON pt.id = ?", product_type_id).Scan(&productTypeID)

	if productTypeID == 0 {
		isProductTypeFound = false
	} else {
		isProductTypeFound = true
	}

	if checkProvider := pr.conn.First(&providerData, "name = ? AND product_type_id = ?", checkName, product_type_id).Error; checkProvider != gorm.ErrRecordNotFound {
		isNameDuplicated = true
	}

	if !isProductTypeFound {
		return providerData.ToDomain(), false, false
	}

	if isNameDuplicated {
		return providerData.ToDomain(), true, true
	}

	if !isNameDuplicated && isProductTypeFound {
		pr.conn.Create(&providerData)
		return providerData.ToDomain(), true, false
	}

	return providerData.ToDomain(), false, false
}

func (pr *providerRepository) GetOne(provider_id int, product_type_id int) (providers.Domain, bool, bool) {
	var providerData Provider
	var productTypeID int

	var isProviderFound bool
	var isProductTypeFound bool

	pr.conn.Raw("SELECT pt.id FROM providers INNER JOIN product_types AS pt ON pt.id = ?", product_type_id).Scan(&productTypeID)

	if productTypeID != 0 {
		isProductTypeFound = true
	}

	if checkProvider := pr.conn.First(&providerData, "id = ? AND product_type_id = ?", provider_id, product_type_id).Error; checkProvider != gorm.ErrRecordNotFound {
		isProviderFound = true
	}

	if isProviderFound && isProductTypeFound {
		pr.conn.Preload("Products").Find(&providerData).Where("id = ? AND product_type_id = ?", provider_id, product_type_id)

		return providerData.ToDomain(), true, true
	}

	if !isProductTypeFound {
		return providerData.ToDomain(), false, false
	}

	if !isProviderFound {
		return providerData.ToDomain(), true, false
	}

	return providerData.ToDomain(), false, false

}

func (pr *providerRepository) GetByPhone(provider string, product_type_id int) (providers.Domain, bool) {
	var providerData Provider
	var productTypeID int

	pr.conn.Raw("SELECT pt.id FROM providers INNER JOIN product_types AS pt ON pt.id = ?", product_type_id).Scan(&productTypeID)

	if productTypeID == 0 {
		return providerData.ToDomain(), false
	}

	pr.conn.Where("product_type_id = ? AND name = ?", product_type_id, provider).Preload("Products", func(db *gorm.DB) *gorm.DB {
		return db.Order("products.price")
	}).Find(&providerData)

	return providerData.ToDomain(), true
}

func (pr *providerRepository) Update(providerDomain *providers.Domain, provider_id int) (providers.Domain, error) {
	providerData := FromDomain(providerDomain)
	var prov Provider

	err := pr.conn.First(&prov, provider_id).Error

	if err != nil {
		return providerData.ToDomain(), err
	}

	pr.conn.Model(&providerData).Where("id = ?", provider_id).Updates(
		Provider{
			Name:  providerDomain.Name,
			Image: providerDomain.Image,
		},
	)

	if len(providerData.Image) != 0 {
		prov.Image = providerData.Image
	}
	prov.Name = providerData.Name

	return prov.ToDomain(), nil
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
