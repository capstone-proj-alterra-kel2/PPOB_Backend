package controllers

import "github.com/labstack/echo/v4"

type Response[T any] struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Data    T      `json:"data"`
}

type ResponseFail struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

func NewResponse[T any](c echo.Context, statusCode int, statusMessage string, message string, data T) error {
	return c.JSON(statusCode, Response[T] {
		Status: statusMessage,
		Message: message,
		Data: data,
	})
}

func NewResponseFail(c echo.Context, statusCode int, statusMessage string, message string) error {
	return c.JSON(statusCode, ResponseFail {
		Status: statusMessage,
		Message: message,
	})
}
