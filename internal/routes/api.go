package routes

import (
	"crypto/subtle"

	"github.com/dusanbre/goPgsqlDocker/internal/controllers"
	"github.com/labstack/echo/v4"
)

func baseAuth(username, password string, c echo.Context) (bool, error) {
	if subtle.ConstantTimeCompare([]byte(username), []byte("joe")) == 1 &&
		subtle.ConstantTimeCompare([]byte(password), []byte("secret")) == 1 {
		return true, nil
	}
	return false, nil
}

func Routes(r *echo.Echo) {

	base := r.Group("/api/v1")

	base.POST("/login", controllers.Login)

	base.GET("/heat", controllers.Health)

}
