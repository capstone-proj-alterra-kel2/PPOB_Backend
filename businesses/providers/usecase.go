package providers

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"time"
)

type providerUsecase struct {
	providerRepository Repository
}

func NewProviderUseCase(providerrepo Repository) Usecase {
	return &providerUsecase{
		providerRepository: providerrepo,
	}
}

func (pu *providerUsecase) GetAll(product_type_id int) ([]Domain, bool) {
	return pu.providerRepository.GetAll(product_type_id)
}
func (pu *providerUsecase) Create(providerDomain *Domain, product_type_id int) (Domain, bool, bool) {
	return pu.providerRepository.Create(providerDomain, product_type_id)
}
func (pu *providerUsecase) GetOne(provider_id int, product_type_id int) (Domain, bool, bool) {
	return pu.providerRepository.GetOne(provider_id, product_type_id)
}
func (pu *providerUsecase) GetByPhone(phone_number string, product_type_id int) (Domain, bool) {
	var provider string
	var prefixes Prefixes
	var parsedStartDate time.Time
	var parsedEndDate time.Time
	var providerResult Domain

	sliced_phone_number := phone_number[:4]
	prefixJSON, err := os.Open("prefix.json")

	if err != nil {
		fmt.Println(err)
	}

	defer prefixJSON.Close()

	byteValuePrefix, _ := io.ReadAll(prefixJSON)
	json.Unmarshal(byteValuePrefix, &prefixes)

	for i := 0; i < len(prefixes.Prefixes); i++ {
		if prefixes.Prefixes[i].Prefix == sliced_phone_number {
			provider = prefixes.Prefixes[i].Provider
		} else {
			continue
		}
	}

	result, isProductTypeFound := pu.providerRepository.GetByPhone(provider, product_type_id)

	if !isProductTypeFound {
		return providerResult, false
	}

	layoutFormat := "2006-01-02"

	currentDate := time.Now()
	formatDate := currentDate.Format("2006-01-02")

	parsedCurrentDate, _ := time.Parse(layoutFormat, formatDate)

	for _, prodProvider := range result.Products {
		if prodProvider.PriceStatus == "promo" {
			parsedStartDate, _ = time.Parse(layoutFormat, prodProvider.PromoStartDate)
			parsedEndDate, _ = time.Parse(layoutFormat, prodProvider.PromoEndDate)

			if parsedCurrentDate.Before(parsedEndDate) && parsedCurrentDate.After(parsedStartDate) {
				*prodProvider.IsPromoActive = true
			} else {
				*prodProvider.IsPromoActive = false
			}
		}

		providerResult.Products = append(providerResult.Products, prodProvider)
	}

	providerResult.ID = result.ID
	providerResult.Image = result.Image
	providerResult.Name = result.Name
	providerResult.CreatedAt = result.CreatedAt
	providerResult.UpdateAt = result.UpdateAt
	providerResult.DeletedAt = result.DeletedAt

	return providerResult, true
}

func (pu *providerUsecase) UpdateCheck(providerDomain *ProviderDomain, provider_id int) Domain {
	return pu.providerRepository.UpdateCheck(providerDomain, provider_id)
}

func (pu *providerUsecase) Update(providerDomain *Domain, provider_id int) Domain {
	return pu.providerRepository.Update(providerDomain, provider_id)
}
func (pu *providerUsecase) Delete(provider_id int) Domain {
	return pu.providerRepository.Delete(provider_id)
}
