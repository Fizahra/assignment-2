package controllers

import (
	"net/http"

	"assignment2/structs"

	"github.com/gin-gonic/gin"
)

func (idb *InDB) GetOrder(c *gin.Context) {
	var (
		order  structs.Order
		result gin.H
	)

	order, err := idb.Orderepository.GetOrderByID()
	if err != nil {
		result = gin.H{
			"result": err.Error(),
			"count":  0,
		}
	} else {
		result = gin.H{
			"result": order,
			"count":  1,
		}
	}
	c.JSON(http.StatusOK, result)

}

func (idb *InDB) GetOrders(c *gin.Context) {
	var (
		orders []structs.Order
		result gin.H
	)

	idb.Orderepository.GetOrders()
	if len(orders) <= 0 {
		result = gin.H{
			"result": nil,
			"count":  0,
		}
	} else {
		result = gin.H{
			"result": orders,
			"count":  len(orders),
		}
	}
	c.JSON(http.StatusOK, result)
}

func (idb *InDB) CreateOrder(c *gin.Context) {
	var (
		order  structs.Order
		result gin.H
	)
	customername := c.PostForm("Customer_Name")
	items := c.PostForm("Items")
	order.Customer_Name = customername
	order.Items = items
	idb.Orderepository.Create()
	result = gin.H{
		"result": order,
	}

	c.JSON(http.StatusOK, result)

}

func (idb *InDB) UpdateOrder(c *gin.Context) {

	customername := c.PostForm("Customer_Name")
	items := c.PostForm("Items")
	var (
		order    structs.Order
		neworder structs.Order
		result   gin.H
	)
	order, err := idb.Orderepository.GetOrderByID()
	if err != nil {
		result = gin.H{
			"result": "data not found",
		}
	}

	neworder.Customer_Name = customername
	neworder.Items = items
	order, err = idb.Orderepository.Update()
	if err != nil {
		result = gin.H{
			"result": "update failed",
		}
	} else {
		result = gin.H{
			"result": "succesfully updated data",
		}
	}
	c.JSON(http.StatusOK, result)
}

func (idb *InDB) DeleteOrder(c *gin.Context) {
	var (
		order  structs.Order
		result gin.H
	)

	order, err := idb.Orderepository.GetOrderByID()
	if err != nil {
		result = gin.H{
			"result": "data not found",
		}
	}
	order, err = idb.Orderepository.Delete()
	if err != nil {
		result = gin.H{
			"result": "delete failed",
		}
	} else {
		result = gin.H{
			"result": "Data succesfully deleted",
		}
	}
	c.JSON(http.StatusOK, result)
}
