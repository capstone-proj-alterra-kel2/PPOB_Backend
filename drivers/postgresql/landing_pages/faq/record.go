package faq

import "PPOB_BACKEND/businesses/landing_pages/faq"

type FAQ struct {
	ID       uint   `json:"id" gorm:"size:100;primaryKey"`
	Question string `json:"question"`
	Answer   string `json:"answer"`
}

func FromDomain(domain *faq.Domain) *FAQ {

	return &FAQ{
		ID:       domain.ID,
		Question: domain.Question,
		Answer:   domain.Answer,
	}
}

func (recFaq *FAQ) ToDomain() faq.Domain {
	return faq.Domain{
		ID:       recFaq.ID,
		Question: recFaq.Question,
		Answer:   recFaq.Answer,
	}
}
