package producttypes

import (
	"PPOB_BACKEND/app/middlewares"
	"PPOB_BACKEND/businesses/producttypes"
	"PPOB_BACKEND/controllers"
	"PPOB_BACKEND/controllers/producttypes/request"
	"PPOB_BACKEND/controllers/producttypes/response"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type ProductTypeController struct {
	productTypeUsecase producttypes.Usecase
}

func NewProductTypeController(productTypeUC producttypes.Usecase) *ProductTypeController {
	return &ProductTypeController{
		productTypeUsecase: productTypeUC,
	}
}

func (ctrl *ProductTypeController) GetAll(c echo.Context) error {
	productTypesData := ctrl.productTypeUsecase.GetAll()

	productTypes := []response.ProductType{}

	for _, productType := range productTypesData {
		productTypes = append(productTypes, response.FromDomain(productType))
	}

	return controllers.NewResponse(c, http.StatusOK, "success", "all product types", productTypes)
}

func (ctrl *ProductTypeController) Create(c echo.Context) error {
	claims := middlewares.GetUser(c)

	if claims.RoleID != 2 || claims.RoleID != 3 {
		return echo.ErrUnauthorized
	}

	input := request.ProductType{}

	if err := c.Bind(&input); err != nil {
		return controllers.NewResponse(c, http.StatusBadRequest, "failed", "invalid request", "")
	}
	if err := input.Validate(); err != nil {
		return controllers.NewResponse(c, http.StatusBadRequest, "failed", "validation failed", "")
	}

	productType := ctrl.productTypeUsecase.Create(input.ToDomain())
	return controllers.NewResponse(c, http.StatusCreated, "success", "product type created", response.FromDomain(productType))
}

func (ctrl *ProductTypeController) GetOne(c echo.Context) error {
	paramID := c.Param("product-type-id")
	productTypeID, _ := strconv.Atoi(paramID)

	productTypeData := ctrl.productTypeUsecase.GetOne(productTypeID)
	return controllers.NewResponse(c, http.StatusOK, "success", "prodcut type", response.FromDomain(productTypeData))
}

func (ctrl *ProductTypeController) Update(c echo.Context) error {
	claims := middlewares.GetUser(c)

	paramID := c.Param("product-type-id")
	productTypeID, _ := strconv.Atoi(paramID)

	if claims.RoleID != 2 || claims.RoleID != 3 {
		return echo.ErrUnauthorized
	}

	input := request.ProductType{}

	if err := c.Bind(&input); err != nil {
		return controllers.NewResponse(c, http.StatusBadRequest, "failed", "invalid request", "")
	}
	if err := input.Validate(); err != nil {
		return controllers.NewResponse(c, http.StatusBadRequest, "failed", "validation failed", "")
	}

	productType := ctrl.productTypeUsecase.Update(input.ToDomain(), productTypeID)
	return controllers.NewResponse(c, http.StatusOK, "success", "product type updated", response.FromDomain(productType))
}

func (ctrl *ProductTypeController) Delete(c echo.Context) error {
	claims := middlewares.GetUser(c)

	paramID := c.Param("product-type-id")
	productID, _ := strconv.Atoi(paramID)

	if claims.RoleID != 2 || claims.RoleID != 3 {
		return echo.ErrUnauthorized
	}

	ctrl.productTypeUsecase.Delete(productID)
	return controllers.NewResponse(c, http.StatusOK, "success", "product type deleted", "")
}
