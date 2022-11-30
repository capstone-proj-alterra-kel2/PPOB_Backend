package stocks

import (
	"PPOB_BACKEND/app/middlewares"
	"PPOB_BACKEND/businesses/stocks"
	"PPOB_BACKEND/controllers"
	"PPOB_BACKEND/controllers/stocks/request"
	"PPOB_BACKEND/controllers/stocks/response"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type StockController struct {
	stockUsecase stocks.Usecase
}

func NewStockController(stockUC stocks.Usecase) *StockController {
	return &StockController{
		stockUsecase: stockUC,
	}
}

func (ctrl *StockController) GetAll(c echo.Context) error {
	stocksData := ctrl.stockUsecase.GetAll()

	stocks := []response.Stock{}

	for _, stock := range stocksData {
		stocks = append(stocks, response.FromDomain(stock))
	}

	return controllers.NewResponse(c, http.StatusOK, "success", "all stocks", stocks)
}

func (ctrl *StockController) Create(c echo.Context) error {
	claims := middlewares.GetUser(c)

	if claims.RoleID != 2 || claims.RoleID != 3 {
		return echo.ErrUnauthorized
	}

	input := request.Stock{}

	if err := c.Bind(&input); err != nil {
		return controllers.NewResponse(c, http.StatusBadRequest, "failed", "invalid request", "")
	}
	if err := input.Validate(); err != nil {
		return controllers.NewResponse(c, http.StatusBadRequest, "failed", "validation failed", "")
	}

	stockData := ctrl.stockUsecase.Create(input.ToDomain())
	return controllers.NewResponse(c, http.StatusCreated, "success", "stock created", response.FromDomain(stockData))
}

func (ctrl *StockController) Get(c echo.Context) error {
	paramID := c.Param("stock-id")
	stockID, _ := strconv.Atoi(paramID)

	stockData := ctrl.stockUsecase.Get(stockID)
	return controllers.NewResponse(c, http.StatusOK, "success", "stock", response.FromDomain(stockData))
}

func (ctrl *StockController) Update(c echo.Context) error {
	claims := middlewares.GetUser(c)

	paramID := c.Param("stock-id")
	stockID, _ := strconv.Atoi(paramID)

	if claims.RoleID != 2 || claims.RoleID != 3 {
		return echo.ErrUnauthorized
	}

	input := request.Stock{}

	if err := c.Bind(&input); err != nil {
		return controllers.NewResponse(c, http.StatusBadRequest, "failed", "invalid request", "")
	}
	if err := input.Validate(); err != nil {
		return controllers.NewResponse(c, http.StatusBadRequest, "failed", "validation failed", "")
	}

	stockData := ctrl.stockUsecase.Update(input.ToDomain(), stockID)
	return controllers.NewResponse(c, http.StatusOK, "success", "stock updated", response.FromDomain(stockData))
}

func (ctrl *StockController) Delete(c echo.Context) error {
	claims := middlewares.GetUser(c)

	paramID := c.Param("stock-id")
	stockID, _ := strconv.Atoi(paramID)

	if claims.RoleID != 2 || claims.RoleID != 3 {
		return echo.ErrUnauthorized
	}

	ctrl.stockUsecase.Delete(stockID)
	return controllers.NewResponse(c, http.StatusOK, "success", "stock deleted", "")
}
