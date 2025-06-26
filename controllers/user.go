package controllers

import (
	"fmt"
	"net/http"
	"github.com/alfredamos/go-meal-api/models"
	"github.com/gin-gonic/gin"
)

func DeleteUserById(context *gin.Context) {
	//----> Declare user type.
	user := models.User{}

	//----> Get the user id from param
	id:= context.Param("id")
	
	//----> Retrieve the user with the given id from database.
	err := user.DeleteUserById(id)

	//----> Check if the user exist.
	if err != nil{
		context.JSON(http.StatusNotFound, gin.H{"status": "failed!", "message": fmt.Sprintf("%v", err)})
		return
	}

	//----> Send back the response
	context.JSON(http.StatusNoContent, gin.H{"status": "success", "message": "User has been deleted successfully!"})

}

func GetAllUsers(context *gin.Context) {
	//----> Declare user type.
	user := models.User{}

	//----> Retrieve the users from the database.
	users, err := user.GetAllUsers()
	
	//----> Check for error.
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"status": "failed!", "message": fmt.Sprintf("%v", err)})
		return
	}

	//----> Send back the response.
	context.JSON(http.StatusOK, users)
}

func GetUserById(context *gin.Context) {
	//----> Declare user type.
	user := models.User{}

	//----> Get the user id from param.
	id := context.Param("id")

	//----> Get the user with the given id from database.
	user, err := user.GetUserById(id)

	//----> Check if the user exist.
	if err != nil{
		context.JSON(http.StatusNotFound, gin.H{"status": "failed!", "message": fmt.Sprintf("%v", err)})
		return
	}

	//----> Send back the response
	context.JSON(http.StatusOK, user)
}