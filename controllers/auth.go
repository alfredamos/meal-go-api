package controllers

import (
	"net/http"
	"github.com/alfredamos/go-meal-api/authenticate"
	"github.com/alfredamos/go-meal-api/models"
	"github.com/gin-gonic/gin"
)

func ChangePasswordController(context *gin.Context){
	//----> Get the type
	changePasswordModel := models.ChangePasswordModel{}

	//----> Get the request payload.
	err := context.ShouldBindJSON(&changePasswordModel)

	//----> Check for error.
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
	context.JSON(http.StatusOK, gin.H{"status": "success","message": "Password has been changed successfully!", "statusCode": http.StatusOK})
}


func EditProfileController(context *gin.Context){
	//----> Get the type
	editProfileModel := models.EditProfileModel{}

	//----> Get the request payload
	err := context.ShouldBindJSON(&editProfileModel)

	//----> Check for error.
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Invalid credentials!"})
		return
	}

	//----> Save the changed profiles into the database.
	err = editProfileModel.EditProfile()
	
	//----> Check for error.
	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid credentials!"})
		return
	}

	//----> Send back the response
	context.JSON(http.StatusOK, gin.H{"status": "success","message": "Profile has been changed successfully!", "statusCode": http.StatusOK})
}


func LoginController(context *gin.Context) {
	//----> Get the login-model type
	loginModel := models.LoginModel{}

	//----> Get the request payload
	err := context.ShouldBindJSON(&loginModel)
	
	//----> Check for error
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Login failed!"})
		return
	}

	//----> Login in the user.
	loginResp, err := loginModel.Login()

	//----> Check for errors.
	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"status": "failed!", "message": "Invalid credentials", "statusCode": http.StatusUnauthorized})
		return
	}

	//----> Set cookie.
	authenticate.SetCookieHandler(context, loginResp.Token)
	
	//----> Send back response.
	context.JSON(http.StatusOK, loginResp)
}

func LogoutController(context *gin.Context){
	//----> Remove the cookie.
	authenticate.DeleteCookieHandler(context)

	//----> Send back the response.
	context.JSON(http.StatusOK, gin.H{"status": "success","message": "Logout is successfully!", "statusCode": http.StatusOK})
}

func SignupController(context *gin.Context){
	//----> Get the type.
	signupModel := models.SignupModel{}

	//----> Get the request payload
	err := context.ShouldBindJSON(&signupModel)
	
	//----> Check for error.
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
	context.JSON(http.StatusCreated, gin.H{"status": "success","message": "Signup is successfully!", "statusCode": http.StatusOK})
}
