package routes

import (
	"PPOB_BACKEND/app/middlewares"
	"PPOB_BACKEND/controllers/users"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type ControllerList struct {
	LoggerMiddleware echo.MiddlewareFunc  // Logger
	JWTMIddleware    middleware.JWTConfig // JWT
	UserController   users.UserController // User
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
	usersAdmin.GET("", cl.UserController.GetAll)                  // Get All User
	usersAdmin.POST("", cl.UserController.CreateUser)             // Create User
	usersAdmin.PUT("/:user_id", cl.UserController.UpdateDataUser) // Update Data User
	usersAdmin.DELETE("/:user_id", cl.UserController.DeleteUser)  // Delete User
	// Only Superadmin
	adminSuperAdmin := v1.Group("/admin/admins", middleware.JWTWithConfig(cl.JWTMIddleware), middlewares.IsSuperAdmin)
	adminSuperAdmin.GET("", cl.UserController.GetAllAdmin)              // Get All Admins
	adminSuperAdmin.POST("", cl.UserController.CreateAdmin)             // Create Admin
	adminSuperAdmin.PUT("/:user_id", cl.UserController.UpdateDataAdmin) // Update Data Admin
	adminSuperAdmin.DELETE("/user_id", cl.UserController.DeleteAdmin)   // Delete Admin
	// User Profile
	user := v1.Group("/user", middleware.JWTWithConfig(cl.JWTMIddleware))
	user.GET("/profile", cl.UserController.Profile)
	user.PUT("/password", cl.UserController.UpdatePassword)
	user.PUT("/data", cl.UserController.UpdateData)
	user.PUT("/image", cl.UserController.UpdateImage)
	// User - Transaction

	// User - Wallet

	// User - Product Type

	// User - Provider

	// Admin

	// Admin - User

	// Admin - Admin

	// Admin - Product Type

	// Admin - Provider

	// Admin - Voucher

	// Admin - Transaction

	// Admin - Wallet

	// Logout
	withAuth := v1.Group("/auth", middleware.JWTWithConfig(cl.JWTMIddleware))
	withAuth.POST("/logout", cl.UserController.Logout)
}
