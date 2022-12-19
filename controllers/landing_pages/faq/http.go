package faq

import (
	"PPOB_BACKEND/businesses/landing_pages/faq"
	"PPOB_BACKEND/controllers"
	"PPOB_BACKEND/controllers/landing_pages/faq/request"
	"PPOB_BACKEND/controllers/landing_pages/faq/response"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type FAQController struct {
	FAQUsecase faq.Usecase
}

func NewFAQController(FAQUC faq.Usecase) *FAQController {
	return &FAQController{
		FAQUsecase: FAQUC,
	}
}

func (fc *FAQController) GetAll(c echo.Context) error {
	var resFAQ []response.FAQ

	faqDomain := fc.FAQUsecase.GetAll()

	for _, faq := range faqDomain {
		resFAQ = append(resFAQ, response.FromDomain(faq))
	}

	return controllers.NewResponse(c, http.StatusOK, "success", "all faqs", resFAQ)
}

func (fc *FAQController) Create(c echo.Context) error {
	input := request.FAQ{}

	if err := c.Bind(&input); err != nil {
		return controllers.NewResponseFail(c, http.StatusBadRequest, "failed", "invalid request")
	}
	if err := input.Validate(); err != nil {
		return controllers.NewResponseFail(c, http.StatusBadRequest, "failed", "validation failed")
	}

	faqCreate := fc.FAQUsecase.Create(input.ToDomain())

	return controllers.NewResponse(c, http.StatusCreated, "success", "faq created", response.FromDomain(faqCreate))
}

func (fc *FAQController) Update(c echo.Context) error {
	input := request.FAQ{}

	paramID := c.Param("faq_id")
	faqID, _ := strconv.Atoi(paramID)

	if err := c.Bind(&input); err != nil {
		return controllers.NewResponseFail(c, http.StatusBadRequest, "failed", "invalid request")
	}

	faqUpdate, errFound := fc.FAQUsecase.Update(input.ToDomain(), faqID)

	if errFound != nil {
		return controllers.NewResponseFail(c, http.StatusNotFound, "failed", "faq not found")
	}

	return controllers.NewResponse(c, http.StatusOK, "success", "update success", response.FromDomain(faqUpdate))
}

func (fc *FAQController) Delete(c echo.Context) error {
	paramID := c.Param("faq_id")
	faqID, _ := strconv.Atoi(paramID)

	_, err := fc.FAQUsecase.Delete(faqID)

	if err != nil {
		return controllers.NewResponseFail(c, http.StatusNotFound, "failed", "faq not found")
	}

	return controllers.NewResponse(c, http.StatusOK, "success", "success deleted", "")
}
