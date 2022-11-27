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
	v1 := e.Group("/v1")
	auth := v1.Group("/auth")
	// Login
	auth.POST("/login", cl.UserController.Login)
	// SignUp
	auth.POST("/register", cl.UserController.Register)
	// User
	usersAdmin := v1.Group("/admin/users",middleware.JWTWithConfig(cl.JWTMIddleware))
	usersAdmin.GET("", cl.UserController.GetAll, middlewares.IsAdmin)

	// User Profile
	user:= v1.Group("/user",middleware.JWTWithConfig(cl.JWTMIddleware))
	user.GET("/profile", cl.UserController.Profile)
 user.PUT("/password", cl.UserController.UpdatePassword)
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
