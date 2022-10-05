package helper

import (
	"github.com/coba/config"
	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo-contrib/jaegertracing"
	"github.com/labstack/echo/v4"
	"time"
)

func CreateJwt(c echo.Context, username string) (string, error) {
	sp := jaegertracing.CreateChildSpan(c, "helper-create-token")
	defer sp.Finish()

	mapClaim := jwt.MapClaims{}
	mapClaim["username"] = username
	mapClaim["exp"] = time.Now().Add(40 * time.Second).Unix()

	generateToken := jwt.NewWithClaims(jwt.SigningMethodHS256, mapClaim)
	return generateToken.SignedString([]byte(config.Cfg.TokenSecret))
}
