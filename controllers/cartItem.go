package controllers

import (
	"net/http"
	"strconv"
	"github.com/alfredamos/go-meal-api/models"
	"github.com/gin-gonic/gin"
)

func CreateCartItem(context *gin.Context) {
	cartItem := models.CartItem{} //----> CartItem variable
	
	//----> Get the cart-item payload from the request.
	err := context.ShouldBindJSON(&cartItem)

	//----> Check for binding error.
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"status": "failed", "message": "All values must be provided!", "statusCode": http.StatusBadRequest})
		return
	}

	//----> Insert the cart-item into the database.
	err = cartItem.CreateCartItem()

	//----> Check for error.
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"status": "failed", "message": "All values must be provided", "statusCode": http.StatusBadRequest})
		return
	}

	//----> send back the response.
	context.JSON(http.StatusBadRequest, gin.H{"status": "Success", "message": "Cart-item has been created successfully!", "statusCode": http.StatusCreated})
}

func DeleteCartItemById(context *gin.Context) {
	cartItem := models.CartItem{} //----> Cart-item variable.
	
	//----> Get the id from param.
	id, err := strconv.Atoi(context.Param("id"))
	
	//----> Check for error.
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": "id cannot be parsed!", "statusCode": http.StatusBadRequest})
		return
	} 

	//----> Delete the cart-item from the database.
	err = cartItem.DeleteCartItemById(uint(id))

	//----> Check for error.
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": "Cart-item with the given id cannot be deleted!", "statusCode": http.StatusNotFound})
		return
	}

	//----> Send back the response.
	context.JSON(http.StatusNoContent, gin.H{"status": "Success", "message": "Cart-item has been deleted successfully!", "statusCode": http.StatusNoContent})
}

func EditCartItemById(context *gin.Context) {
	cartItem := models.CartItem{} //----> Cart-item variable.

	//----> Get the id from param.
	id, err := strconv.Atoi(context.Param("id"))

	//----> Check for error.
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": "id cannot be parsed!", "statusCode": http.StatusBadRequest})
		return
	}

	//----> Get the request payload
	err = context.ShouldBindJSON(&cartItem)

	//----> Check for error.
	if err != nil {
	 context.JSON(http.StatusBadRequest, gin.H{"message": "All values must be provided!"})
	 return
	}

	//----> Update cart-item in the database.
	err = cartItem.EditCartItemId(uint(id))

	//----> Check for error.
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": "Cart-item cannot be edited", "statusCode": http.StatusNotFound})
		return
	}

	//----> send back the response.
	context.JSON(http.StatusNoContent, gin.H{"status": "Success", "message": "Cart-item has been edited successfully!", "statusCode": http.StatusNoContent})
}

func GetAllCartItems(context *gin.Context) {
	cartItem := models.CartItem{} //----> Cart-item variable.

	//----> Retrieve all the cart-items from database.
	cartItems, err := cartItem.GetAllCartItems()

	//----> Check for error.
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": "Cart-items cannot be edited", "statusCode": http.StatusNotFound})
		return
	}

	//----> send back response.
	context.JSON(http.StatusOK, cartItems)
}

func GetCartItemById(context *gin.Context) {
	cartItem := models.CartItem{} //----> Cart-item variable.
	
	//----> Get the id from param.
	id, err := strconv.Atoi(context.Param("id"))

	//----> Check for error.
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": "id cannot be parsed!", "statusCode": http.StatusBadRequest})
		return
	} 

	//----> Retrieve cart-item from database.
	cartItem, err = cartItem.GetCartItemById(uint(id))

	//----> Check for error.
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": "cart-item is not found!", "statusCode": http.StatusNotFound})
		return
	}

	//----> send back the response.
	context.JSON(http.StatusOK, cartItem)
}