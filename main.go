package main

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type Products struct {
	Id int `json:"id"`
	Name string `json:"name"`
	Price int `json:"price"`
}

type BaseResponse struct {
	Status bool `json:"status"`
	Message string `json:"message"`
	Data interface{} `json:"data"`
}

func main(){
	e := echo.New()
	e.GET("/products", GetProductController)
	e.POST("/products", AddProductController)
	e.GET("/products/:id", GetDetailProductController)
	e.Start(":8000")
}

func AddProductController(c echo.Context) error{

	var product Products
	c.Bind(&product)
	
	return c.JSON(http.StatusCreated, BaseResponse{
		Status: true,
		Message: "Berhasil Ditambahkan",
		Data: product,
	})
}

func GetDetailProductController(c echo.Context) error{

	id, _ := strconv.Atoi(c.Param("id"))
	
	product := Products{id, "Baju", 10000}
	
	return c.JSON(http.StatusOK, BaseResponse{
		Status: true,
		Message: "Berhasil",
		Data: product,
	})
}

func GetProductController(c echo.Context) error{
	country := c.QueryParam("country")
	harga := c.QueryParam("harga")
	
	var dataProducts []Products
	
	product := Products{1, country, 10000}
	dataProducts = append(dataProducts, product)
	product = Products{2, harga, 10000}
	dataProducts = append(dataProducts, product)
	
	return c.JSON(http.StatusOK, BaseResponse{
		Status: true,
		Message: "Berhasil",
		Data: dataProducts,
	})
}
