package controllers

import (
	"fmt"
	"net/http"

	"github.com/alfredamos/go-meal-api/authenticate"
	"github.com/alfredamos/go-meal-api/models"
	"github.com/gin-gonic/gin"
)

func ChangePasswordController(context *gin.Context){
	//----> Get the type
	var changePasswordModel models.ChangePasswordModel

	//----> Get the request payload.
	err := context.ShouldBindJSON(&changePasswordModel)

	//----> Check for error
	if err != nil{
		context.JSON(http.StatusBadRequest, gin.H{"message": "Invalid credentials!"})
		return
	}

	//----> Save the new password in the database.
	err = changePasswordModel.ChangePassword()

	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid credentials!"})
		return
	}

	//----> Send back the response.
	context.JSON(http.StatusOK, gin.H{"message": "Password has been changed successfully!"})
}


func EditProfileController(context *gin.Context){
	//----> Get the type
	var editProfileModel models.EditProfileModel

	//----> Get the request payload
	err := context.ShouldBindJSON(&editProfileModel)

	//----> Check for error.
	if err != nil {
		fmt.Println("At point 1, error : ", err)
		context.JSON(http.StatusBadRequest, gin.H{"message": "Invalid credentials!"})
		return
	}

	//----> Save the changed profiles into the database.
	user, err := editProfileModel.EditProfile()
	fmt.Println("At point 2, error : ", err)
	//----> Check for error.
	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid credentials!"})
		return
	}

	//----> Send back the response
	context.JSON(http.StatusOK, gin.H{"message": "Profile has been changed successfully!", "user": user})
}


func LoginController(context *gin.Context) {
	//----> Get the login-model type
	var loginModel  models.LoginModel

	//----> Get the request payload
	err := context.ShouldBindJSON(&loginModel)
	
	fmt.Println("Error : ", err)
	//----> Check for error
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Login failed!"})
		return
	}

	//----> Login
	token, err := loginModel.Login()

	//----> Check for errors.
	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"status": "failed!", "message": "Invalid credentials", "statusCode": http.StatusUnauthorized})
		return
	}

	//----> Set cookie.
	authenticate.SetCookieHandler(context, token)
	
	//----> Send back response.
	context.JSON(http.StatusOK, gin.H{"status": "success", "message": "Login is successful!", "token": token})
}

func LogoutController(context *gin.Context){
	//----> Remove the cookie.
	authenticate.DeleteCookieHandler(context)

	//----> Send back the response.
	context.JSON(http.StatusOK, gin.H{"status": "success", "message": "Logout is successful!"})
}

func SignupController(context *gin.Context){
	//----> Get the type.
	var signupModel models.SignupModel

	//----> Get the request payload
	err := context.ShouldBindJSON(&signupModel)
	
	//----> Check for error
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Please provide all values!"})
		return
	}

	//----> Save the new user in the database.
	err = signupModel.Signup()

	//----> Check for error.
	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"status": "Unauthorized", "message": "Invalid credentials", "statusCode": http.StatusUnauthorized})
		return
	}

	//----> Send back the response.
	context.JSON(http.StatusCreated, gin.H{"message": "Signup is successful"})
}
