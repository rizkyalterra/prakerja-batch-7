package routes

import (
	"prakerja7/controllers"

	"github.com/labstack/echo/v4"
)

func InitRoute(e *echo.Echo) *echo.Echo {
	e.GET("/products", controllers.GetProductController)
	e.POST("/products", controllers.AddProductController)
	e.GET("/products/:id", controllers.GetDetailProductController)
	return e
}