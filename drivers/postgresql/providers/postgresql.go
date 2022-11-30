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

func (pr *providerRepository) GetAll() []providers.Domain {
	var providersData []Provider

	pr.conn.Preload("Products").Find(&providersData)

	providerDomain := []providers.Domain{}
	for _, provider := range providersData {
		providerDomain = append(providerDomain, provider.ToDomain())
	}

	return providerDomain
}

func (pr *providerRepository) Create(providerDomain *providers.Domain) providers.Domain {
	providerData := FromDomain(providerDomain)

	pr.conn.Create(&providerData)
	return providerData.ToDomain()
}

func (pr *providerRepository) GetOne(provider_id int) providers.Domain {
	var providerData Provider

	pr.conn.Preload("Products").First(&providerData, provider_id)
	return providerData.ToDomain()
}

func (pr *providerRepository) GetByPhone(provider string, product_type_id int) providers.Domain {
	var providerData Provider

	pr.conn.Preload("Products").First(&providerData).Where("product_type_id = ? AND name = ?", product_type_id, provider)
	return providerData.ToDomain()
}

func (pr *providerRepository) Update(providerDomain *providers.Domain, provider_id int) providers.Domain {
	providerData := FromDomain(providerDomain)

	pr.conn.Model(&providerData).Where("id = ?", provider_id).Updates(
		Provider{
			Name: providerDomain.Name,
		},
	)

	return providerData.ToDomain()
}

func (pr *providerRepository) Delete(provider_id int) providers.Domain {
	var providerData Provider

	pr.conn.Unscoped().Where("id = ?", provider_id).Delete(&providerData)
	return providerData.ToDomain()
}
