package controllers

import (
	"fmt"
	"net/http"
	"github.com/alfredamos/go-meal-api/models"
	"github.com/gin-gonic/gin"
)

func CreateCartItem(context *gin.Context) {
	cartItem := models.CartItem{} //----> CartItem variable
	
	//----> Get the cart-item payload from the request.
	err := context.ShouldBindJSON(&cartItem)

	//----> Check for binding error.
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"status": "failed!", "message": fmt.Sprintf("%v", err)})
		return
	}

	//----> Insert the cart-item into the database.
	err = cartItem.CreateCartItem()

	//----> Check for error.
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"status": "failed!", "message": fmt.Sprintf("%v", err)})
		return
	}

	//----> send back the response.
	context.JSON(http.StatusBadRequest, gin.H{"status": "Success", "message": "Cart-item has been created successfully!"})
}

func DeleteCartItemById(context *gin.Context) {
	cartItem := models.CartItem{} //----> Cart-item variable.
	
	//----> Get the id from param.
	id := context.Param("id")
	
	//----> Delete the cart-item from the database.
	err := cartItem.DeleteCartItemById(id)

	//----> Check for error.
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"status": "failed!", "message": fmt.Sprintf("%v", err)})
		return
	}

	//----> Send back the response.
	context.JSON(http.StatusNoContent, gin.H{"status": "Success", "message": "Cart-item has been deleted successfully!"})
}

func EditCartItemById(context *gin.Context) {
	cartItem := models.CartItem{} //----> Cart-item variable.

	//----> Get the id from param.
	id := context.Param("id")

	//----> Get the request payload
	err := context.ShouldBindJSON(&cartItem)

	//----> Check for error.
	if err != nil {
	 context.JSON(http.StatusBadRequest, gin.H{"status": "failed!", "message": fmt.Sprintf("%v", err)})
	 return
	}

	//----> Update cart-item in the database.
	err = cartItem.EditCartItemId(id)

	//----> Check for error.
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"status": "failed!", "message": fmt.Sprintf("%v", err)})
		return
	}

	//----> send back the response.
	context.JSON(http.StatusNoContent, gin.H{"status": "Success", "message": "Cart-item has been edited successfully!"})
}

func GetAllCartItems(context *gin.Context) {
	cartItem := models.CartItem{} //----> Cart-item variable.

	//----> Retrieve all the cart-items from database.
	cartItems, err := cartItem.GetAllCartItems()

	//----> Check for error.
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"status": "failed!", "message": fmt.Sprintf("%v", err)})
		return
	}

	//----> send back response.
	context.JSON(http.StatusOK, cartItems)
}

func GetCartItemById(context *gin.Context) {
	cartItem := models.CartItem{} //----> Cart-item variable.
	
	//----> Get the id from param.
	id := context.Param("id")

	//----> Retrieve cart-item from database.
	cartItem, err := cartItem.GetCartItemById(id)

	//----> Check for error.
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"status": "failed!", "message": fmt.Sprintf("%v", err)})
		return
	}

	//----> send back the response.
	context.JSON(http.StatusOK, cartItem)
}