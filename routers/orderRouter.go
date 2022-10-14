package routers

import (
	"assignment-2/controllers"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func StartServer(db *gorm.DB) *gin.Engine {
	router := gin.Default()

	var controller = controllers.InDB{
		DB: db,
	}

	router.POST("/orders/create", controller.CreateOrder)
	router.GET("/orders", controller.GetOrder)
	router.GET("/orders/:id", controller.GetOrderDetail)
	router.DELETE("/orders/:id", controller.DeleteOrder)
	router.PUT("/orders/:id", controller.UpdateOrder)
	return router
}
