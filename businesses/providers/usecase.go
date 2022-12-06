package providers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
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

func (pu *providerUsecase) GetAll(product_type_id int) ([]Domain, error) {
	return pu.providerRepository.GetAll(product_type_id)
}
func (pu *providerUsecase) Create(providerDomain *Domain, product_type_id int) (Domain, error) {
	return pu.providerRepository.Create(providerDomain, product_type_id)
}
func (pu *providerUsecase) GetOne(provider_id int, product_type_id int) (Domain, error) {
	res, err := pu.providerRepository.GetOne(provider_id, product_type_id)

	return res, err
}
func (pu *providerUsecase) GetByPhone(phone_number string, product_type_id int) Domain {
	var provider string
	var prefixes []Prefix
	var parsedStartDate time.Time
	var parsedEndDate time.Time
	var parsedCurrentTime time.Time
	var providerResult Domain

	sliced_phone_number := phone_number[:4]
	prefixJSON, err := os.Open("prefix.json")

	if err != nil {
		fmt.Println(err)
	}

	defer prefixJSON.Close()

	byteValuePrefix, _ := ioutil.ReadAll(prefixJSON)
	json.Unmarshal(byteValuePrefix, &prefixes)

	for i := 0; i < len(prefixes); i++ {
		fmt.Println(prefixes[i].Prefix)
		if prefixes[i].Prefix == sliced_phone_number {
			provider = prefixes[i].Provider
		} else {
			continue
		}
	}

	result := pu.providerRepository.GetByPhone(provider, product_type_id)

	layoutFormat := "2006-01-02 15:04:05"

	currentTime := time.Now().Local().Format(layoutFormat)
	parsedCurrentTime, _ = time.Parse(layoutFormat, currentTime)

	for _, prodProvider := range result.Products {
		if prodProvider.IsPromo {
			parsedStartDate, _ = time.Parse(layoutFormat, prodProvider.PromoStartDate)
			parsedEndDate, _ = time.Parse(layoutFormat, prodProvider.PromoEndDate)

			if parsedCurrentTime.Before(parsedEndDate) && parsedCurrentTime.After(parsedStartDate) {
				prodProvider.IsPromoActive = true
			} else {
				prodProvider.IsPromoActive = false
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

	return providerResult
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
