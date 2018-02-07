package controllers

import (
    "net/http"
	"uktrav_echo/app"
	"github.com/labstack/echo"
)

func Init() {
        app.Server.GET("/users", func(c echo.Context) error{
    //dO something here
     return c.String(http.StatusOK, "OK")
    })
}
