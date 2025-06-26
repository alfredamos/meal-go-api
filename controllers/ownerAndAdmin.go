package controllers

import (
	"net/http"
	"github.com/alfredamos/go-meal-api/authenticate"
	"github.com/alfredamos/go-meal-api/models"
	"github.com/gin-gonic/gin"
)

func OwnerAndAdmin(c *gin.Context){
	//----> Get the order id from param.
	id := c.Param("id")

	//----> Get the order with the given id.
	orderVariable := models.Order{} //----> Order variable
	order, err := orderVariable.GetOrderById(id)

	//----> Check for error in fetching order.
	if err != nil {
		//----> Invalid user.
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"status": "fail","message": "Order is not found in the database!"})
		return
	}

	//----> Get the user-id from the order.
	userIdFromOrder := order.ID

	//----> Get admin user.
	_, userId, isAdmin := authenticate.GetUserAuthFromContext(c)

	//----> Check for equality of userId.
	userIsSame := authenticate.IsSameUser(userIdFromOrder, userId) 


	//----> Admin and owner user are allowed.
	if isAdmin || userIsSame {
			//----> Invalid user.
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"status": "fail","message": "You are not permitted to access this page!"})
			return
	}

	//----> Owner user and admin are allowed to access this page.
	c.Next()
}

