package controllers

import (
	"net/http"
	"strconv"
	"github.com/alfredamos/go-meal-api/authenticate"
	"github.com/alfredamos/go-meal-api/models"
	"github.com/gin-gonic/gin"
)

func DeleteUserById(context *gin.Context) {
	//----> Declare user type.
	user := models.User{}

	//----> Get the user id from param
	id, err := strconv.Atoi(context.Param("id"))

 //----> Check for error
 if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Please provide a valid id!"})
		return
 }
	
	//----> Retrieve the user with the given id from database.
	err = user.DeleteUserById(uint(id))

	//----> Check if the user exist.
	if err != nil{
		context.JSON(http.StatusNotFound, gin.H{"message": "This user cannot be deleted!"})
		return
	}

	//----> Send back the response
	context.JSON(http.StatusNoContent, gin.H{"status": "success", "message": "User has been deleted successfully!", "statusCode": http.StatusNoContent})

}

func GetAllUsers(context *gin.Context) {
	//----> Declare user type.
	user := models.User{}

	//----> Retrieve the users from the database.
	users, err := user.GetAllUsers()
	
	//----> Check for error.
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"status": "failed", "message": "There are no users in the database!", "statusCode": http.StatusNotFound})
		return
	}

	//----> Send back the response.
	context.JSON(http.StatusOK, users)
}

func GetUserById(context *gin.Context) {
	//----> Declare user type.
	user := models.User{}

	//----> Get the user id from param.
	id, err := strconv.Atoi(context.Param("id"))

 //----> Check for error
 if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Please provide a valid id!"})
		return
 }
	
	//----> Check for ownership permission
	err = authenticate.OwnerAuthorize(uint(id), context)

	//----> Check for ownership.
	if err != nil {
		context.JSON(http.StatusForbidden, gin.H{"status": "fail", "message": "You are not permitted to view or perform any action on this page!"})
		return
	}

	//----> Get the user with the given id from database.
	user, err = user.GetUserById(uint(id))

	//----> Check if the user exist.
	if err != nil{
		context.JSON(http.StatusNotFound, gin.H{"message": "The user is not available in the database!"})
		return
	}

	//----> Send back the response
	context.JSON(http.StatusOK, user)
}