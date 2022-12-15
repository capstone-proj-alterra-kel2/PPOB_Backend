package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	_userUseCase "PPOB_BACKEND/businesses/users"
	_userController "PPOB_BACKEND/controllers/users"

	_paymentmethodUsecase "PPOB_BACKEND/businesses/payment_method"
	_paymentmethodController "PPOB_BACKEND/controllers/payment_method"
	_providerUseCase "PPOB_BACKEND/businesses/providers"
	_providerController "PPOB_BACKEND/controllers/providers"

	_productTypeUseCase "PPOB_BACKEND/businesses/producttypes"
	_productTypeController "PPOB_BACKEND/controllers/producttypes"

	_productUseCase "PPOB_BACKEND/businesses/products"
	_productController "PPOB_BACKEND/controllers/products"

	_walletUseCase "PPOB_BACKEND/businesses/wallets"
	_walletController "PPOB_BACKEND/controllers/wallets"

	_walletHistoryUseCase "PPOB_BACKEND/businesses/wallet_histories"
	_walletHistoryController "PPOB_BACKEND/controllers/wallet_histories"

	_transactionUseCase "PPOB_BACKEND/businesses/transactions"
	_transactionController "PPOB_BACKEND/controllers/transactions"

	_categoryUseCase "PPOB_BACKEND/businesses/category"
	_categoryController "PPOB_BACKEND/controllers/category"

	_driverFactory "PPOB_BACKEND/drivers"

	_middleware "PPOB_BACKEND/app/middlewares"
	_routes "PPOB_BACKEND/app/routes"
	_dbDriver "PPOB_BACKEND/drivers/postgresql"

	util "PPOB_BACKEND/utils"

	"github.com/labstack/echo/v4"
)

type operation func(ctx context.Context) error

func main() {
	configDB := _dbDriver.ConfigDB{
		DB_USERNAME: util.GetEnv("DB_USERNAME"),
		DB_PASSWORD: util.GetEnv("DB_PASSWORD"),
		DB_HOST:     util.GetEnv("DB_HOST"),
		DB_PORT:     util.GetEnv("DB_PORT"),
		DB_NAME:     util.GetEnv("DB_NAME"),
	}

	db := configDB.InitDB()
	_dbDriver.DBMigrate(db)

	configJWT := _middleware.ConfigJWT{
		SecretJWT:      util.GetEnv("JWT_SECRET_KEY"),
		ExpireDuration: 24,
	}

	configLogger := _middleware.ConfigLogger{
		Format: "[${time_rfc3339}] method=${method}, uri=${uri}, status=${status}, latency_human=${latency_human}\n",
	}

	e := echo.New()

	// User
	userRepo := _driverFactory.NewUserRepository(db)
	userUseCase := _userUseCase.NewUserUseCase(userRepo, &configJWT)
	userCtrl := _userController.NewUserController(userUseCase)

	// Product
	productRepo := _driverFactory.NewProductRepository(db)
	productUseCase := _productUseCase.NewProductUseCase(productRepo)
	productCtrl := _productController.NewProductController(productUseCase)

	// Provider
	providerRepo := _driverFactory.NewProviderRepository(db)
	providerUsecase := _providerUseCase.NewProviderUseCase(providerRepo)
	providerCtrl := _providerController.NewProviderController(providerUsecase)
	// Wallet History
	walletHistoryRepo := _driverFactory.NewWalletHistoryRepository(db)
	walletHistoryUseCase := _walletHistoryUseCase.NewWalletHistoryUseCase(walletHistoryRepo)
	walletHistoryCtrl := _walletHistoryController.NewWalletHistoryController(walletHistoryUseCase, userUseCase)
	// Wallet
	walletRepo := _driverFactory.NewWalletRepository(db)
	walletUseCase := _walletUseCase.NewWalletUseCase(walletRepo)
	walletCtrl := _walletController.NewWalletController(walletUseCase, walletHistoryUseCase)
	// Transaction
	transactionRepo := _driverFactory.NewTransactionRepository(db)
	transactionUsecase := _transactionUseCase.NewTransactionUsecase(transactionRepo)
	transactionCtrl := _transactionController.NewTransactionController(transactionUsecase, productUseCase, userUseCase, walletUseCase, walletHistoryUseCase)

	// Product Type
	productTypeRepo := _driverFactory.NewProductTypeRepository(db)
	productTypeUseCase := _productTypeUseCase.NewProductTypeUseCase(productTypeRepo)

	// Product

	// Voucher

	// Payment Method
	paymentmethodRepo := _driverFactory.NewPaymentMethodRepository(db)
	paymentmethodUsecase := _paymentmethodUsecase.NewPaymentMethodUsecase(paymentmethodRepo)
	paymentmethodController := _paymentmethodController.NewPaymentController(paymentmethodUsecase)

	productTypeCtrl := _productTypeController.NewProductTypeController(productTypeUseCase)

	// Category
	categoryRepo := _driverFactory.NewCategoryRepository(db)
	categoryUseCase := _categoryUseCase.NewCategoryUseCase(categoryRepo)
	categoryCtrl := _categoryController.NewCategoryController(categoryUseCase)

	routesInit := _routes.ControllerList{
		LoggerMiddleware:        configLogger.Init(),
		JWTMIddleware:           configJWT.Init(),
		UserController:          *userCtrl,
		WalletHistoryController: *walletHistoryCtrl,
		WalletController:        *walletCtrl,
		ProviderController:      *providerCtrl,
		ProductTypeController:   *productTypeCtrl,
		ProductController:       *productCtrl,
		TransactionController:   *transactionCtrl,
		PaymentController: *paymentmethodController,
		CategoryController:      *categoryCtrl,
	}
	routesInit.RouteRegister(e)

	go func() {
		if err := e.Start(":" + util.GetEnv("PORT")); err != nil && err != http.ErrServerClosed {
			e.Logger.Fatal("shutting down server")
		}
	}()
	wait := gracefulShutdown(context.Background(), 2*time.Second, map[string]operation{
		"database": func(ctx context.Context) error {
			return _dbDriver.CloseDB(db)
		},
		"http-server": func(ctx context.Context) error {
			return e.Shutdown(context.Background())
		},
	})

	<-wait

}

// gracefulShutdown performs gracefully shutdown
func gracefulShutdown(ctx context.Context, timeout time.Duration, ops map[string]operation) <-chan struct{} {
	wait := make(chan struct{})
	go func() {
		s := make(chan os.Signal, 1)

		signal.Notify(s, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)
		<-s

		log.Println("shutting down")

		timeoutFunc := time.AfterFunc(timeout, func() {
			log.Printf("timeout %d ms has been elased, force exit", timeout.Milliseconds())
			os.Exit(0)
		})

		defer timeoutFunc.Stop()

		var wg sync.WaitGroup

		for key, op := range ops {
			wg.Add(1)
			innerOp := op
			innerKey := key
			go func() {
				defer wg.Done()

				log.Printf("cleaning up: %s", innerKey)
				if err := innerOp(ctx); err != nil {
					log.Printf("%s: clean up failed : %s", innerKey, err.Error())
					return
				}

				log.Printf("%s was shutdown gracefully", innerKey)
			}()
		}
		wg.Wait()

		close(wait)
	}()

	return wait
}
