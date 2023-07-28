package controllers

import (
	"net/http"

	"prakerja7/config"
	"prakerja7/models"

	"github.com/labstack/echo/v4"
)


func AddProductController(c echo.Context) error{

	var product models.Product
	c.Bind(&product)
	
	result := config.DB.Create(&product)
	
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError,  models.BaseResponse{
			Status: false,
			Message: "Failed create product into database",
			Data: nil,
		})
	}
	
	return c.JSON(http.StatusCreated,  models.BaseResponse{
		Status: true,
		Message: "Success add data product",
		Data: product,
	})
}

func GetDetailProductController(c echo.Context) error{

	// id, _ := strconv.Atoi(c.Param("id"))

	product :=  models.Product{}
	
	return c.JSON(http.StatusOK,  models.BaseResponse{
		Status: true,
		Message: "Berhasil",
		Data: product,
	})
}

func GetProductController(c echo.Context) error{

	var dataProducts [] models.Product
	
	result := config.DB.Find(&dataProducts)
	
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, models. BaseResponse{
			Status: false,
			Message: "Failed get product from database",
			Data: nil,
		})
	}

	return c.JSON(http.StatusOK,  models.BaseResponse{
		Status: true,
		Message: "Berhasil",
		Data: dataProducts,
	})
}