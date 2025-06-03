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
	changePasswordModel := models.ChangePasswordModel{}

	//----> Get the request payload.
	err := context.ShouldBindJSON(&changePasswordModel)

	//----> Check for error.
	if err != nil{
		context.JSON(http.StatusBadRequest, gin.H{"status": "failed!", "message": fmt.Sprintf("%v", err)})
		return
	}

	//----> Save the new password in the database.
	err = changePasswordModel.ChangePassword()
	
	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"status": "failed!", "message": fmt.Sprintf("%v", err)})
		return
	}

	//----> Send back the response.
	context.JSON(http.StatusOK, gin.H{"status": "success","message": "Password has been changed successfully!"})
}

func GetCurrentUserController(context *gin.Context){
	//----> Get the user-id of the current login user.
	userId := authenticate.GetUserIdFromContext(context)

	//----> Get the current login user.
	user, err := models.GetCurrentUser(userId)

	//----> Check for error.
	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"status": "failed", "message": err.Error()})
		return
	}
	
	//----> Send back the response.
	context.JSON(http.StatusOK, user)
}


func EditProfileController(context *gin.Context){
	//----> Get the type
	editProfileModel := models.EditProfileModel{}

	//----> Get the request payload
	err := context.ShouldBindJSON(&editProfileModel)
	
	//----> Check for error.
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"status": "failed!", "message": fmt.Sprintf("%v", err)})
		return
	}

	//----> Save the changed profiles into the database.
	err = editProfileModel.EditProfile()
	
	//----> Check for error.
	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"status": "failed!", "message": fmt.Sprintf("%v", err)})
		return
	}

	//----> Send back the response
	context.JSON(http.StatusOK, gin.H{"status": "success","message": "Profile has been changed successfully!"})
}


func LoginController(context *gin.Context) {
	//----> Get the login-model type
	loginModel := models.LoginModel{}

	//----> Get the request payload
	err := context.ShouldBindJSON(&loginModel)
	
	//----> Check for error
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"status": "failed!", "message": fmt.Sprintf("%v", err)})
		return
	}

	//----> Login in the user.
	loginResp, err := loginModel.Login()

	//----> Check for errors.
	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"status": "failed!", "message": fmt.Sprintf("%v", err)})
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
	context.JSON(http.StatusOK, gin.H{"status": "success","message": "Logout is successfully!"})
}

func SignupController(context *gin.Context){
	//----> Get the type.
	signupModel := models.SignupModel{}

	//----> Get the request payload
	err := context.ShouldBindJSON(&signupModel)
	
	//----> Check for error.
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"status": "failed!", "message": fmt.Sprintf("%v", err)})
		return
	}

	//----> Save the new user in the database.
	err = signupModel.Signup()

	//----> Check for error.
	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"status": "failed", "message": fmt.Sprintf("%v", err), "statusCode": http.StatusUnauthorized})
		return
	}

	//----> Send back the response.
	context.JSON(http.StatusCreated, gin.H{"status": "success","message": "Signup is successfully!"})
}
