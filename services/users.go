package services

import (
	"gin-learning/common"

	"github.com/gin-gonic/gin"
)

type User struct {
	Id            string `json:"id"`
	ContactNumber string `json:"contact_number"`
}

func GetUsers(c *gin.Context) *common.GonicResonse {
	u := User{
		Id:            "Arkose101",
		ContactNumber: "9893729155",
	}
	return &common.GonicResonse{
		Data:    u,
		Message: "successfull",
	}
}
