package services

import (
	"github.com/gin-gonic/gin"
)

type Product struct {
	ProductId   string `json:"product_id"`
	ProductName string `json:"product_name"`
	Description string `json:"product_description"`
	ModelNumber string `json:"model_number"`
}

func GetProduct(c *gin.Context) *Product {
	p := &Product{
		ProductId:   "prd__011",
		ProductName: "Water Bottle 101",
		Description: "just a simple water bottler",
		ModelNumber: "wtr__001",
	}
	return p
}

func GetProducts(c *gin.Context) []Product {
	products := []Product{
		{
			ProductId:   "prd__011",
			ProductName: "Water Bottle 101",
			Description: "just a simple water bottler",
			ModelNumber: "wtr__001",
		},
		{
			ProductId:   "prd__012",
			ProductName: "Clay water glass",
			Description: "Water glass made of clay , just a bit broken",
			ModelNumber: "wtr__002",
		},
	}
	return products
}
