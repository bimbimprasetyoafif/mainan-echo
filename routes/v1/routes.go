package v1

import (
	"fmt"
	"github.com/coba/config"
	c "github.com/coba/constant"
	"github.com/coba/controllers/rest-echo/helo"
	usersEcho "github.com/coba/controllers/rest-echo/users"
	usersHttp "github.com/coba/controllers/rest-http/users"
	"github.com/coba/routes"
	"github.com/labstack/echo-contrib/jaegertracing"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"io"
	"net/http"
)

func InitRoute(payload *routes.Payload) (*echo.Echo, io.Closer) {
	e := echo.New()

	e.Use(middleware.Logger())
	trace := jaegertracing.New(e, nil)

	v1 := e.Group("/v1")
	v1.POST("/login", helo.HandlerLogin)
	//v1.Use(MiddlewareAPIKey)

	h := v1.Group(c.HelloIndex, middleware.CORSWithConfig(
		middleware.CORSConfig{
			AllowOrigins: []string{"https://localhost:9999"},
			//AllowMethods: []string{"GET"},
		}))
	h.GET("", helo.HandlerHello)

	hUser := usersEcho.HandlerUser{
		UserService: payload.GetUserService(),
	}

	u := v1.Group(c.UsersIndex)
	u.GET("", hUser.HandlerUsersAll)
	u.POST("", usersEcho.HandlerCreateUsers)
	u.GET("/:id", usersEcho.HandlerUsersByID)
	u.DELETE("/:id", usersEcho.HandlerDeleteUsersByID)
	u.PUT("/:id", usersEcho.HandlerUpdateUsers)

	return e, trace
}

func MidBasicAuth(username string, password string, c echo.Context) (bool, error) {
	isAdmin := false

	if username == "admin" && password == "admin" {
		isAdmin = true
	}

	return isAdmin, nil
}

func MidEcho(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		fmt.Println("ini log echo")

		return next(c)
	}
}

func MiddlewareAPIKey(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {

		apiKey := c.Request().Header.Get("x-api-key")

		if len(apiKey) < 1 {
			return c.JSON(400, map[string]string{
				"message": "your request body is awesome",
			})
		}

		if apiKey != config.Cfg.APIKey {
			return c.JSON(401, map[string]string{
				"message": "unauthorized key",
			})
		}

		c.Request().Header.Del("x-api-key")

		return next(c)
	}
}

func InitHttpRoute() {
	http.HandleFunc("/hello", usersHttp.HandlerUsersAll)
}
