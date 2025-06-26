package controllers

import (
	"fmt"
	"net/http"
	"github.com/alfredamos/go-meal-api/models"
	"github.com/gin-gonic/gin"
)

func CheckOutOrder(context *gin.Context){
	//----> Declare the type.
	order := models.OrderPayload{}
	
	//----> Get the request payload
	err := context.ShouldBindJSON(&order)
	
	//----> Check for error.
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"status": "failed!", "message": fmt.Sprintf("%v", err)})
		return
	}

	//----> Save the order in the database.
	err = order.CheckOutOrder()

	//----> Check for error.
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"status": "failed!", "message": fmt.Sprintf("%v", err)})
		return
	}

	//----> send back the response
	context.JSON(http.StatusCreated, gin.H{"status": "success!", "message": "Order has been created successfully!"})

}

func DeleteOrderById(context *gin.Context){
	//----> Declare the type.
	order := models.Order{}

	//----> Get order id from params.
	id := context.Param("id")
	
	//----> Delete order with this id.
	err := order.DeleteOrderById(id)

	//----> Check for error.
	if err != nil {
		context.JSON(http.StatusOK, gin.H{"status": "failed!", "message": fmt.Sprintf("%v", err)})
		return
	}

	//----> Send back the response.
	context.JSON(http.StatusNoContent, gin.H{"status": "Success", "message": "Order has been deleted successfully!"})
}

func DeleteOrderByUserId(context *gin.Context){
	//----> Declare the order type.
	order := models.Order{}

	//----> Get the user-id from param.
	userId := context.Param("userId")

	//----> Delete all orders attach to this userId.
	err := order.DeleteOrderByUserId(userId)

	//----> Check for error
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"status": "failed!", "message": fmt.Sprintf("%v", err)})
		return
	}

	//----> Send back the response.
	context.JSON(http.StatusNoContent, gin.H{"status": "Success", "message": "Order has been deleted successfully!"})
}

func DeleteAllOrders(context *gin.Context){
	//----> Declare the order type.
	order := models.Order{}

	//----> Delete all orders.
	err := order.DeleteAllOrders()

	//----> Check for error
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"status": "failed!", "message": fmt.Sprintf("%v", err)})
		return
	}

	//----> Send back the response.
	context.JSON(http.StatusNoContent, gin.H{"status": "Success", "message": "All Orders have been deleted successfully!"})
}

func GetAllOrders(context *gin.Context){
	//----> declare the order type.
	order := models.Order{}

	//----> Get all orders from the database.
	orders, err := order.GetAllOrders()

	//----> Check for error.
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"status": "failed!", "message": fmt.Sprintf("%v", err)})
		return
	}

	//----> Send back the response.
	context.JSON(http.StatusOK, orders)
}

func GetAllOrderByUserId(context *gin.Context){
	//----> declare the order variable.
	order := models.Order{}

	//----> Get the user-id from param.
	userId := context.Param("userId")

	//----> Get all the orders by given user-id.
	orders, err := order.GetAllOrdersByUserId(userId)

	//----> Check for error.
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"status": "failed!", "message": fmt.Sprintf("%v", err)})
		return
	}

	//----> Send back the response.
	context.JSON(http.StatusOK, orders)
}

func GetOrderById(context *gin.Context){
	//----> declare the order variable.
	order := models.Order{}

	//----> The id from params.
	id := context.Param("id")

	//----> Get order by order-id.
	order, err := order.GetOrderById(id)

	//----> Check for error.
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"status": "failed!", "message": fmt.Sprintf("%v", err)})
		return
	}

	//----> Send back the response.
	context.JSON(http.StatusOK, order)
}

func OrderDelivered(context *gin.Context){
	//----> declare the order variable.
	order := models.Order{}

	//----> Get the order-id from param.
	id := context.Param("id")

	//----> Change the order status.
	orderEdited, err := order.OrderDelivered(id)

	//----> Check for error.
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"status": "failed!", "message": fmt.Sprintf("%v", err)})
	return
  }

 //----> Send back the response.
 context.JSON(http.StatusOK, orderEdited)

}

func OrderShipped(context *gin.Context){
	//----> declare the order variable.
	order := models.Order{}

 //----> Get the order-id from param.
 id := context.Param("id")

 //----> Change the order status.
 orderEdited, err := order.OrderShipped(id)

 //----> Check for error.
 if err != nil {
	context.JSON(http.StatusBadRequest, gin.H{"status": "failed!", "message": fmt.Sprintf("%v", err)})
	return
 }

 //----> Send back the response.
 context.JSON(http.StatusOK, orderEdited)

}