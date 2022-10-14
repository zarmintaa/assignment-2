package controllers

import (
	"assignment-2/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"strconv"
	"time"
)

type InDB struct {
	DB *gorm.DB
}

func (idb *InDB) CreateOrder(ctx *gin.Context) {
	var newOrder models.Orders
	err := ctx.ShouldBindJSON(&newOrder)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"result": nil,
			"err":    err,
		})
	}

	newOrder.Ordered_at = time.Now()
	err = idb.DB.Debug().Create(&newOrder).Error

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"result": nil,
			"err":    err,
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"result": newOrder,
	})
}

func (idb *InDB) GetOrder(ctx *gin.Context) {
	var orders []models.Orders

	err := idb.DB.Preload("Items").Find(&orders).Error

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"orders": nil,
			"err":    err,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"orders": orders,
		"count":  len(orders),
	})
}

func (idb InDB) GetOrderDetail(ctx *gin.Context) {
	id := ctx.Param("id")
	idOrder, _ := strconv.Atoi(id)

	var order models.Orders
	idb.DB.Debug().Preload("Items").First(&order, idOrder)

	if order.Order_id == 0 {
		ctx.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No order found!"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"result": order,
	})

}

func (idb *InDB) DeleteOrder(ctx *gin.Context) {
	id := ctx.Param("id")
	idOrder, _ := strconv.Atoi(id)

	var order models.Orders
	idb.DB.Debug().First(&order, idOrder)

	if order.Order_id == 0 {
		ctx.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No order found!"})
		return
	}

	errDelete := idb.DB.Where("order_id = ?", id).Delete(&order).Association("Items").Clear()

	if errDelete != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"result": nil,
			"err":    errDelete,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"result": "deleted",
	})
}

func (idb *InDB) UpdateOrder(ctx *gin.Context) {
	id := ctx.Param("id")
	idOrder, _ := strconv.Atoi(id)

	var order models.Orders
	idb.DB.Debug().First(&order, idOrder)

	if order.Order_id == 0 {
		ctx.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No order found!"})
		return
	}

	err := ctx.ShouldBindJSON(&order)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"result": nil,
			"err":    err,
		})
	}

	err = idb.DB.Debug().Model(&order).Association("Items").Replace(order.Items)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"result": nil,
			"err":    err,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"result": order,
	})
}
