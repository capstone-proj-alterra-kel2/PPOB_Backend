package routes

import (
	"PPOB_BACKEND/app/middlewares"
	"PPOB_BACKEND/controllers/products"
	"PPOB_BACKEND/controllers/producttypes"
	"PPOB_BACKEND/controllers/providers"
	"PPOB_BACKEND/controllers/transactions"
	"PPOB_BACKEND/controllers/users"
	"PPOB_BACKEND/controllers/wallet_histories"
	"PPOB_BACKEND/controllers/wallets"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type ControllerList struct {
	LoggerMiddleware        echo.MiddlewareFunc  // Logger
	JWTMIddleware           middleware.JWTConfig // JWT
	UserController          users.UserController // User
	ProductController       products.ProductController
	ProviderController      providers.ProviderController
	ProductTypeController   producttypes.ProductTypeController
	TransactionController   transactions.TransactionController
	WalletController        wallets.WalletController                 // Wallet
	WalletHistoryController wallet_histories.WalletHistoryController // Wallet Histories
	// Businesse
}

func (cl *ControllerList) RouteRegister(e *echo.Echo) {
	// Logger
	e.Use(cl.LoggerMiddleware)
	// CORS
	e.Use(middleware.CORS())
	v1 := e.Group("/v1")
	auth := v1.Group("/auth")
	// Login
	auth.POST("/login", cl.UserController.Login)
	// SignUp
	auth.POST("/register", cl.UserController.Register)
	auth.POST("/register/check", cl.UserController.CheckDuplicateUser)
	// Only Admin & Superadmin - User
	usersAdmin := v1.Group("/admin/users", middleware.JWTWithConfig(cl.JWTMIddleware), middlewares.IsAdmin)
	usersAdmin.Use(middlewares.CheckStatusToken)
	usersAdmin.GET("", cl.UserController.GetAll)                  // Get All User
	usersAdmin.POST("", cl.UserController.CreateUser)             // Create User
	usersAdmin.PUT("/:user_id", cl.UserController.UpdateDataUser) // Update Data User
	usersAdmin.DELETE("/:user_id", cl.UserController.DeleteUser)  // Delete User
	usersAdmin.GET("/:user_id", cl.UserController.DetailUser)     // Get Detail User
	// Only Superadmin - Admin
	adminSuperAdmin := v1.Group("/admin/admins", middleware.JWTWithConfig(cl.JWTMIddleware), middlewares.IsSuperAdmin)
	adminSuperAdmin.Use(middlewares.CheckStatusToken)
	adminSuperAdmin.GET("", cl.UserController.GetAllAdmin)              // Get All Admins
	adminSuperAdmin.POST("", cl.UserController.CreateAdmin)             // Create Admin
	adminSuperAdmin.PUT("/:admin_id", cl.UserController.UpdateDataAdmin) // Update Data Admin
	adminSuperAdmin.DELETE("/:admin_id", cl.UserController.DeleteAdmin)  // Delete Admin
	adminSuperAdmin.GET("/:admin_id", cl.UserController.DetailAdmin)     // Get Detaul Admin
	// User - User Profile
	user := v1.Group("/user", middleware.JWTWithConfig(cl.JWTMIddleware))
	user.Use(middlewares.CheckStatusToken)
	user.GET("/profile", cl.UserController.Profile)
	user.PUT("/password", cl.UserController.UpdatePassword)
	user.PUT("/data", cl.UserController.UpdateData)
	user.PUT("/image", cl.UserController.UpdateImage)
	user.GET("/wallet", cl.WalletController.GetWalletUser)
	user.GET("/wallet/cashin-cashout", cl.WalletHistoryController.GetCashInCashOutMonthly)
	user.GET("/wallet/histories", cl.WalletHistoryController.GetWalletHistories)
	user.POST("/wallet/isi-saldo", cl.WalletController.IsiSaldo)

	// User - Transaction
	userTransaction := v1.Group("/users/transactions", middleware.JWTWithConfig(cl.JWTMIddleware))
	userTransaction.Use(middlewares.CheckStatusToken)
	userTransaction.GET("/history", cl.TransactionController.GetTransactionHistory)
	userTransaction.GET("/:transaction_id", cl.TransactionController.GetDetail)
	userTransaction.POST("/create", cl.TransactionController.Create)

	// User - Product
	userProduct := v1.Group("/users/products", middleware.JWTWithConfig(cl.JWTMIddleware))
	userProduct.Use(middlewares.CheckStatusToken)
	userProduct.GET("/:product_id", cl.ProductController.GetOne)

	// User - Product Type
	usersProductType := v1.Group("/users/producttypes", middleware.JWTWithConfig(cl.JWTMIddleware), middlewares.CheckStatusToken)
	usersProductType.GET("", cl.ProductTypeController.GetAll)
	usersProductType.GET("/:product_type_id", cl.ProductTypeController.GetOne)

	// User - Provider
	usersProvider := usersProductType.Group("/:product_type_id/providers")
	usersProvider.Use(middlewares.CheckStatusToken)
	usersProvider.POST("/phone", cl.ProviderController.GetByPhone)

	// Admin - Product
	adminProduct := v1.Group("/admin/products", middleware.JWTWithConfig(cl.JWTMIddleware), middlewares.IsAdmin)
	adminProduct.Use(middlewares.CheckStatusToken)
	adminProduct.GET("", cl.ProductController.GetAll)
	adminProduct.GET("/:product_id", cl.ProductController.GetOne)
	adminProduct.POST("", cl.ProductController.Create)
	adminProduct.PUT("/:product_id", cl.ProductController.UpdateData)
	adminProduct.DELETE("/:product_id", cl.ProductController.Delete)

	// Admin - Product Type
	adminProductType := v1.Group("/admin/producttypes", middleware.JWTWithConfig(cl.JWTMIddleware), middlewares.IsAdmin)
	adminProductType.Use(middlewares.CheckStatusToken)
	adminProductType.GET("", cl.ProductTypeController.GetAll)
	adminProductType.GET("/:product_type_id", cl.ProductTypeController.GetOne)
	adminProductType.POST("", cl.ProductTypeController.Create)
	adminProductType.PUT("/:product_type_id", cl.ProductTypeController.Update)
	adminProductType.DELETE("/:product_type_id", cl.ProductTypeController.Delete)

	// Admin - Provider
	adminProvider := adminProductType.Group("/:product_type_id/providers", middleware.JWTWithConfig(cl.JWTMIddleware), middlewares.IsAdmin)
	adminProvider.Use(middlewares.CheckStatusToken)
	adminProvider.GET("", cl.ProviderController.GetAll)
	adminProvider.GET("/:provider_id", cl.ProviderController.GetOne)
	adminProvider.POST("", cl.ProviderController.Create)
	adminProvider.PUT("/:provider_id", cl.ProviderController.Update)
	adminProvider.DELETE("/:provider_id", cl.ProviderController.Delete)

	// Admin - Voucher

	// Admin - Transaction
	adminTransaction := v1.Group("/admin/transactions", middleware.JWTWithConfig(cl.JWTMIddleware), middlewares.IsAdmin)
	adminTransaction.Use(middlewares.CheckStatusToken)
	adminTransaction.GET("", cl.TransactionController.GetAll)

	// Admin - Wallet
	adminWallet := v1.Group("/admin/wallets", middleware.JWTWithConfig(cl.JWTMIddleware), middlewares.CheckStatusToken, middlewares.IsAdmin)
	adminWallet.GET("", cl.WalletController.GetAllWallet)
	adminWallet.GET("/:user_id", cl.WalletController.GetDetailWallet)
	adminWallet.PUT("/:user_id", cl.WalletController.UpdateBalance)
	// Admin - Wallet History
	adminWalletHistory := adminWallet.Group("/:no_wallet/histories")
	adminWalletHistory.GET("", cl.WalletHistoryController.GetWalletHistories)
	adminWalletHistory.GET("/:history_wallet_id", cl.WalletHistoryController.GetDetailWalletHistories)
	adminWalletHistory.PUT("/:history_wallet_id", cl.WalletHistoryController.UpdateWalletHistories)
	adminWalletHistory.DELETE("/:history_wallet_id", cl.WalletHistoryController.DeleteWalletHistories)
	adminWalletHistory.POST("", cl.WalletHistoryController.CreateWalletHistories)

	// Logout
	withAuth := v1.Group("/auth", middleware.JWTWithConfig(cl.JWTMIddleware))

	withAuth.POST("/logout", cl.UserController.Logout)
}
