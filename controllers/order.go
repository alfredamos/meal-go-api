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
	context.JSON(http.StatusCreated, gin.H{"message": "Order created successfully"})

}

func DeleteOrderById(context *gin.Context){
	//----> Declare the type.
	order := models.Order{}

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
	order := models.Order{}

	//----> Get the id from param.
	userIdd := context.Param("userId")
	userId, err := strconv.ParseUint(userIdd, 10, 32)

	//----> Check for error.
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Please provide valid id!"})
		return
	}

	//----> Check for ownership permission
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
	context.JSON(http.StatusOK, gin.H{"message": "orders deleted by userId successfully!"})
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
	context.JSON(http.StatusOK, gin.H{"message": "orders deleted successfully!"})
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
	context.JSON(http.StatusOK, gin.H{"status": "success" , "message": "Orders are retrieved successfully!", "orders": orders})
}

func GetAllOrderByUserId(context *gin.Context){
	//----> declare the order variable.
	order := models.Order{}

	//----> Get the user-id from param.
	userIdd := context.Param("userId")
	userId, err := strconv.ParseUint(userIdd, 10, 32)
	
	//----> Check for error.
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Please provide valid userId!"})
		return
	}

	//----> Check for ownership permission
	err = authenticate.OwnerAuthorize(uint(userId), context)

	//----> Check for ownership.
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
	context.JSON(http.StatusOK, gin.H{"status": "success" , "message": "Orders are retrieved successfully!", "orders": orders})
}

func GetOrderById(context *gin.Context){
	//----> declare the order variable.
	order := models.Order{}

	//----> The id from params.
	idd:= context.Param("id")
	id, err := strconv.ParseUint(idd, 10, 32)

	//----> Check for error.
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Please provide a valid id!"})
		return
	}

	//----> Get all the orders by given user-id.
	order, err = order.GetOrderById(uint(id))

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
	order := models.Order{}

	//----> Get the user-id from param.
	idd := context.Param("id")
	id, err := strconv.ParseUint(idd, 10, 32)

	//----> Check for error.
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Please provide valid userId!"})
		return
	}

	//----> Change the order status.
	err = order.OrderDelivered(uint(id))

	//----> Check for error.
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": "Order is yet to be shipped or has already been deliver"})
	return
  }

 //----> Send back the response.
 context.JSON(http.StatusOK, gin.H{"status": "success", "message": "Order delivered successfully!"})

}
func OrderShipped(context *gin.Context){
	//----> declare the order variable.
	order := models.Order{}

 //----> Get the user-id from param.
 idd := context.Param("id")
 id, errUserId := strconv.ParseUint(idd, 10, 32)

 //----> Check for error.
 if errUserId != nil {
	context.JSON(http.StatusBadRequest, gin.H{"message": "Please provide valid userId!"})
	return
 }

 //----> Change the order status.
 err := order.OrderShipped(uint(id))

 //----> Check for error.
 if err != nil {
	context.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": "Order has already been shipped or deliver"})
	return
 }

 //----> Send back the response.
 context.JSON(http.StatusOK, gin.H{"status": "success", "message": "Order shipped successfully!"})

}