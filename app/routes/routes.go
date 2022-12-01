package routes

import (
	"PPOB_BACKEND/app/middlewares"
	"PPOB_BACKEND/controllers/users"
	"PPOB_BACKEND/controllers/payment_method"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type ControllerList struct {
	LoggerMiddleware echo.MiddlewareFunc  // Logger
	JWTMIddleware    middleware.JWTConfig // JWT
	UserController   users.UserController // User
	PaymentController payment_method.PaymentController
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
	// User - Transaction
	
	// User - Wallet

	// User - Product Type

	// User - Provider

	// User - Payment Method
	payment := v1.Group("/payments", middleware.JWTWithConfig(cl.JWTMIddleware))
	payment.GET("", cl.PaymentController.GetAll)
	payment.GET("/:id", cl.PaymentController.GetSpecificPayment)
	payment.POST("", cl.PaymentController.CreatePayment, middlewares.IsAdmin)
	payment.PUT("/:id", cl.PaymentController.UpdatePaymentByID, middlewares.IsAdmin)
	payment.DELETE("/:id", cl.PaymentController.DeletePayment, middlewares.IsAdmin)


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
