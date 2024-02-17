package controllers

import (
	"gingonic/src/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetAllDay(c *gin.Context) {
	res := models.SelectAllDay()
	c.JSON(200, gin.H{
		"Message": "Success",
		"data": res,
	})
}

func GetDayById(c *gin.Context) {
	idParam := c.Param("id")
	id, _ := strconv.Atoi(idParam)
	res := models.SelectDayById(id)
	c.JSON(200, gin.H{
		"Message": "Success",
		"data": res,
	})
}

func CreateDay(c *gin.Context) {
	var input models.Day
	err := c.ShouldBindJSON(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	models.CreateDay(input.Name)
	c.JSON(200, gin.H{
		"message": "Day Created",
	})
}

func UpdateDay(c *gin.Context) {
	idParam := c.Param("id")
	id, _ := strconv.Atoi(idParam)
	var input models.Day
	err := c.ShouldBindJSON(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	models.UpdateDay(id, input.Name)
	c.JSON(200, gin.H{
		"message": "Day Updated",
	})
}

func DeleteDay(c *gin.Context) {
	idParam := c.Param("id")
	id, _ := strconv.Atoi(idParam)
	models.DeleteDay(id)
	c.JSON(200, gin.H{
		"message": "Day Deleted",
	})
}