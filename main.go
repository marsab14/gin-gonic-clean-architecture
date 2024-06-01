package main

import (
	"fmt"
	"net/http"
	"os"

	"gin-learning/common"
	"gin-learning/config"
	"gin-learning/routes"

	"github.com/gin-gonic/gin"
	"github.com/kelseyhightower/envconfig"
)

func setupRouter() *gin.Engine {
	r := gin.Default()
	return r
}

type Person struct {
	Name string
	Age  int
}

func stopForcefullyWithError(msg string, err error) {
	fmt.Printf("%s: %v\n", msg, err.Error())
	os.Exit(1)
}

func TestCustomMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Print("Inside custom middleware")
	}
}

func main() {
	// validating configuration
	var s config.Specification
	err := envconfig.Process("ig", &s)
	if err != nil {
		stopForcefullyWithError("Error processing configuration ", err)
	}
	err = config.ValidateSpecification(s)
	if err != nil {
		stopForcefullyWithError("Invalid Configuration", err)
	}
	// setting up router
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

	routes.SetupRouters(r)
	r.Run(":" + string(s.Port))
}
