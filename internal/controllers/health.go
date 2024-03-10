package controllers

import "github.com/labstack/echo/v4"

func Health(c echo.Context) error {
	return c.JSON(200, "OK")
}
