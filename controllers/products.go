package controllers

import (
	"net/http"

	"gin-learning/common"
	"gin-learning/services"

	"github.com/gin-gonic/gin"
)

func GetProduct(c *gin.Context) {
	data := common.GonicResonse{
		Data:    services.GetProduct(c),
		Message: "product fetch successfully",
	}
	c.JSON(http.StatusOK, data)
}

func GetProducts(c *gin.Context) {
	data := common.GonicResonse{
		Data:    services.GetProducts(c),
		Message: "products fetched succesfully",
	}
	c.JSON(http.StatusOK, data)
}
