package controller

import (
	"fmt"
	"net/http"

	"github.com/pilinux/gorest/database"
	"github.com/pilinux/gorest/database/model"

	"github.com/gin-gonic/gin"
)

// GetHobbies - GET /hobbies
func GetHobbies(c *gin.Context) {
	db := database.GetDB()
	hobbies := []model.Hobby{}

	if err := db.Find(&hobbies).Error; err != nil {
		fmt.Println(err)
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, hobbies)
	}
}
