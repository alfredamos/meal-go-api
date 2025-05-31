package controllers

import (
	"net/http"
	"strconv"
	"github.com/alfredamos/go-meal-api/authenticate"
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
		context.JSON(http.StatusBadRequest, gin.H{"message": "All values must be provided!"})
		return
	}

	//----> Save the order in the database.
	order.CheckOutOrder()

	//----> send back the response
	context.JSON(http.StatusCreated, gin.H{"status": "Success", "message": "Order has been created successfully!", "statusCode": http.StatusCreated})

}

func DeleteOrderById(context *gin.Context){
	//----> Declare the type.
	order := models.Order{}

	//----> Get order id from params.
	id, err := strconv.Atoi(context.Param("id"))

	//----> Check for error.
	if err != nil {
		context.JSON(http.StatusOK, gin.H{"message": "Order-id couldn't be parsed!"})
		return
	}
	
	//----> Delete order with this id.
	order.DeleteOrderById(uint(id))

	//----> Send back the response.
	context.JSON(http.StatusNoContent, gin.H{"status": "Success", "message": "Order has been deleted successfully!", "statusCode": http.StatusNoContent})
}

func DeleteOrderByUserId(context *gin.Context){
	//----> Declare the order type.
	order := models.Order{}

	//----> Get the user-id from param.
	userId, err := strconv.Atoi(context.Param("id"))

	//----> Check for error.
	if err != nil {
		context.JSON(http.StatusOK, gin.H{"message": "User-id couldn't be parsed!"})
		return
	}

	//----> Check for ownership permission or admin privilege.
	err = authenticate.OwnerAuthorize(uint(userId), context)

	//----> Check for ownership.
	if err != nil {
		context.JSON(http.StatusForbidden, gin.H{"status": "fail", "message": "You are not permitted to view or perform any action on this page!"})
		return
	}

	//----> Delete all orders attach to this userId.
	err = order.DeleteOrderByUserId(uint(userId))

	//----> Check for error
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": "orders are not available in the database!"})
		return
	}

	//----> Send back the response.
	context.JSON(http.StatusNoContent, gin.H{"status": "Success", "message": "Order has been deleted successfully!", "statusCode": http.StatusNoContent})
}

func DeleteAllOrders(context *gin.Context){
	//----> Declare the order type.
	order := models.Order{}

	//----> Delete all orders.
	err := order.DeleteAllOrders()

	//----> Check for error
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": "orders are not available in the database!"})
		return
	}

	//----> Send back the response.
	context.JSON(http.StatusNoContent, gin.H{"status": "Success", "message": "Order has been deleted successfully!", "statusCode": http.StatusNoContent})
}

func GetAllOrders(context *gin.Context){
	//----> declare the order type.
	order := models.Order{}

	//----> Get all orders from the database.
	orders, err := order.GetAllOrders()

	//----> Check for error.
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": "orders cannot be retrieved!"})
		return
	}

	//----> Send back the response.
	context.JSON(http.StatusOK, orders)
}

func GetAllOrderByUserId(context *gin.Context){
	//----> declare the order variable.
	order := models.Order{}

	//----> Get the user-id from param.
	userId, err := strconv.Atoi(context.Param("id"))

	//----> Check for parsing error.
	if err != nil {
		context.JSON(http.StatusForbidden, gin.H{"status": "fail", "message": "You are not permitted to view or perform any action on this page!"})
		return
	}

	//----> Check for ownership permission or admin privilege.
	err = authenticate.OwnerAuthorize(uint(userId), context)

	//----> Check for error.
	if err != nil {
		context.JSON(http.StatusForbidden, gin.H{"status": "fail", "message": "You are not permitted to view or perform any action on this page!"})
		return
	}

	//----> Get all the orders by given user-id.
	orders, err := order.GetAllOrdersByUserId(uint(userId))

	//----> Check for error.
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": "orders cannot be retrieved!"})
		return
	}

	//----> Send back the response.
	context.JSON(http.StatusOK, orders)
}

func GetOrderById(context *gin.Context){
	//----> declare the order variable.
	order := models.Order{}

	//----> Check for ownership permission or admin privilege.
	err := authenticate.OwnerAuthorize(order.UserID, context)

	//----> Check for ownership.
	if err != nil {
		context.JSON(http.StatusForbidden, gin.H{"status": "fail", "message": "You are not permitted to view or perform any action on this page!"})
		return
	}

	//----> The id from params.
	id, err := strconv.Atoi(context.Param("id"))

	//----> Check for error.
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Please provide a valid id!"})
		return
	}

	//----> Get order by order-id.
	order, err = order.GetOrderById(uint(id))

	//----> Check for error.
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": "order cannot be retrieved!"})
		return
	}

	//----> Send back the response.
	context.JSON(http.StatusOK, order)
}

func OrderDelivered(context *gin.Context){
	//----> declare the order variable.
	order := models.Order{}

	//----> Get the order-id from param.
	id, err := strconv.Atoi(context.Param("id"))

	//----> Check for error.
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Please provide valid userId!"})
		return
	}

	//----> Change the order status.
	orderEdited, err := order.OrderDelivered(uint(id))

	//----> Check for error.
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": "Order is yet to be shipped or has already been deliver"})
	return
  }

 //----> Send back the response.
 context.JSON(http.StatusOK, orderEdited)

}

func OrderShipped(context *gin.Context){
	//----> declare the order variable.
	order := models.Order{}

 //----> Get the order-id from param.
 id, err := strconv.Atoi(context.Param("id"))

 //----> Check for error.
 if err != nil {
	context.JSON(http.StatusBadRequest, gin.H{"message": "Please provide valid userId!"})
	return
 }

 //----> Change the order status.
 orderEdited, err := order.OrderShipped(uint(id))

 //----> Check for error.
 if err != nil {
	context.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": "Order has already been shipped or deliver"})
	return
 }

 //----> Send back the response.
 context.JSON(http.StatusOK, orderEdited)

}