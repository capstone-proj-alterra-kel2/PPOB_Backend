package producttypes

import (
	"PPOB_BACKEND/app/aws"
	"PPOB_BACKEND/businesses/producttypes"
	"PPOB_BACKEND/controllers"
	"PPOB_BACKEND/controllers/producttypes/request"
	"PPOB_BACKEND/controllers/producttypes/response"
	"PPOB_BACKEND/utils"
	"net/http"
	"strconv"
	"time"

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
	var result string
	input := request.ProductType{}

	image, _ := c.FormFile("image")

	if image != nil {
		isValid, message := utils.IsFileValid(image)

		if !isValid {
			return controllers.NewResponseFail(c, http.StatusBadRequest, "failed", message)
		}

		image.Filename = time.Now().String() + ".png"
		src, _ := image.Open()
		defer src.Close()
		result, _ = aws.UploadToS3(c, "product-type/", image.Filename, src)
		input.Image = result
	}

	if err := c.Bind(&input); err != nil {
		return controllers.NewResponseFail(c, http.StatusBadRequest, "failed", "invalid request")
	}
	if err := input.Validate(); err != nil {
		return controllers.NewResponseFail(c, http.StatusBadRequest, "failed", "validation failed")
	}

	productType := ctrl.productTypeUsecase.Create(input.ToDomain())

	return controllers.NewResponse(c, http.StatusCreated, "success", "product type created", response.FromDomain(productType))
}

func (ctrl *ProductTypeController) GetOne(c echo.Context) error {
	paramID := c.Param("product_type_id")
	productTypeID, _ := strconv.Atoi(paramID)

	productTypeData, err := ctrl.productTypeUsecase.GetOne(productTypeID)

	if err != nil {
		return controllers.NewResponseFail(c, http.StatusNotFound, "failed", "product type not found")
	}

	return controllers.NewResponse(c, http.StatusOK, "success", "prodcut type", response.FromDomain(productTypeData))
}

func (ctrl *ProductTypeController) Update(c echo.Context) error {
	var result string

	paramID := c.Param("product_type_id")
	productTypeID, _ := strconv.Atoi(paramID)

	input := request.UpdateProductType{}

	image, _ := c.FormFile("image")

	if image != nil {
		isValid, message := utils.IsFileValid(image)

		if !isValid {
			return controllers.NewResponseFail(c, http.StatusBadRequest, "failed", message)
		}

		image.Filename = time.Now().String() + ".png"
		src, _ := image.Open()
		defer src.Close()
		result, _ = aws.UploadToS3(c, "product-type/", image.Filename, src)
		input.Image = result
	}

	if err := c.Bind(&input); err != nil {
		return controllers.NewResponseFail(c, http.StatusBadRequest, "failed", "invalid request")
	}

	productType, err := ctrl.productTypeUsecase.Update(input.ToDomain(), productTypeID)

	if err != nil {
		return controllers.NewResponseFail(c, http.StatusBadRequest, "failed", "product type not found")
	}

	return controllers.NewResponse(c, http.StatusOK, "success", "product type updated", response.FromDomain(productType))
}

func (ctrl *ProductTypeController) Delete(c echo.Context) error {
	paramID := c.Param("product_type_id")
	productID, _ := strconv.Atoi(paramID)

	_, err := ctrl.productTypeUsecase.Delete(productID)

	if err != nil {
		return controllers.NewResponseFail(c, http.StatusBadRequest, "failed", "product type not found")
	}

	return controllers.NewResponse(c, http.StatusOK, "success", "product type deleted", "")
}
