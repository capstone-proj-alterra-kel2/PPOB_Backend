package faq

type Domain struct {
	ID       uint
	Question string
	Answer   string
}

type Repository interface {
	GetAll() []Domain
	Create(faqDomain *Domain) Domain
	Update(faqDomain *Domain, faq_id int) (Domain, error)
	Delete(faq_id int) (Domain, error)
}

type Usecase interface {
	GetAll() []Domain
	Create(faqDomain *Domain) Domain
	Update(faqDomain *Domain, faq_id int) (Domain, error)
	Delete(faq_id int) (Domain, error)
}
