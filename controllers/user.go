package controllers

import (
	"net/http"
	"strconv"

	"github.com/alfredamos/go-meal-api/models"
	"github.com/gin-gonic/gin"
)

func DeleteUserById(context *gin.Context) {
	//----> Declare user type.
	var user models.User

	//----> Get the user id from param
	idd := context.Param("id")
	id, errId := strconv.ParseUint(idd, 10, 32)

	//----> Check the error.
	if errId != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Please provide a valid id!"})
		return
	}

	//----> Retrieve the pizza with the given id from database.
	err := user.DeleteUserById(uint(id))

	//----> Check if the user exist.
	if err != nil{
		context.JSON(http.StatusNotFound, gin.H{"message": "This user cannot be deleted!"})
	}

	//----> Send back the response
	context.JSON(http.StatusOK, gin.H{"message": "User is deleted successfully!"})

}

func GetAllUsers(context *gin.Context) {
	//----> Declare user type.
	var user models.User

	//----> Retrieve the users from the database.
	users, err := user.GetAllUsers()
	
	//----> Check for error.
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"status": "failed", "message": "There are no users in the database!", "statusCode": http.StatusNotFound})
	}

	//----> Send back the response.
	context.JSON(http.StatusOK, gin.H{"message": "Users are retrieve successfully!", "users": users})
}

func GetUserById(context *gin.Context) {
	//----> Declare user type.
	var user models.User

	//----> Get the user id from param.
	idd := context.Param("id")
	id, errId := strconv.ParseUint(idd, 10, 32)

	//----> Check for error
	if errId != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Please provide a valid id!"})
		return
	}

	//----> Get the user with the given id from database.
	user, err := user.GetUserById(uint(id))

	//----> Check if the user exist.
	if err != nil{
		context.JSON(http.StatusNotFound, gin.H{"message": "The user is not available in the database!"})
	}

	//----> Send back the response
	context.JSON(http.StatusOK, gin.H{"message": "The user is retrieved successfully!", "user": user})
}