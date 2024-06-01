package controllers

import (
	"fmt"
	"net/http"

	"gin-learning/common"
	"gin-learning/config"
	"gin-learning/services"

	"github.com/gin-gonic/gin"
)

var s config.Specification

func GetProduct(c *gin.Context) {
	data := common.GonicResonse{
		Data:    services.GetProduct(c),
		Message: "product fetch successfully",
	}
	c.JSON(http.StatusOK, data)
}

func GetProducts(c *gin.Context) {
	fmt.Print("env is", s.Envirnoment)
	fmt.Print("port is", s.Port)
	tempData := map[string]interface{}{
		"products": services.GetProducts(c),
		"env":      s.Envirnoment,
		"port":     s.Port,
	}
	data := common.GonicResonse{
		Data:    tempData,
		Message: "products fetched succesfully updated",
	}
	c.JSON(http.StatusOK, data)
}
