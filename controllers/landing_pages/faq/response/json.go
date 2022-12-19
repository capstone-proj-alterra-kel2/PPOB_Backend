package response

import "PPOB_BACKEND/businesses/landing_pages/faq"

type FAQ struct {
	ID       int    `json:"id"`
	Question string `json:"question"`
	Answer   string `json:"answer"`
}

func FromDomain(Domain faq.Domain) FAQ {
	return FAQ{
		ID:       int(Domain.ID),
		Question: Domain.Question,
		Answer:   Domain.Answer,
	}
}
