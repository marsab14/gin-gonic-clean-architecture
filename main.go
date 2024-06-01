package main

import (
	"fmt"
	"net/http"

	"gin-learning/common"
	"gin-learning/routes"
	"gin-learning/services"

	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	r := gin.Default()

	return r
}

type Person struct {
	Name string
	Age  int
}

func TestCustomMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Print("Inside custom middleware")
	}
}

func main() {
	r := setupRouter()
	// custom middleware
	r.Use(TestCustomMiddleware())
	r.GET("/ping", func(c *gin.Context) {
		p := Person{
			Name: "John",
			Age:  25,
		}
		response := &common.GonicResonse{
			Data:    p,
			Message: "Successfull",
		}
		c.JSON(http.StatusOK, response)
	})
	r.GET("/users", func(c *gin.Context) {
		data := services.GetUsers(c)
		c.JSON(http.StatusOK, data)
	})
	routes.SetupRouters(r)
	r.Run(":3201")
}
