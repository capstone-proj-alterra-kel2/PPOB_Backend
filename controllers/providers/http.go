package providers

import (
	"PPOB_BACKEND/app/middlewares"
	"PPOB_BACKEND/businesses/providers"
	"PPOB_BACKEND/controllers"
	"PPOB_BACKEND/controllers/providers/request"
	"PPOB_BACKEND/controllers/providers/response"
	"net/http"
	"strconv"

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
	claims := middlewares.GetUser(c)

	if claims.RoleID != 2 || claims.RoleID != 3 {
		return echo.ErrUnauthorized
	}

	paramID := c.Param("product-type-id")
	productTypeID, _ := strconv.Atoi(paramID)

	input := request.Provider{}
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
	return controllers.NewResponse(c, http.StatusOK, "success", "provider", response.FromDomain(providerData))
}

func (ctrl *ProviderController) Update(c echo.Context) error {
	claims := middlewares.GetUser(c)

	paramID := c.Param("provider-id")
	providerID, _ := strconv.Atoi(paramID)

	if claims.RoleID != 2 || claims.RoleID != 3 {
		return echo.ErrUnauthorized
	}

	input := request.Provider{}

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
	claims := middlewares.GetUser(c)

	paramID := c.Param("provider-id")
	providerID, _ := strconv.Atoi(paramID)

	if claims.RoleID != 2 || claims.RoleID != 3 {
		return echo.ErrUnauthorized
	}

	ctrl.providerUsecase.Delete(providerID)
	return controllers.NewResponse(c, http.StatusOK, "success", "provider deleted", "")
}
