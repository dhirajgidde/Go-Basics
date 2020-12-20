package main

import (
	"net/http"

	"github.com/labstack/echo"
)

// e.GET("/users/:id", getUser)
func getUser(c echo.Context) error {
	id := c.Param("id")
	return c.String(http.StatusOK, id)
}

// func saveUser(c echo.Context) error {
// 	name := c.FormValue("Name")
// 	email := c.FormValue("Email")
// 	return c.String(http.StatusOK, "Name : "+name+" Email: "+email)
// }

func main() {
	e := echo.New()
	e.GET("/users/:id", getUser)
	//e.POST("/users", saveUser)
	e.Logger.Fatal(e.Start(":9099"))

}
