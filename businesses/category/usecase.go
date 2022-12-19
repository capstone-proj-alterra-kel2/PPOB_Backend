package category

type categoryUsecase struct {
	categoryRepository Repository
}

func NewCategoryUseCase(cr Repository) *categoryUsecase {
	return &categoryUsecase{
		categoryRepository: cr,
	}
}

func (cu *categoryUsecase) GetAll(Sort string) []Domain {
	return cu.categoryRepository.GetAll(Sort)
}
func (cu *categoryUsecase) Create(categoryDomain *Domain) Domain {
	return cu.categoryRepository.Create(categoryDomain)
}
func (cu *categoryUsecase) GetDetail(category_id int) Domain {
	return cu.categoryRepository.GetDetail(category_id)
}
func (cu *categoryUsecase) Update(categoryDomain *Domain, category_id int) Domain {
	return cu.categoryRepository.Update(categoryDomain, category_id)
}
func (cu *categoryUsecase) Delete(category_id int) (Domain, error) {
	return cu.categoryRepository.Delete(category_id)
}
