package request

import (
	"PPOB_BACKEND/businesses/landing_pages/faq"

	"github.com/go-playground/validator"
)

type FAQ struct {
	Question string `json:"question" form:"question" validate:"required"`
	Answer   string `json:"answer" form:"answer" validate:"required"`
}

func (faqReq *FAQ) ToDomain() *faq.Domain {
	return &faq.Domain{
		Question: faqReq.Question,
		Answer:   faqReq.Answer,
	}
}

func (req *FAQ) Validate() error {
	validate := validator.New()

	err := validate.Struct(req)
	return err
}
