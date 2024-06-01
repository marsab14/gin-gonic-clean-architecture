package routes

import (
	"gin-learning/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRouters(r *gin.Engine) {
	productApi := r.Group("/product")
	{
		productApi.GET("/", controllers.GetProduct)
		productApi.GET("/all", controllers.GetProducts)
	}
}
