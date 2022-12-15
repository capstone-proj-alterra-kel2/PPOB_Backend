package products

import (
	"time"

	"gorm.io/gorm"
)

type productUsecase struct {
	productRepository Repository
}

func NewProductUseCase(pr Repository) Usecase {
	return &productUsecase{
		productRepository: pr,
	}
}

func (pu *productUsecase) GetAll(Page int, Size int, Sort string, Search string) (*gorm.DB, []Domain) {
	return pu.productRepository.GetAll(Page, Size, Sort, Search)
}

func (pu *productUsecase) Create(productDomain *Domain) Domain {
	return pu.productRepository.Create(productDomain)
}

func (pu *productUsecase) GetOne(product_id int) Domain {
	var parsedStartDate time.Time
	var parsedEndDate time.Time

	layoutFormat := "2006-01-02"

	currentDate := time.Now()
	formatDate := currentDate.Format("2006-01-02")

	parsedCurrentDate, _ := time.Parse(layoutFormat, formatDate)

	result := pu.productRepository.GetOne(product_id)

	if *result.IsPromo {
		parsedStartDate, _ = time.Parse(layoutFormat, result.PromoStartDate)
		parsedEndDate, _ = time.Parse(layoutFormat, result.PromoEndDate)

		if parsedCurrentDate.Before(parsedEndDate) && parsedCurrentDate.After(parsedStartDate) {
			*result.IsPromoActive = true
		} else {
			*result.IsPromoActive = false
		}
	}

	return result
}

func (pu *productUsecase) UpdateData(productDomain *UpdateDataDomain, product_id int) (Domain, error) {
	return pu.productRepository.UpdateData(productDomain, product_id)
}

func (pu *productUsecase) UpdatePromo(productDomain *Domain) Domain {
	return pu.productRepository.UpdatePromo(productDomain)
}

func (pu *productUsecase) UpdateStockStatus(productDomain *UpdateStockStatusDomain, product_id int) Domain {
	return pu.productRepository.UpdateStockStatus(productDomain, product_id)
}

func (pu *productUsecase) Delete(product_id int) (Domain, error) {
	return pu.productRepository.Delete(product_id)
}
