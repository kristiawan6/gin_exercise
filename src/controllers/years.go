package controllers

import (
	"gingonic/src/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetAllYear(c *gin.Context) {
	res := models.SelectAllYear()
	c.JSON(200, gin.H{
		"Message": "Success",
		"data": res,
	})
}

func GetYearById(c *gin.Context) {
	idParam := c.Param("id")
	id, _ := strconv.Atoi(idParam)
	res := models.SelectYearById(id)
	c.JSON(200, gin.H{
		"Message": "Success",
		"data": res,
	})
}

func CreateYear(c *gin.Context) {
	var input models.Year
	err := c.ShouldBindJSON(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	models.CreateYear(input.Name)
	c.JSON(200, gin.H{
		"message": "Year Created",
	})
}

func UpdateYear(c *gin.Context) {
	idParam := c.Param("id")
	id, _ := strconv.Atoi(idParam)
	var input models.Year
	err := c.ShouldBindJSON(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	models.UpdateYear(id, input.Name)
	c.JSON(200, gin.H{
		"message": "Year Updated",
	})
}

func DeleteYear(c *gin.Context) {
	idParam := c.Param("id")
	id, _ := strconv.Atoi(idParam)
	models.DeleteYear(id)
	c.JSON(200, gin.H{
		"message": "Year Deleted",
	})
}