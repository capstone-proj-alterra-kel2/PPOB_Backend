package wallet_histories

import (
	"PPOB_BACKEND/app/middlewares"
	"PPOB_BACKEND/businesses/users"
	"PPOB_BACKEND/businesses/wallet_histories"
	"PPOB_BACKEND/controllers"
	"PPOB_BACKEND/controllers/wallet_histories/request"
	"PPOB_BACKEND/controllers/wallet_histories/response"
	"net/http"

	"github.com/labstack/echo/v4"
)

type WalletHistoryController struct {
	walletHistoryUsecase wallet_histories.Usecase
	userUsecase          users.Usecase
}

func NewWalletHistoryController(walletHistoryUC wallet_histories.Usecase, userUC users.Usecase) *WalletHistoryController {
	return &WalletHistoryController{
		walletHistoryUsecase: walletHistoryUC,
		userUsecase:          userUC,
	}
}

func (ctrl *WalletHistoryController) GetWalletHistories(c echo.Context) error {
	idUser := middlewares.GetUserID(c)
	profile := ctrl.userUsecase.Profile(idUser)
	walletHistoryData := ctrl.walletHistoryUsecase.GetWalletHistories(profile.Wallet.NoWallet)

	return controllers.NewResponse(c, http.StatusOK, "success", "data wallet histories", walletHistoryData)
}

func (ctrl *WalletHistoryController) GetCashInCashOutMonthly(c echo.Context) error {
	idUser := middlewares.GetUserID(c)
	profile := ctrl.userUsecase.Profile(idUser)
	CashInCashOutData := ctrl.walletHistoryUsecase.GetCashInCashOutMonthly(profile.Wallet.NoWallet)

	return controllers.NewResponse(c, http.StatusOK, "success", "data cash-in cash-out", response.FromDomainCashInCashOut(CashInCashOutData ))
}

func (ctrl *WalletHistoryController) GetWalletHistoriesByUserID(c echo.Context) error {
	idUser := c.Param("user_id")
	profile := ctrl.userUsecase.Profile(idUser)
	walletHistoryData := ctrl.walletHistoryUsecase.GetWalletHistories(profile.Wallet.NoWallet)

	return controllers.NewResponse(c, http.StatusOK, "success", "data wallet histories", walletHistoryData)
}

func (ctrl *WalletHistoryController) GetDetailWalletHistories(c echo.Context) error {
	idHistory := c.Param("history_wallet_id")
	walletHistoryData := ctrl.walletHistoryUsecase.GetDetailWalletHistories(idHistory)
	if walletHistoryData.HistoryWalletID == 0 {
		return controllers.NewResponse(c, http.StatusNotFound, "failed", "wallet history not found", "")
	}

	return controllers.NewResponse(c, http.StatusOK, "success", "data wallet history", response.FromDomain(walletHistoryData))
}

func (ctrl *WalletHistoryController) UpdateWalletHistories(c echo.Context) error {
	input := request.WalletHistory{}
	idHistory := c.Param("history_wallet_id")
	if err := c.Bind(&input); err != nil {
		return controllers.NewResponse(c, http.StatusBadRequest, "failed", "invalid request", "")
	}
	if err := input.Validate(); err != nil {
		return controllers.NewResponse(c, http.StatusBadRequest, "failed", "validation failed", "")
	}
	walletHistory := ctrl.walletHistoryUsecase.UpdateWalletHistories(idHistory, input.ToDomain())
	if walletHistory.HistoryWalletID == 0 {
		return controllers.NewResponse(c, http.StatusNotFound, "failed", "wallet history not found", "")
	}

	return controllers.NewResponse(c, http.StatusOK, "success", "data wallet history updated", response.FromDomain(walletHistory))
}

func (ctrl *WalletHistoryController) CreateWalletHistories(c echo.Context) error {

	NoWallet := c.Param("no_wallet")
	input := request.WalletHistory{}
	if err := c.Bind(&input); err != nil {
		return controllers.NewResponse(c, http.StatusBadRequest, "failed", "invalid request", "")
	}
	if err := input.Validate(); err != nil {
		return controllers.NewResponse(c, http.StatusBadRequest, "failed", "validation failed", "")
	}
	walletHistory := ctrl.walletHistoryUsecase.CreateManual(NoWallet, input.ToDomain())
	if walletHistory.HistoryWalletID == 0 {
		return controllers.NewResponse(c, http.StatusNotFound, "failed", "wallet history not found", "")
	}

	return controllers.NewResponse(c, http.StatusOK, "success", "data wallet history created", response.FromDomain(walletHistory))
}

func (ctrl *WalletHistoryController) DeleteWalletHistories(c echo.Context) error {
	idHistory := c.Param("history_wallet_id")

	if isSuccess := ctrl.walletHistoryUsecase.DeleteWalletHistories(idHistory); !isSuccess {
		return controllers.NewResponse(c, http.StatusNotFound, "failed", "wallet history not found", "")
	}
	return controllers.NewResponse(c, http.StatusOK, "success", "data wallet history deleted", "")

}
