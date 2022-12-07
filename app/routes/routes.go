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
	// CORS
	e.Use(middleware.CORS())
	v1 := e.Group("/v1")
	auth := v1.Group("/auth")
	// Login
	auth.POST("/login", cl.UserController.Login)
	// SignUp
	auth.POST("/register", cl.UserController.Register)
	// Only Admin & Superadmin
	usersAdmin := v1.Group("/admin/users", middleware.JWTWithConfig(cl.JWTMIddleware), middlewares.IsAdmin)
	usersAdmin.Use(middlewares.CheckStatusToken)
	usersAdmin.GET("", cl.UserController.GetAll)                  // Get All User
	usersAdmin.POST("", cl.UserController.CreateUser)             // Create User
	usersAdmin.PUT("/:user_id", cl.UserController.UpdateDataUser) // Update Data User
	usersAdmin.DELETE("/:user_id", cl.UserController.DeleteUser)  // Delete User
	usersAdmin.GET("/:user_id", cl.UserController.DetailUser)     // Get Detail User
	// Only Superadmin
	adminSuperAdmin := v1.Group("/admin/admins", middleware.JWTWithConfig(cl.JWTMIddleware), middlewares.IsSuperAdmin)
	adminSuperAdmin.Use(middlewares.CheckStatusToken)
	adminSuperAdmin.GET("", cl.UserController.GetAllAdmin)              // Get All Admins
	adminSuperAdmin.POST("", cl.UserController.CreateAdmin)             // Create Admin
	adminSuperAdmin.PUT("/:user_id", cl.UserController.UpdateDataAdmin) // Update Data Admin
	adminSuperAdmin.DELETE("/:user_id", cl.UserController.DeleteAdmin)  // Delete Admin
	adminSuperAdmin.GET("/:user_id", cl.UserController.DetailAdmin)     // Get Detaul Admin
	// User Profile
	user := v1.Group("/user", middleware.JWTWithConfig(cl.JWTMIddleware))
	user.Use(middlewares.CheckStatusToken)
	user.GET("/profile", cl.UserController.Profile)
	user.PUT("/password", cl.UserController.UpdatePassword)
	user.PUT("/data", cl.UserController.UpdateData)
	user.PUT("/image", cl.UserController.UpdateImage)
	// User - Transaction

	// User - Wallet

	// User - Product
	userProduct := v1.Group("/users/products", middleware.JWTWithConfig(cl.JWTMIddleware))
	userProduct.Use(middlewares.CheckStatusToken)
	userProduct.GET("/:product-id", cl.ProductController.GetOne)

	// User - Product Type
	usersProductType := v1.Group("/users/producttypes")
	usersProductType.GET("", cl.ProductTypeController.GetAll)
	usersProductType.GET("/:product-type-id", cl.ProductTypeController.GetOne)

	// User - Provider
	usersProvider := usersProductType.Group("/:product-type-id/providers", middleware.JWTWithConfig(cl.JWTMIddleware))
	usersProvider.Use(middlewares.CheckStatusToken)
	usersProvider.POST("/phone", cl.ProviderController.GetByPhone)

	// Admin

	// Admin - User

	// Admin - Admin

	// Admin - Product
	adminProduct := v1.Group("/admin/products", middleware.JWTWithConfig(cl.JWTMIddleware))
	adminProduct.Use(middlewares.CheckStatusToken)
	adminProduct.GET("", cl.ProductController.GetAll)
	adminProduct.GET("/:product-id", cl.ProductController.GetOne)
	adminProduct.POST("", cl.ProductController.Create)
	adminProduct.PUT("/:product-id", cl.ProductController.Update)
	adminProduct.DELETE("/:product-id", cl.ProductController.Delete)

	// Admin - Product Type
	adminProductType := v1.Group("/admin/producttypes", middleware.JWTWithConfig(cl.JWTMIddleware))
	adminProductType.Use(middlewares.CheckStatusToken)
	adminProductType.GET("", cl.ProductTypeController.GetAll)
	adminProductType.GET("/:product-type-id", cl.ProductTypeController.GetOne)
	adminProductType.POST("", cl.ProductTypeController.Create)
	adminProductType.PUT("/:product-type-id", cl.ProductTypeController.Update)
	adminProductType.DELETE("/:product-type-id", cl.ProductTypeController.Delete)

	// Admin - Provider
	adminProvider := adminProductType.Group("/:product-type-id/providers")
	adminProvider.Use(middlewares.CheckStatusToken)
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
