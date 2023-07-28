package routes

import (
	"os"
	"prakerja7/controllers"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func InitRoute(e *echo.Echo) *echo.Echo {
	e.Use(middleware.Logger())
	e.POST("/login", controllers.LoginController)

	productRoute := e.Group("")
	productRoute.Use(echojwt.JWT([]byte(os.Getenv("SECRET_KEY"))))
	productRoute.GET("/products", controllers.GetProductController)
	productRoute.POST("/products", controllers.AddProductController)
	productRoute.GET("/products/:id", controllers.GetDetailProductController)
	return e
}