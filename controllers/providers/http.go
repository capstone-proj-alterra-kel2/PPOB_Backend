package providers

import (
	"PPOB_BACKEND/app/aws"
	"PPOB_BACKEND/businesses/providers"
	"PPOB_BACKEND/controllers"
	"PPOB_BACKEND/controllers/providers/request"
	"PPOB_BACKEND/controllers/providers/response"
	"PPOB_BACKEND/utils"
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
	paramID := c.Param("product_type_id")
	productTypeID, _ := strconv.Atoi(paramID)

	providersData, isProductTypeFound := ctrl.providerUsecase.GetAll(productTypeID)

	if !isProductTypeFound {
		return controllers.NewResponseFail(c, http.StatusBadRequest, "failed", "product type not found")
	}

	providers := []response.Provider{}

	for _, provider := range providersData {
		providers = append(providers, response.FromDomain(provider))
	}

	return controllers.NewResponse(c, http.StatusOK, "success", "all providers", providers)
}

func (ctrl *ProviderController) Create(c echo.Context) error {
	paramID := c.Param("product_type_id")
	productTypeID, _ := strconv.Atoi(paramID)

	input := request.Provider{}
	var result string

	image, _ := c.FormFile("image")

	if image != nil {
		isValid, message := utils.IsFileValid(image)

		if !isValid {
			return controllers.NewResponseFail(c, http.StatusBadRequest, "failed", message)
		}

		image.Filename = time.Now().String() + ".png"
		src, _ := image.Open()
		defer src.Close()
		result, _ = aws.UploadToS3(c, "provider/", image.Filename, src)
		input.Image = result
	}

	input.ProductTypeID = productTypeID

	if err := c.Bind(&input); err != nil {
		return controllers.NewResponseFail(c, http.StatusBadRequest, "failed", "invalid request")
	}
	if err := input.Validate(); err != nil {
		return controllers.NewResponseFail(c, http.StatusBadRequest, "failed", "validation failed")
	}

	providerData, isProductTypeFound, isNameDuplicated := ctrl.providerUsecase.Create(input.ToDomain(), productTypeID)

	if !isProductTypeFound || providerData.ID == 0 {
		return controllers.NewResponseFail(c, http.StatusBadRequest, "failed", "product types not found")
	}

	if isNameDuplicated {
		return controllers.NewResponseFail(c, http.StatusBadRequest, "failed", "provider already exist")
	}

	return controllers.NewResponse(c, http.StatusCreated, "success", "provider created", response.FromDomain(providerData))
}

func (ctrl *ProviderController) GetOne(c echo.Context) error {
	paramProviderID := c.Param("provider_id")
	providerID, _ := strconv.Atoi(paramProviderID)

	paramProductTypeID := c.Param("product_type_id")
	productTypeID, _ := strconv.Atoi(paramProductTypeID)

	providerData, isProductTypeFound, isProviderFound := ctrl.providerUsecase.GetOne(providerID, productTypeID)

	if !isProductTypeFound {
		return controllers.NewResponseFail(c, http.StatusBadRequest, "failed", "product type not found")
	}

	if !isProviderFound {
		return controllers.NewResponseFail(c, http.StatusBadRequest, "failed", "provider not found")
	}

	return controllers.NewResponse(c, http.StatusOK, "success", "provider", response.FromDomain(providerData))
}

func (ctrl *ProviderController) GetByPhone(c echo.Context) error {
	phoneNumber := request.InputPhone{}

	if err := c.Bind(&phoneNumber); err != nil {
		return controllers.NewResponseFail(c, http.StatusBadRequest, "failed", "invalid request")
	}
	if err := phoneNumber.Validate(); err != nil {
		return controllers.NewResponseFail(c, http.StatusBadRequest, "failed", "validation failed")
	}

	paramID := c.Param("product_type_id")
	productTypeID, _ := strconv.Atoi(paramID)

	providerData, isProductTypeFound := ctrl.providerUsecase.GetByPhone(phoneNumber.PhoneNumber, productTypeID)

	if !isProductTypeFound {
		return controllers.NewResponseFail(c, http.StatusBadRequest, "failed", "product type not found")
	}

	if providerData.ID == 0 {
		return controllers.NewResponseFail(c, http.StatusBadRequest, "failed", "provider not found")
	}

	inputProviderUpdates := []request.UpdateCheckProduct{}
	inputProvider := request.InputProvider{}

	for _, v := range providerData.Products {
		inputProviderUpdates = append(inputProviderUpdates, request.UpdateCheckProduct{
			ID:                    int(v.ID),
			Name:                  v.Name,
			Description:           v.Description,
			Price:                 v.Price,
			ProviderID:            v.ProviderID,
			Stock:                 v.Stock,
			Status:                v.Status,
			AdditionalInformation: v.AdditionalInformation,
			IsAvailable:           v.IsAvailable,
			PriceStatus:           v.PriceStatus,
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
	paramID := c.Param("provider_id")
	providerID, _ := strconv.Atoi(paramID)

	input := request.UpdateData{}

	var result string

	image, _ := c.FormFile("image")
	if image != nil {
		isValid, message := utils.IsFileValid(image)

		if !isValid {
			return controllers.NewResponseFail(c, http.StatusBadRequest, "failed", message)
		}

		image.Filename = time.Now().String() + ".png"
		src, _ := image.Open()
		defer src.Close()
		result, _ = aws.UploadToS3(c, "provider/", image.Filename, src)
		input.Image = result
	}

	if err := c.Bind(&input); err != nil {
		return controllers.NewResponseFail(c, http.StatusBadRequest, "failed", "invalid request")
	}
	if err := input.Validate(); err != nil {
		return controllers.NewResponseFail(c, http.StatusBadRequest, "failed", "validation failed")
	}

	providerData := ctrl.providerUsecase.Update(input.ToDomain(), providerID)
	return controllers.NewResponse(c, http.StatusOK, "success", "provider updated", response.FromDomain(providerData))
}

func (ctrl *ProviderController) Delete(c echo.Context) error {
	paramID := c.Param("provider_id")
	providerID, _ := strconv.Atoi(paramID)

	ctrl.providerUsecase.Delete(providerID)
	return controllers.NewResponse(c, http.StatusOK, "success", "provider deleted", "")
}
