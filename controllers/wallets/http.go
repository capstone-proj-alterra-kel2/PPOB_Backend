package wallets

import (
	"PPOB_BACKEND/app/middlewares"
	"PPOB_BACKEND/businesses/wallets"
	"PPOB_BACKEND/controllers"
	"PPOB_BACKEND/controllers/wallets/request"
	"PPOB_BACKEND/controllers/wallets/response"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/morkid/paginate"
)

type WalletController struct {
	walletUsecase wallets.Usecase
}

func NewWalletController(walletUC wallets.Usecase) *WalletController {
	return &WalletController{
		walletUsecase: walletUC,
	}
}

func (ctrl *WalletController) GetWalletUser(c echo.Context) error {
	idUser := middlewares.GetUserID(c)
	walletData := ctrl.walletUsecase.GetWalletUser(idUser)

	if walletData.NoWallet == "" {
		return controllers.NewResponseFail(c, http.StatusNotFound, "failed", " wallet not found")
	}
	return controllers.NewResponse(c, http.StatusOK, "success", "data wallet", walletData)
}

func (ctrl *WalletController) GetWalletUserByUserID(c echo.Context) error {
	idUser := c.Param("user_id")
	walletData := ctrl.walletUsecase.GetWalletUser(idUser)

	if walletData.NoWallet == "" {
		return controllers.NewResponseFail(c, http.StatusNotFound, "failed", "wallet not found")
	}
	return controllers.NewResponse(c, http.StatusOK, "success", "data wallet", walletData)
}

func (ctrl *WalletController) GetAllWallet(c echo.Context) error {
	pg := paginate.New()
	size, _ := strconv.Atoi(c.QueryParam("size"))
	page, _ := strconv.Atoi(c.QueryParam("page"))
	sort := c.QueryParam("sort")
	walletsData, walletDomain := ctrl.walletUsecase.GetAllWallet(page, size, sort)
	wallets := []response.Wallet{}
	for _, wallet := range walletDomain {
		wallets = append(wallets, response.FromDomain(wallet))
	}
	return controllers.NewResponse(c, http.StatusOK, "success", "all wallet", pg.Response(walletsData, c.Request(), &wallets))
}

func (ctrl *WalletController) GetDetailWallet(c echo.Context) error {
	noWallet := c.Param("no_wallet")
	wallet := ctrl.walletUsecase.GetDetailWallet(noWallet)

	if wallet.UserID == 0 {
		return controllers.NewResponseFail(c, http.StatusNotFound, "failed", "cannot get detail wallet")
	}
	return controllers.NewResponse(c, http.StatusOK, "success", "get detail wallet", response.FromDomain(wallet))
}

func (ctrl *WalletController) UpdateBalance(c echo.Context) error {
	idUser := middlewares.GetUserID(c)
	input := request.UpdateBalance{}

	if err := c.Bind(&input); err != nil {
		return controllers.NewResponseFail(c, http.StatusBadRequest, "failed", "invalid request")
	}
	if err := input.Validate(); err != nil {
		return controllers.NewResponseFail(c, http.StatusBadRequest, "failed", "validation failed")
	}
	user, err := ctrl.walletUsecase.UpdateBalance(idUser, input.ToDomain())
	if err != nil {
		return controllers.NewResponseFail(c, http.StatusBadRequest, "failed", err.Error())
	}

	return controllers.NewResponse(c, http.StatusOK, "success", "image updated", response.FromDomain(user))
}

func (ctrl *WalletController) IsiSaldo(c echo.Context) error {
	idUser := middlewares.GetUserID(c)
	input := request.UpdateBalance{}

	if err := c.Bind(&input); err != nil {
		return controllers.NewResponseFail(c, http.StatusBadRequest, "failed", "invalid request")
	}
	if err := input.Validate(); err != nil {
		return controllers.NewResponseFail(c, http.StatusBadRequest, "failed", "validation failed")
	}
	user, err := ctrl.walletUsecase.IsiSaldo(idUser, input.ToDomain())
	if err != nil {
		return controllers.NewResponseFail(c, http.StatusBadRequest, "failed", err.Error())
	}

	return controllers.NewResponse(c, http.StatusOK, "success", "image updated", response.FromDomain(user))

}
