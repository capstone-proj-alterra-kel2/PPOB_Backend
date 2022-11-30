package providers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type providerUsecase struct {
	providerRepository Repository
}

func NewProviderUseCase(pr Repository) Usecase {
	return &providerUsecase{
		providerRepository: pr,
	}
}

func (pu *providerUsecase) GetAll() []Domain {
	return pu.providerRepository.GetAll()
}
func (pu *providerUsecase) Create(providerDomain *Domain) Domain {
	return pu.providerRepository.Create(providerDomain)
}
func (pu *providerUsecase) GetOne(provider_id int) Domain {
	return pu.providerRepository.GetOne(provider_id)
}
func (pu *providerUsecase) GetByPhone(phone_number string, product_type_id int) Domain {
	var provider string
	sliced_phone_number := phone_number[:4]
	fmt.Println(sliced_phone_number)
	prefixJSON, err := os.Open("prefix.json")

	if err != nil {
		fmt.Println(err)
	}

	defer prefixJSON.Close()

	var prefixes []Prefix

	byteValuePrefix, _ := ioutil.ReadAll(prefixJSON)
	json.Unmarshal(byteValuePrefix, &prefixes)

	for i := 0; i < len(prefixes); i++ {
		fmt.Println(prefixes[i].Prefix)
		if prefixes[i].Prefix == sliced_phone_number {
			provider = prefixes[i].Type
		} else {
			continue
		}
	}

	return pu.providerRepository.GetByPhone(provider, product_type_id)
}
func (pu *providerUsecase) Update(providerDomain *Domain, provider_id int) Domain {
	return pu.providerRepository.Update(providerDomain, provider_id)
}
func (pu *providerUsecase) Delete(provider_id int) Domain {
	return pu.providerRepository.Delete(provider_id)
}
