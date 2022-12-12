package producttypes

type productTypeUsecase struct {
	productTypeRepository Repository
}

func NewProductTypeUseCase(ptr Repository) Usecase {
	return &productTypeUsecase{
		productTypeRepository: ptr,
	}
}

func (ptu *productTypeUsecase) GetAll() []Domain {
	return ptu.productTypeRepository.GetAll()
}

func (ptu *productTypeUsecase) Create(productTypeDomain *Domain) Domain {
	return ptu.productTypeRepository.Create(productTypeDomain)
}

func (ptu *productTypeUsecase) GetOne(product_type_id int) Domain {
	return ptu.productTypeRepository.GetOne(product_type_id)
}

func (ptu *productTypeUsecase) Update(productTypeDomain *Domain, product_type_id int) (Domain, error) {
	return ptu.productTypeRepository.Update(productTypeDomain, product_type_id)
}

func (ptu *productTypeUsecase) Delete(product_type_id int) (Domain, error) {
	return ptu.productTypeRepository.Delete(product_type_id)
}
