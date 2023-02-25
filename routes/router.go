package routes

import (
	"github.com/gin-gonic/gin"
	controllers "github.com/mrehanabbasi/appraisal-system-backend/controller"
)

func NewRouter() *gin.Engine {
	router := gin.Default()
	v1 := router.Group("/v1")
	// TODO: Define/Call routes here
	sc := controllers.NewSupervisorController()

	v1.POST("/supervisors", sc.ConvertSupervisorToEmployee)
	v1.GET("/supervisors", sc.GetSupervisors)
	v1.GET("/supervisors/:id", sc.GetSupervisorById)
	v1.PUT("/supervisors/:id", sc.UpdateSupervisor)
	v1.DELETE("/supervisors/:id", sc.DeleteSupervisor)

	return router
}
