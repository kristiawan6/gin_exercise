package routes

import (
	controllers "gingonic/src/controllers"

	"github.com/gin-gonic/gin"
)

func Router() {
	router := gin.Default()

	v1 := router.Group("/api/v1")
	{
		month := v1.Group("/month")
		{
			month.GET("/list", controllers.GetAllMonth)
			month.GET("/:id", controllers.GetMonthById)
			month.POST("/create", controllers.CreateMonth)
			month.PUT("/update/:id", controllers.UpdateMonth)
			month.DELETE("/delete/:id", controllers.DeleteMonth)
		}
		day := v1.Group("/day")
		{
			day.GET("/list", controllers.GetAllDay)
			day.GET("/:id", controllers.GetDayById)
			day.POST("/create", controllers.CreateDay)
			day.PUT("/update/:id", controllers.UpdateDay)
			day.DELETE("/delete/:id", controllers.DeleteDay)
		}
		year := v1.Group("/year")
		{
			year.GET("/list", controllers.GetAllYear)
			year.GET("/:id", controllers.GetYearById)
			year.POST("/create", controllers.CreateYear)
			year.PUT("/update/:id", controllers.UpdateYear)
			year.DELETE("/delete/:id", controllers.DeleteYear)
		}
	}
	router.Run(":8080")
}
