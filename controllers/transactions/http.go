package transactions

import (
	"PPOB_BACKEND/app/middlewares"
	"PPOB_BACKEND/businesses/products"
	"PPOB_BACKEND/businesses/transactions"
	"PPOB_BACKEND/businesses/users"
	"PPOB_BACKEND/controllers"
	productRequest "PPOB_BACKEND/controllers/products/request"
	"PPOB_BACKEND/controllers/transactions/request"
	"PPOB_BACKEND/controllers/transactions/response"
	userRequest "PPOB_BACKEND/controllers/users/request"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/morkid/paginate"
)

type TransactionController struct {
	transactionUsecase transactions.Usecase
	productUsecase     products.Usecase
	userUsecase        users.Usecase
}

func NewTransactionController(transactionUC transactions.Usecase, productUC products.Usecase, userUC users.Usecase) *TransactionController {
	return &TransactionController{
		transactionUsecase: transactionUC,
		productUsecase:     productUC,
		userUsecase:        userUC,
	}
}

func (tc *TransactionController) GetAll(c echo.Context) error {
	pg := paginate.New()

	//get user
	userID := middlewares.GetUserID(c)
	user := tc.userUsecase.Profile(userID)

	// if user.RoleName != "admin" || user.RoleName != "superadmin" {
	// 	return controllers.NewResponseFail(c, http.StatusBadRequest, "failed", "unauthorized")
	// }
	if user.ID == 0 {
		return controllers.NewResponseFail(c, http.StatusNotFound, "failed", "cannot get detail user")
	}

	size, _ := strconv.Atoi(c.QueryParam("size"))
	page, _ := strconv.Atoi(c.QueryParam("page"))
	search := c.QueryParam("search")
	sort := c.QueryParam("sort")

	transactionDatas, modelTransaction := tc.transactionUsecase.GetAll(page, size, sort, search)

	transactions := []response.Transaction{}
	for _, transaction := range transactionDatas {
		transactions = append(transactions, response.FromDomain(transaction))
	}

	return controllers.NewResponse(c, http.StatusOK, "success", "all transactions", pg.Response(modelTransaction, c.Request(), &transactions))

}

func (tc *TransactionController) GetDetail(c echo.Context) error {
	paramID := c.Param("transaction-id")
	transactionID, _ := strconv.Atoi(paramID)

	transactionDetail, isTransactionFound := tc.transactionUsecase.GetDetail(transactionID)

	if !isTransactionFound {
		return controllers.NewResponseFail(c, http.StatusNotFound, "failed", "transaction not found")
	}

	return controllers.NewResponse(c, http.StatusOK, "success", "transaction", response.FromDomain(transactionDetail))
}

func (tc *TransactionController) GetTransactionHistory(c echo.Context) error {

	middlewareID := middlewares.GetUserID(c)
	userID, _ := strconv.Atoi(middlewareID)

	transactionHistoryData := tc.transactionUsecase.GetTransactionHistory(userID)

	transactions := []response.Transaction{}

	for _, transaction := range transactionHistoryData {
		transactions = append(transactions, response.FromDomain(transaction))
	}

	return controllers.NewResponse(c, http.StatusOK, "success", "transaction history", transactions)
}

func (tc *TransactionController) Create(c echo.Context) error {
	totalAmount := 0
	productDiscount := 0

	request := request.Transaction{}
	c.Bind(&request)

	//get user
	userID := middlewares.GetUserID(c)

	user := tc.userUsecase.Profile(userID)

	if user.RoleName == "admin" || user.RoleName == "superadmin" {
		return controllers.NewResponseFail(c, http.StatusBadRequest, "failed", "prevent getting detail admin & superadmin")
	}
	if user.ID == 0 {
		return controllers.NewResponseFail(c, http.StatusNotFound, "failed", "cannot get detail user")
	}

	//get product
	product := tc.productUsecase.GetOne(request.ProductID)

	if product.ID == 0 {
		return controllers.NewResponseFail(c, http.StatusNotFound, "failed", "product not found")
	}

	if *product.Stock < 1 {
		return controllers.NewResponseFail(c, http.StatusBadRequest, "failed", "product stock empty")
	}

	if *product.IsPromoActive {
		if user.Balance < (product.Price - product.Discount) {
			return controllers.NewResponseFail(c, http.StatusBadRequest, "failed", "not enough balance")
		}

		totalAmount = product.Price - product.Discount
		productDiscount = product.Discount
	} else {
		if user.Balance < product.Price {
			return controllers.NewResponseFail(c, http.StatusBadRequest, "failed", "not enough balance")
		}

		totalAmount = product.Price
	}

	transaction := tc.transactionUsecase.Create(&product, &user, totalAmount, productDiscount)

	var updatedBalance int

	// update balance
	if *product.IsPromoActive {
		updatedBalance = user.Balance - (product.Price - product.Discount)
	} else {
		updatedBalance = user.Balance - product.Price
	}

	requestUpdateBalance := userRequest.UpdateBalance{}
	requestUpdateBalance.Balance = updatedBalance
	tc.userUsecase.UpdateBalance(userID, requestUpdateBalance.ToDomain())

	// update stock and status
	requestUpdateStockStatus := productRequest.UpdateStockStatus{}
	updatedStock := *product.Stock - 1
	updatedTotalPurchased := product.TotalPurchased + 1

	requestUpdateStockStatus.Stock = updatedStock
	requestUpdateStockStatus.TotalPurchased = updatedTotalPurchased

	if updatedStock == 0 {
		requestUpdateStockStatus.IsAvailable = false
		requestUpdateStockStatus.Status = "Habis"
	}

	tc.productUsecase.UpdateStockStatus(requestUpdateStockStatus.ToDomain(), int(product.ID))

	return controllers.NewResponse(c, http.StatusOK, "success", "transaction success", response.FromDomain(transaction))
}

func (tc *TransactionController) Delete(c echo.Context) error {
	paramID := c.Param("transaction-id")
	transactionID, _ := strconv.Atoi(paramID)

	_, isTransactionFound := tc.transactionUsecase.Delete(transactionID)

	if !isTransactionFound {
		return controllers.NewResponseFail(c, http.StatusNotFound, "failed", "transaction not found")
	}

	return controllers.NewResponse(c, http.StatusOK, "success", "delete tranasaction", "")
}
