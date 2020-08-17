package main

import (
	"github.com/labstack/echo"
	"net/http"
	"web/routes"
)

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	//e.POST("/users", SaveUser)
	e.GET("/users/:id", routes.GetUser)
	//e.PUT("/users/:id", UpdateUser)
	//e.DELETE("/users/:id", DeleteUser)
	e.Logger.Fatal(e.Start(":1323"))
}

