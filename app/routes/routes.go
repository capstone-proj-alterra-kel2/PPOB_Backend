package routes

import (
	"PPOB_BACKEND/app/middlewares"
	"PPOB_BACKEND/controllers/products"
	"PPOB_BACKEND/controllers/producttypes"
	"PPOB_BACKEND/controllers/providers"
	"PPOB_BACKEND/controllers/users"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type ControllerList struct {
	LoggerMiddleware      echo.MiddlewareFunc  // Logger
	JWTMIddleware         middleware.JWTConfig // JWT
	UserController        users.UserController // User
	ProductController     products.ProductController
	ProviderController    providers.ProviderController
	ProductTypeController producttypes.ProductTypeController
	// Admin
	// Businesse
}

func (cl *ControllerList) RouteRegister(e *echo.Echo) {
	// Logger
	e.Use(cl.LoggerMiddleware)
	v1 := e.Group("/v1")
	auth := v1.Group("/auth")
	// Login
	auth.POST("/login", cl.UserController.Login)
	// SignUp
	auth.POST("/register", cl.UserController.Register)
	// User
	usersAdmin := v1.Group("/admin/users", middleware.JWTWithConfig(cl.JWTMIddleware))
	usersAdmin.GET("", cl.UserController.GetAll, middlewares.IsAdmin)
	// User - Transaction

	// User - Wallet

	// User - Product
	userProduct := v1.Group("/users/products", middleware.JWTWithConfig(cl.JWTMIddleware))
	userProduct.GET("/:product-id", cl.ProductController.GetOne)

	// User - Product Type
	usersProductType := v1.Group("/users/producttypes")
	usersProductType.GET("", cl.ProductTypeController.GetAll)
	usersProductType.GET("/:product-type-id", cl.ProductTypeController.GetOne)

	// User - Provider
	usersProvider := usersProductType.Group("/:product-type-id/providers", middleware.JWTWithConfig(cl.JWTMIddleware))
	usersProvider.POST("/phone", cl.ProviderController.GetByPhone)

	// Admin

	// Admin - User

	// Admin - Admin

	// Admin - Product
	adminProduct := v1.Group("/admin/products", middleware.JWTWithConfig(cl.JWTMIddleware))
	adminProduct.GET("", cl.ProductController.GetAll)
	adminProduct.GET("/:product-id", cl.ProductController.GetOne)
	adminProduct.POST("", cl.ProductController.Create)
	adminProduct.PUT("/:product-id", cl.ProductController.Update)
	adminProduct.DELETE("/:product-id", cl.ProductController.Delete)

	// Admin - Product Type
	adminProductType := v1.Group("/admin/producttypes", middleware.JWTWithConfig(cl.JWTMIddleware))
	adminProductType.GET("", cl.ProductTypeController.GetAll)
	adminProductType.GET("/:product-type-id", cl.ProductTypeController.GetOne)
	adminProductType.POST("", cl.ProductTypeController.Create)
	adminProductType.PUT("/:product-type-id", cl.ProductTypeController.Update)
	adminProductType.DELETE("/:product-type-id", cl.ProductTypeController.Delete)

	// Admin - Provider
	adminProvider := adminProductType.Group("/:product-type-id/providers")
	adminProvider.GET("", cl.ProviderController.GetAll)
	adminProvider.GET("/:provider-id", cl.ProviderController.GetOne)
	adminProvider.POST("", cl.ProviderController.Create)
	adminProvider.PUT("/:provider-id", cl.ProviderController.Update)
	adminProvider.DELETE("/:provider-id", cl.ProviderController.Delete)

	// Admin - Voucher

	// Admin - Transaction

	// Admin - Wallet

	// Logout
	withAuth := v1.Group("/auth", middleware.JWTWithConfig(cl.JWTMIddleware))
	withAuth.POST("/logout", cl.UserController.Logout)
}
