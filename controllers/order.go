package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/alfredamos/go-meal-api/models"
	"github.com/gin-gonic/gin"
)

func Home(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{"message": "This is home!!!"})
}

func CheckOutOrder(context *gin.Context){
	fmt.Println("I want to checkout order please!")
	//----> Declare the type.
	var order models.OrderPayload
	
	//----> Get the request payload
	err := context.ShouldBindJSON(&order)
	fmt.Println("Error : ", err)
	//----> Check for error.
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "All values must be provided!"})
		return
	}

	//----> Save the order in the database.
	order.CheckOutOrder()

	//----> send back the response
	context.JSON(http.StatusCreated, gin.H{"message": "Order created successfully"})

}

func DeleteOrderById(context *gin.Context){
	//----> Declare the type.
	var order models.Order

	//----> The id from params.
	idd:= context.Param("id")
	id, errId:= strconv.ParseUint(idd, 10, 32)

	//----> Check for error.
	if errId != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Please provide a valid id!"})
		return
	}

	//----> Delete order with this id.
	order.DeleteOrderById(uint(id))

	//----> Send back the response.
	context.JSON(http.StatusOK, gin.H{"message": "Order deleted successfully!"})
}

func DeleteOrderByUserId(context *gin.Context){
	//----> Declare the order type.
	var order models.Order

	//----> Get the id from param.
	userIdd := context.Param("userId")
	userId, errId := strconv.ParseUint(userIdd, 10, 32)

	//----> Check for error.
	if errId != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Please provide valid id!"})
		return
	}

	//----> Delete all orders attach to this userId.
	err := order.DeleteOrderByUserId(uint(userId))

	//----> Check for error
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": "orders are not available in the database!"})
		return
	}

	//----> Send back the response.
	context.JSON(http.StatusOK, gin.H{"message": "orders deleted by userId successfully!"})
}

func DeleteAllOrders(context *gin.Context){
	//----> Declare the order type.
	var order models.Order

	//----> Delete all orders.
	err := order.DeleteAllOrders()

	//----> Check for error
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": "orders are not available in the database!"})
		return
	}

	//----> Send back the response.
	context.JSON(http.StatusOK, gin.H{"message": "orders deleted successfully!"})
}

func GetAllOrders(context *gin.Context){
	//----> declare the order type.
	var order models.Order

	//----> Get all orders from the database.
	orders, err := order.GetAllOrders()

	//----> Check for error.
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": "orders cannot be retrieved!"})
		return
	}

	//----> Send back the response.
	context.JSON(http.StatusOK, gin.H{"status": "success" , "message": "Orders are retrieved successfully!", "orders": orders})
}

func GetAllOrderByUserId(context *gin.Context){
	//----> declare the order variable.
	var order models.Order

	//----> Get the user-id from param.
	userIdd := context.Param("userId")
	userId, errUserId := strconv.ParseUint(userIdd, 10, 32)
	
	//----> Check for error.
	if errUserId != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Please provide valid userId!"})
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
	context.JSON(http.StatusOK, gin.H{"status": "success" , "message": "Orders are retrieved successfully!", "orders": orders})
}

func GetOrderById(context *gin.Context){
	//----> declare the order variable.
	var order models.Order

	//----> The id from params.
	idd:= context.Param("id")
	id, errId:= strconv.ParseUint(idd, 10, 32)

	//----> Check for error.
	if errId != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Please provide a valid id!"})
		return
	}

	//----> Get all the orders by given user-id.
	order, err := order.GetOrderById(uint(id))

	//----> Check for error.
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": "order cannot be retrieved!"})
		return
	}

	//----> Send back the response.
	context.JSON(http.StatusOK, gin.H{"status": "success" , "message": "Orders are retrieved successfully!", "order": order})
}

func OrderDelivered(context *gin.Context){
//----> declare the order variable.
var order models.Order

//----> Get the user-id from param.
idd := context.Param("id")
id, errUserId := strconv.ParseUint(idd, 10, 32)

//----> Check for error.
if errUserId != nil {
	context.JSON(http.StatusBadRequest, gin.H{"message": "Please provide valid userId!"})
	return
}

//----> Change the order status.
order.OrderDelivered(uint(id))

}
func OrderShipped(context *gin.Context){
	//----> declare the order variable.
  var order models.Order

 //----> Get the user-id from param.
 idd := context.Param("id")
 id, errUserId := strconv.ParseUint(idd, 10, 32)

 //----> Check for error.
 if errUserId != nil {
	context.JSON(http.StatusBadRequest, gin.H{"message": "Please provide valid userId!"})
	return
 }

 //----> Change the order status.
 order.OrderShipped(uint(id))

}