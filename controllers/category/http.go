package category

import (
	"PPOB_BACKEND/businesses/category"
	"PPOB_BACKEND/controllers"
	"PPOB_BACKEND/controllers/category/request"
	"PPOB_BACKEND/controllers/category/response"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type CategoryController struct {
	categoryUsecase category.Usecase
}

func NewCategoryController(categoryUC category.Usecase) *CategoryController {
	return &CategoryController{
		categoryUsecase: categoryUC,
	}
}

func (ctrl *CategoryController) GetAll(c echo.Context) error {
	categoryName := c.QueryParam("category")

	categoryDomain := ctrl.categoryUsecase.GetAll(categoryName)

	categories := []response.Category{}
	for _, category := range categoryDomain {
		categories = append(categories, response.FromDomain(category))
	}

	return controllers.NewResponse(c, http.StatusOK, "success", "categories data", categories)
}

func (ctrl *CategoryController) Create(c echo.Context) error {
	input := request.Category{}

	if err := c.Bind(&input); err != nil {
		return controllers.NewResponseFail(c, http.StatusBadRequest, "failed", "invalid request")
	}
	if err := input.Validate(); err != nil {
		return controllers.NewResponseFail(c, http.StatusBadRequest, "failed", "validation failed")
	}

	category := ctrl.categoryUsecase.Create(input.ToDomain())
	return controllers.NewResponse(c, http.StatusCreated, "success", "category created", response.FromDomain(category))
}

func (ctrl *CategoryController) GetDetail(c echo.Context) error {
	paramID := c.Param("category_id")
	categoryID, _ := strconv.Atoi(paramID)

	categoryData := ctrl.categoryUsecase.GetDetail(categoryID)

	if categoryData.ID == 0 {
		return controllers.NewResponseFail(c, http.StatusNotFound, "failed", "category not found")
	}

	return controllers.NewResponse(c, http.StatusOK, "success", "product", response.FromDomain(categoryData))
}

func (ctrl *CategoryController) Update(c echo.Context) error {
	paramID := c.Param("category_id")
	categoryID, _ := strconv.Atoi(paramID)

	input := request.Category{}

	if err := c.Bind(&input); err != nil {
		return controllers.NewResponseFail(c, http.StatusBadRequest, "failed", "invalid request")
	}

	categoryData := ctrl.categoryUsecase.Update(input.ToDomain(), categoryID)

	return controllers.NewResponse(c, http.StatusOK, "success", "category updated", response.FromDomain(categoryData))
}

func (ctrl *CategoryController) Delete(c echo.Context) error {
	paramID := c.Param("category_id")
	categoryID, _ := strconv.Atoi(paramID)

	_, err := ctrl.categoryUsecase.Delete(categoryID)

	if err != nil {
		return controllers.NewResponseFail(c, http.StatusNotFound, "failed", "category not found")
	}

	return controllers.NewResponse(c, http.StatusOK, "success", "category deleted", "")
}
