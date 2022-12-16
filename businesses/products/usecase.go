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
	var parsedStartDate time.Time
	var parsedEndDate time.Time

	model, result := pu.productRepository.GetAll(Page, Size, Sort, Search)

	layoutFormat := "2006-01-02"
	layoutFormatCurrent := "2006-01-02 15:04:05"

	currentDate := time.Now()
	formatDate := currentDate.Format("2006-01-02 15:04:05")

	parsedCurrentDate, _ := time.Parse(layoutFormatCurrent, formatDate)
	updatedDataDomain := []Domain{}

	for _, value := range result {
		if value.PriceStatus == "promo" {
			parsedStartDate, _ = time.Parse(layoutFormat, value.PromoStartDate)
			parsedEndDate, _ = time.Parse(layoutFormat, value.PromoEndDate)

			if parsedCurrentDate.Before(parsedEndDate) && parsedCurrentDate.After(parsedStartDate) {
				*value.IsPromoActive = true
			} else {
				*value.IsPromoActive = false
			}
		}

		if *value.Stock > 0 {
			*value.IsAvailable = true
		}
		updatedDataDomain = append(updatedDataDomain, value)
	}

	return model, updatedDataDomain
}

func (pu *productUsecase) GetAllForUser() []Domain {
	var parsedStartDate time.Time
	var parsedEndDate time.Time

	result := pu.productRepository.GetAllForUser()

	layoutFormat := "2006-01-02"
	layoutFormatCurrent := "2006-01-02 15:04:05"

	currentDate := time.Now()
	formatDate := currentDate.Format("2006-01-02 15:04:05")

	parsedCurrentDate, _ := time.Parse(layoutFormatCurrent, formatDate)
	updatedDataDomain := []Domain{}

	for _, value := range result {
		if value.PriceStatus == "promo" {
			parsedStartDate, _ = time.Parse(layoutFormat, value.PromoStartDate)
			parsedEndDate, _ = time.Parse(layoutFormat, value.PromoEndDate)

			if parsedCurrentDate.Before(parsedEndDate) && parsedCurrentDate.After(parsedStartDate) {
				*value.IsPromoActive = true
			} else {
				*value.IsPromoActive = false
			}
		}

		if *value.Stock > 0 {
			*value.IsAvailable = true
		}
		updatedDataDomain = append(updatedDataDomain, value)
	}

	return updatedDataDomain
}

func (pu *productUsecase) Create(productDomain *Domain) (Domain, bool) {
	var isDateValid bool

	layoutFormat := "2006-01-02"

	_, errOne := time.Parse(layoutFormat, productDomain.PromoStartDate)
	_, errTwo := time.Parse(layoutFormat, productDomain.PromoEndDate)

	if productDomain.PriceStatus == "promo" {
		if errOne == nil && errTwo == nil {
			isDateValid = true
		} else {
			isDateValid = false
		}
	} else {
		isDateValid = true
	}

	res := pu.productRepository.Create(productDomain)

	return res, isDateValid
}

func (pu *productUsecase) GetOne(product_id int) (Domain, error) {
	var parsedStartDate time.Time
	var parsedEndDate time.Time

	layoutFormat := "2006-01-02"
	layoutFormatCurrent := "2006-01-02 15:04:05"

	currentDate := time.Now()
	formatDate := currentDate.Format("2006-01-02 15:04:05")

	parsedCurrentDate, _ := time.Parse(layoutFormatCurrent, formatDate)

	result, err := pu.productRepository.GetOne(product_id)

	if err != nil {
		return result, err
	}

	if result.PriceStatus == "promo" {
		parsedStartDate, _ = time.Parse(layoutFormat, result.PromoStartDate)
		parsedEndDate, _ = time.Parse(layoutFormat, result.PromoEndDate)

		if parsedCurrentDate.Before(parsedEndDate) && parsedCurrentDate.After(parsedStartDate) {
			*result.IsPromoActive = true
		} else {
			*result.IsPromoActive = false
		}
	}

	if *result.Stock > 0 {
		*result.IsAvailable = true
	}

	return result, nil
}

func (pu *productUsecase) UpdateData(productDomain *UpdateDataDomain, product_id int) (Domain, error, bool) {
	var isDateValid bool

	layoutFormat := "2006-01-02"

	_, errOne := time.Parse(layoutFormat, productDomain.PromoStartDate)
	_, errTwo := time.Parse(layoutFormat, productDomain.PromoEndDate)

	if productDomain.PriceStatus == "promo" {
		if errOne == nil && errTwo == nil {
			isDateValid = true
		} else {
			isDateValid = false
		}
	} else {
		isDateValid = true
	}

	res, err := pu.productRepository.UpdateData(productDomain, product_id)

	return res, err, isDateValid
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
