package middleware

import (
	"github.com/labstack/echo/v4"
)

func Login(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Set("uid", 114514)
		return next(c)
	}
}
