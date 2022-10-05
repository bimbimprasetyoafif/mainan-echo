package helo

import (
	"fmt"
	"github.com/coba/helper"
	"github.com/labstack/echo-contrib/jaegertracing"
	"github.com/labstack/echo/v4"
)

func HandlerHello(c echo.Context) error {
	sp := jaegertracing.CreateChildSpan(c, "handler-hello-v1")
	defer sp.Finish()

	fmt.Println(c.Request().Header)

	return c.String(200, "ini halo")
}

func HandlerLogin(c echo.Context) error {
	sp := jaegertracing.CreateChildSpan(c, "handler-login-v1")
	defer sp.Finish()

	req := c.Request()
	username := req.Header.Get("username")
	password := req.Header.Get("password")

	// logic ke db cek user

	token, err := helper.CreateJwt(c, username+password)
	if err != nil {
		return c.JSON(500, err.Error())
	}

	return c.JSON(200, map[string]string{
		"token": token,
	})
}
