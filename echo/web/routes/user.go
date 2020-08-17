package routes

import (
	"github.com/labstack/echo"
	"net/http"
)

func GetUser(c echo.Context) error {
	id := c.Param("id")
	return c.String(http.StatusOK, id)
}