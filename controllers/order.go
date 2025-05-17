package controllers

import (
	"net/http"
	"strconv"

	"github.com/alfredamos/go-meal-api/models"
	"github.com/gin-gonic/gin"
)

func Home(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{"message": "This is home!!!"})
}

func CreateOrder(context *gin.Context){
	//----> Declare the type.
	var order models.OrderPayload
	
	//----> Get the request payload
	err := context.ShouldBindJSON(&order)
	//fmt.Println("Error : ", err)
	//----> Check for error.
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "All values must be provided!"})
		return
	}

	//----> Save the order in the database.
	order.CreateOrder()

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

	//----> delete all orders attach to this userId.
	order.DeleteOrderByUserId(uint(userId))

	//----> Send back the response.
	context.JSON(http.StatusOK, gin.H{"message": "orders deleted by userId successfully!"})
}

func EditOrderById(context *gin.Context){
	//----> Declare the order type.
	var order models.Order

	//----> Get the id from param.
	idd := context.Param("id")
	id, errId := strconv.ParseUint(idd, 10, 32)

	//----> Check for error.
	if errId != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Please provide a valid id!"})
		return
	}

	//----> delete the order associated with this id.
	order.EditOrderId(uint(id))

	//----> send back the response.
	context.JSON(http.StatusOK, gin.H{"message": "Order is edited successfully!"})
}

func GetAllOrders(context *gin.Context){
	//----> declare the order type.
	var order models.Order

	//----> Get all orders from the database.
	order.GetAllOrders()

	//----> Send back the response.
	context.JSON(http.StatusOK, gin.H{"message": "Orders are retrieved successfully!"})
}

func GetAllOrderByUserId(context *gin.Context){
	//----> declare the order variable.
	var order models.Order

	//----> Get the user-id from param.
	userIdd := context.Param("id")
	userId, errUserId := strconv.ParseUint(userIdd, 10, 32)
	
	//----> Check for error.
	if errUserId != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Please provide valid userId!"})
		return
	}

	//----> Get all the orders by given user-id.
	order.GetAllOrdersByUserId(uint(userId))

	//----> Send back the response.
	context.JSON(http.StatusOK, gin.H{"message": "Orders are retrieved successfully!"})
}

func GetOrderById(context *gin.Context){
//----> declare the order variable.
var order models.Order

//----> Get all the orders by given user-id.
order.GetAllOrders()

//----> Send back the response.
context.JSON(http.StatusOK, gin.H{"message": "Orders are retrieved successfully!"})
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