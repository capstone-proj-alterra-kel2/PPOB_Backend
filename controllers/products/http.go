package products

import (
	"PPOB_BACKEND/businesses/products"
	"PPOB_BACKEND/controllers"
	"PPOB_BACKEND/controllers/products/request"
	"PPOB_BACKEND/controllers/products/response"
	"fmt"
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
	input := request.Product{}

	if err := c.Bind(&input); err != nil {
		return controllers.NewResponseFail(c, http.StatusBadRequest, "failed", "invalid request")
	}
	if err := input.Validate(); err != nil {
		fmt.Println(err.Error())
		return controllers.NewResponseFail(c, http.StatusBadRequest, "failed", "validation failed")
	}

	if *input.IsPromo {
		if input.Discount == 0 {
			return controllers.NewResponseFail(c, http.StatusBadRequest, "failed", "discount isn't allowed empty")
		}

		if input.PromoStartDate == "" || input.PromoEndDate == "" {
			return controllers.NewResponseFail(c, http.StatusBadRequest, "failed", "promo start or end date aren't allowed empty")
		}
	}

	product := ctrl.productUsecase.Create(input.ToDomain())
	return controllers.NewResponse(c, http.StatusCreated, "success", "product created", response.FromDomain(product))
}

func (ctrl *ProductController) Update(c echo.Context) error {
	paramID := c.Param("product-id")
	productID, _ := strconv.Atoi(paramID)

	input := request.Product{}

	if err := c.Bind(&input); err != nil {
		return controllers.NewResponseFail(c, http.StatusBadRequest, "failed", "invalid request")
	}

	product := ctrl.productUsecase.Update(input.ToDomain(), productID)
	return controllers.NewResponse(c, http.StatusOK, "success", "product updated", response.FromDomain(product))
}

func (ctrl *ProductController) Delete(c echo.Context) error {
	paramID := c.Param("product-id")
	productID, _ := strconv.Atoi(paramID)

	ctrl.productUsecase.Delete(productID)
	return controllers.NewResponse(c, http.StatusOK, "success", "product deleted", "")
}
