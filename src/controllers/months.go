package controllers

import (
	"gingonic/src/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetAllMonth(c *gin.Context) {
	res := models.SelectAllMonth()
	c.JSON(200, gin.H{
		"Message": "Success",
		"Data": res,
	})
}

func GetMonthById(c *gin.Context) {
	idParam := c.Param("id")
	id, _ := strconv.Atoi(idParam)
	res := models.SelectMonthById(id)
	c.JSON(200, gin.H{
		"Message": "Success",
		"data": res,
	})
}

func CreateMonth(c *gin.Context) {
	var input models.Month
	err := c.ShouldBindJSON(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	models.CreateMonth(input.Name, input.Day)

	c.JSON(200, gin.H{
		"Message": "Month Created",
	})
}

func UpdateMonth(c *gin.Context) {
	idParam := c.Param("id")
	id, _ := strconv.Atoi(idParam)
	var input models.Month
	err := c.ShouldBindJSON(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	models.UpdateMonth(id, input.Name, input.Day)

	c.JSON(200, gin.H{
		"Message": "Month Updated",
	})
}

func DeleteMonth(c *gin.Context) {
	idParam := c.Param("id")
	id, _ := strconv.Atoi(idParam)
	models.DeleteMonth(id)

	c.JSON(200, gin.H{
		"Message": "Month Deleted",
	})
}