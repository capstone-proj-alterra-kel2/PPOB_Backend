package products

import "gorm.io/gorm"

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
	return pu.productRepository.GetOne(product_id)
}

func (pu *productUsecase) Update(productDomain *Domain, product_id int) Domain {
	return pu.productRepository.Update(productDomain, product_id)
}

func (pu *productUsecase) Delete(product_id int) Domain {
	return pu.productRepository.Delete(product_id)
}
