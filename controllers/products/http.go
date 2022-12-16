package products

import (
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
	pg := paginate.New()

	size, _ := strconv.Atoi(c.QueryParam("size"))
	page, _ := strconv.Atoi(c.QueryParam("page"))
	search := c.QueryParam("search")
	sort := c.QueryParam("sort")

	modelProduct, productDomain := ctrl.productUsecase.GetAll(page, size, sort, search)

	for _, value := range productDomain {
		input := request.UpdatePromoProduct{
			ID:             value.ID,
			IsAvailable:    value.IsAvailable,
			PriceStatus:    value.PriceStatus,
			IsPromoActive:  value.IsPromoActive,
			Discount:       value.Discount,
			PromoStartDate: value.PromoStartDate,
			PromoEndDate:   value.PromoEndDate,
		}
		updateResult := ctrl.productUsecase.UpdatePromo(input.ToDomain())

		value.IsAvailable = updateResult.IsAvailable
		value.PriceStatus = updateResult.PriceStatus
		value.IsPromoActive = updateResult.IsPromoActive
		value.Discount = updateResult.Discount
		value.PromoStartDate = updateResult.PromoStartDate
		value.PromoEndDate = updateResult.PromoEndDate
	}

	products := []response.Product{}
	for _, product := range productDomain {
		products = append(products, response.FromDomain(product))
	}

	return controllers.NewResponse(c, http.StatusOK, "success", "all products", pg.With(modelProduct).Request(c.Request()).Response(&products))
}

func (ctrl *ProductController) GetAllForUSer(c echo.Context) error {
	productDomain := ctrl.productUsecase.GetAllForUser()

	for _, value := range productDomain {
		input := request.UpdatePromoProduct{
			ID:             value.ID,
			IsAvailable:    value.IsAvailable,
			PriceStatus:    value.PriceStatus,
			IsPromoActive:  value.IsPromoActive,
			Discount:       value.Discount,
			PromoStartDate: value.PromoStartDate,
			PromoEndDate:   value.PromoEndDate,
		}
		updateResult := ctrl.productUsecase.UpdatePromo(input.ToDomain())

		value.IsAvailable = updateResult.IsAvailable
		value.PriceStatus = updateResult.PriceStatus
		value.IsPromoActive = updateResult.IsPromoActive
		value.Discount = updateResult.Discount
		value.PromoStartDate = updateResult.PromoStartDate
		value.PromoEndDate = updateResult.PromoEndDate
	}

	products := []response.Product{}
	for _, product := range productDomain {
		products = append(products, response.FromDomain(product))
	}

	return controllers.NewResponse(c, http.StatusOK, "success", "all products", &products)
}

func (ctrl *ProductController) GetOne(c echo.Context) error {
	paramID := c.Param("product_id")
	productID, _ := strconv.Atoi(paramID)

	productData, err := ctrl.productUsecase.GetOne(productID)

	if err != nil {
		return controllers.NewResponseFail(c, http.StatusNotFound, "failed", "product not found")
	}

	input := request.UpdatePromoProduct{
		ID:             productData.ID,
		IsAvailable:    productData.IsAvailable,
		PriceStatus:    productData.PriceStatus,
		IsPromoActive:  productData.IsPromoActive,
		Discount:       productData.Discount,
		PromoStartDate: productData.PromoStartDate,
		PromoEndDate:   productData.PromoEndDate,
	}

	updateResult := ctrl.productUsecase.UpdatePromo(input.ToDomain())
	productData.IsAvailable = updateResult.IsAvailable
	productData.PriceStatus = updateResult.PriceStatus
	productData.IsPromoActive = updateResult.IsPromoActive
	productData.Discount = updateResult.Discount
	productData.PromoStartDate = updateResult.PromoStartDate
	productData.PromoEndDate = updateResult.PromoEndDate

	return controllers.NewResponse(c, http.StatusOK, "success", "product", response.FromDomain(productData))
}

func (ctrl *ProductController) Create(c echo.Context) error {
	input := request.Product{}
	input.TotalPurchased = 0
	zeroValue := 0

	if err := c.Bind(&input); err != nil {
		return controllers.NewResponseFail(c, http.StatusBadRequest, "failed", "invalid request")
	}
	if err := input.Validate(); err != nil {
		return controllers.NewResponseFail(c, http.StatusBadRequest, "failed", "validation failed")
	}

	if input.PriceStatus == "promo" {
		if input.Discount == nil || *input.Discount == 0 {
			return controllers.NewResponseFail(c, http.StatusBadRequest, "failed", "discount isn't allowed empty")
		}
	} else {
		input.Discount = &zeroValue
		input.PromoStartDate = ""
		input.PromoEndDate = ""
	}

	product, isDateValid, isProviderFound := ctrl.productUsecase.Create(input.ToDomain())

	if !isDateValid {
		return controllers.NewResponseFail(c, http.StatusNotFound, "failed", "invalid date input")
	}

	if !isProviderFound || product.ID == 0 {
		return controllers.NewResponseFail(c, http.StatusNotFound, "failed", "provider not found")
	}

	product.PriceStatus = input.PriceStatus

	return controllers.NewResponse(c, http.StatusCreated, "success", "product created", response.FromDomain(product))
}

func (ctrl *ProductController) UpdateData(c echo.Context) error {
	paramID := c.Param("product_id")
	productID, _ := strconv.Atoi(paramID)

	promoActiveFalse := false
	zeroValue := 0

	input := request.UpdateDataProduct{}

	if err := c.Bind(&input); err != nil {
		return controllers.NewResponseFail(c, http.StatusBadRequest, "failed", "invalid request")
	}

	if input.PriceStatus == "promo" {
		if input.Discount == nil || *input.Discount == 0 {
			return controllers.NewResponseFail(c, http.StatusBadRequest, "failed", "discount isn't allowed empty")
		}

	} else {
		input.Discount = &zeroValue
		input.IsPromoActive = &promoActiveFalse
		input.PromoStartDate = ""
		input.PromoEndDate = ""
	}

	product, err, isDateValid, isProviderFound := ctrl.productUsecase.UpdateData(input.ToDomain(), productID)

	if err != nil {
		return controllers.NewResponseFail(c, http.StatusNotFound, "failed", "product not found")
	}

	if !isProviderFound {
		return controllers.NewResponseFail(c, http.StatusNotFound, "failed", "provider not found")
	}

	if !isDateValid {
		return controllers.NewResponseFail(c, http.StatusNotFound, "failed", "invalid date input")
	}

	product.PriceStatus = input.PriceStatus
	return controllers.NewResponse(c, http.StatusOK, "success", "product updated", response.FromDomain(product))
}

func (ctrl *ProductController) Delete(c echo.Context) error {
	paramID := c.Param("product_id")
	productID, _ := strconv.Atoi(paramID)

	_, err := ctrl.productUsecase.Delete(productID)

	if err != nil {
		return controllers.NewResponseFail(c, http.StatusNotFound, "failed", "product not found")
	}

	return controllers.NewResponse(c, http.StatusOK, "success", "product deleted", "")
}
