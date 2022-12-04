package products

import (
	"PPOB_BACKEND/app/middlewares"
	"PPOB_BACKEND/businesses/products"
	"PPOB_BACKEND/controllers"
	"PPOB_BACKEND/controllers/products/request"
	"PPOB_BACKEND/controllers/products/response"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/morkid/paginate"
)

type ProductController struct {
	productUsecase products.Usecase
}

func NewProductController(productUC products.Usecase) *ProductController {
	return &ProductController{
		productUsecase: productUC,
	}
}

func (ctrl *ProductController) GetAll(c echo.Context) error {
	model := ctrl.productUsecase.GetAll()
	pg := paginate.New(&paginate.Config{
		DefaultSize: 6,
	})

	products := []response.Product{}

	return controllers.NewResponse(c, http.StatusOK, "success", "all products", pg.With(model).Request(c.Request()).Response(&products))
}

func (ctrl *ProductController) GetOne(c echo.Context) error {

	paramID := c.Param("product-id")
	productID, _ := strconv.Atoi(paramID)

	productData := ctrl.productUsecase.GetOne(productID)

	return controllers.NewResponse(c, http.StatusOK, "success", "product", response.FromDomain(productData))
}

func (ctrl *ProductController) Create(c echo.Context) error {
	claims := middlewares.GetUser(c)

	// if claims.RoleID != 2 || claims.RoleID != 3 {
	// 	return echo.ErrUnauthorized
	// }

	if claims.RoleID != 1 {
		return echo.ErrUnauthorized
	}

	input := request.Product{}

	if err := c.Bind(&input); err != nil {
		return controllers.NewResponse(c, http.StatusBadRequest, "failed", "invalid request", "")
	}
	if err := input.Validate(); err != nil {
		return controllers.NewResponse(c, http.StatusBadRequest, "failed", "validation failed", "")
	}

	product := ctrl.productUsecase.Create(input.ToDomain())
	return controllers.NewResponse(c, http.StatusCreated, "success", "product created", response.FromDomain(product))
}

func (ctrl *ProductController) Update(c echo.Context) error {
	claims := middlewares.GetUser(c)

	paramID := c.Param("product-id")
	productID, _ := strconv.Atoi(paramID)

	if claims.RoleID != 1 {
		return echo.ErrUnauthorized
	}

	input := request.Product{}

	if err := c.Bind(&input); err != nil {
		return controllers.NewResponse(c, http.StatusBadRequest, "failed", "invalid request", "")
	}

	product := ctrl.productUsecase.Update(input.ToDomain(), productID)
	return controllers.NewResponse(c, http.StatusOK, "success", "product updated", response.FromDomain(product))
}

func (ctrl *ProductController) Delete(c echo.Context) error {
	claims := middlewares.GetUser(c)

	paramID := c.Param("product-id")
	productID, _ := strconv.Atoi(paramID)

	if claims.RoleID != 1 {
		return echo.ErrUnauthorized
	}

	ctrl.productUsecase.Delete(productID)
	return controllers.NewResponse(c, http.StatusOK, "success", "product deleted", "")
}
