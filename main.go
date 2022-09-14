package main

import (
	"assignment2/config"
	"assignment2/controllers"
	"assignment2/repository"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db := config.DBInit()
	repoOrder := repository.NewRepository(db)
	inDB := &controllers.InDB{
		Orderepository: repoOrder,
	}

	router := gin.Default()

	router.GET("/order/:id", inDB.GetOrder)
	router.GET("/orders", inDB.GetOrders)
	router.POST("/orders", inDB.CreateOrder)
	router.PUT("/orders", inDB.UpdateOrder)
	router.DELETE("/orders/:id", inDB.DeleteOrder)
	router.Run(":8080")
}
