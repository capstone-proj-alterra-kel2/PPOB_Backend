package providers

import (
	"PPOB_BACKEND/app/aws"
	"PPOB_BACKEND/businesses/providers"
	"PPOB_BACKEND/controllers"
	"PPOB_BACKEND/controllers/providers/request"
	"PPOB_BACKEND/controllers/providers/response"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
)

type ProviderController struct {
	providerUsecase providers.Usecase
}

func NewProviderController(providerUC providers.Usecase) *ProviderController {
	return &ProviderController{
		providerUsecase: providerUC,
	}
}

func (ctrl *ProviderController) GetAll(c echo.Context) error {
	providersData := ctrl.providerUsecase.GetAll()

	providers := []response.Provider{}

	for _, provider := range providersData {
		providers = append(providers, response.FromDomain(provider))
	}

	return controllers.NewResponse(c, http.StatusOK, "success", "all providers", providers)
}

func (ctrl *ProviderController) Create(c echo.Context) error {
	paramID := c.Param("product-type-id")
	productTypeID, _ := strconv.Atoi(paramID)

	input := request.Provider{}
	var result string

	image, _ := c.FormFile("image")
	image.Filename = time.Now().String() + ".png"
	if image != nil {
		src, _ := image.Open()
		defer src.Close()
		result, _ = aws.UploadToS3(c, "provider/", image.Filename, src)
		input.Image = result
	}

	input.ProductTypeID = productTypeID

	if err := c.Bind(&input); err != nil {
		return controllers.NewResponse(c, http.StatusBadRequest, "failed", "invalid request", "")
	}
	if err := input.Validate(); err != nil {
		return controllers.NewResponse(c, http.StatusBadRequest, "failed", "validation failed", "")
	}

	providerData := ctrl.providerUsecase.Create(input.ToDomain())
	return controllers.NewResponse(c, http.StatusCreated, "success", "provider created", response.FromDomain(providerData))
}

func (ctrl *ProviderController) GetOne(c echo.Context) error {
	paramID := c.Param("provider-id")
	providerID, _ := strconv.Atoi(paramID)

	providerData := ctrl.providerUsecase.GetOne(providerID)
	return controllers.NewResponse(c, http.StatusOK, "success", "provider", response.FromDomain(providerData))
}

func (ctrl *ProviderController) GetByPhone(c echo.Context) error {
	phoneNumber := c.FormValue("phone_number")

	paramID := c.Param("product-type-id")
	productTypeID, _ := strconv.Atoi(paramID)

	providerData := ctrl.providerUsecase.GetByPhone(phoneNumber, productTypeID)

	inputProviderUpdates := []request.UpdateCheckProduct{}
	inputProvider := request.InputProvider{}

	for _, v := range providerData.Products {
		inputProviderUpdates = append(inputProviderUpdates, request.UpdateCheckProduct{
			ID:                    int(v.ID),
			Name:                  v.Name,
			Category:              v.Category,
			Description:           v.Description,
			Price:                 v.Price,
			ProviderID:            v.ProviderID,
			Stock:                 v.Stock,
			Status:                v.Status,
			AdditionalInformation: v.AdditionalInformation,
			IsAvailable:           v.IsAvailable,
			IsPromo:               v.IsPromo,
			IsPromoActive:         v.IsPromoActive,
			Discount:              v.Discount,
			PromoStartDate:        v.PromoStartDate,
			PromoEndDate:          v.PromoEndDate,
		})
	}

	inputProvider.ID = int(providerData.ID)
	inputProvider.Name = providerData.Name
	inputProvider.Image = providerData.Image
	inputProvider.ProductTypeID = providerData.ProductTypeID
	inputProvider.Products = inputProviderUpdates

	providerUpdatedData := ctrl.providerUsecase.UpdateCheck(inputProvider.ToDomain(), int(providerData.ID))

	return controllers.NewResponse(c, http.StatusOK, "success", "provider", response.FromDomain(providerUpdatedData))
}

func (ctrl *ProviderController) Update(c echo.Context) error {
	paramID := c.Param("provider-id")
	providerID, _ := strconv.Atoi(paramID)

	input := request.Provider{}

	var result string

	image, _ := c.FormFile("image")
	image.Filename = time.Now().String() + ".png"
	if image != nil {
		src, _ := image.Open()
		defer src.Close()
		result, _ = aws.UploadToS3(c, "provider/", image.Filename, src)
		input.Image = result
	}

	if err := c.Bind(&input); err != nil {
		return controllers.NewResponse(c, http.StatusBadRequest, "failed", "invalid request", "")
	}
	if err := input.Validate(); err != nil {
		return controllers.NewResponse(c, http.StatusBadRequest, "failed", "validation failed", "")
	}

	providerData := ctrl.providerUsecase.Update(input.ToDomain(), providerID)
	return controllers.NewResponse(c, http.StatusOK, "success", "provider updated", response.FromDomain(providerData))
}

func (ctrl *ProviderController) Delete(c echo.Context) error {
	paramID := c.Param("provider-id")
	providerID, _ := strconv.Atoi(paramID)

	ctrl.providerUsecase.Delete(providerID)
	return controllers.NewResponse(c, http.StatusOK, "success", "provider deleted", "")
}
