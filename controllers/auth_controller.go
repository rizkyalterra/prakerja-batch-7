package controllers

import (
	"net/http"
	"os"
	"prakerja7/models"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

type jwtCustomClaims struct {
	Name  string `json:"name"`
	Id  int `json:"id"`
	Admin bool   `json:"admin"`
	jwt.RegisteredClaims
}

func LoginController(c echo.Context) error {
	var loginRequest models.LoginRequest
	c.Bind(&loginRequest)

	// database => COCOK

	claims := &jwtCustomClaims{
		"Alterra",
		1,
		true,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 72)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString([]byte(os.Getenv("SECRET_KEY")))
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, models.BaseResponse{
		Status: true,
		Message: "Success login",
		Data: echo.Map{
			"token": t,
		},
	})
}