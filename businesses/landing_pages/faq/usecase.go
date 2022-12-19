package faq

type FAQUsecase struct {
	faqUseCase Usecase
}

func NewFaqUsecase(FAQUC Usecase) *FAQUsecase {
	return &FAQUsecase{
		faqUseCase: FAQUC,
	}
}

func (fuc *FAQUsecase) GetAll() []Domain {
	return fuc.faqUseCase.GetAll()
}
func (fuc *FAQUsecase) Create(faqDomain *Domain) Domain {
	return fuc.faqUseCase.Create(faqDomain)
}
func (fuc *FAQUsecase) Update(faqDomain *Domain, faq_id int) (Domain, error) {
	return fuc.faqUseCase.Update(faqDomain, faq_id)
}
func (fuc *FAQUsecase) Delete(faq_id int) (Domain, error) {
	return fuc.faqUseCase.Delete(faq_id)
}
