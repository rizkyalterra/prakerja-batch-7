package main

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
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
	connectDatabase()
	e := echo.New()
	e.GET("/products", GetProductController)
	e.POST("/products", AddProductController)
	e.GET("/products/:id", GetDetailProductController)
	e.Start(":8000")
}

var DB *gorm.DB

func connectDatabase(){
	dsn := "root:123ABC4d.@tcp(127.0.0.1:3306)/prakerja7?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed connect into database")
	}
	migration()
}

func migration(){
	DB.AutoMigrate(&Products{})

}

func AddProductController(c echo.Context) error{

	var product Products
	c.Bind(&product)

	result := DB.Create(&product)
	
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, BaseResponse{
			Status: false,
			Message: "Failed create product into database",
			Data: nil,
		})
	}
	
	return c.JSON(http.StatusCreated, BaseResponse{
		Status: true,
		Message: "Success add data product",
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

	var dataProducts []Products
	
	result := DB.Find(&dataProducts)
	
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, BaseResponse{
			Status: false,
			Message: "Failed get product from database",
			Data: nil,
		})
	}

	return c.JSON(http.StatusOK, BaseResponse{
		Status: true,
		Message: "Berhasil",
		Data: dataProducts,
	})
}
