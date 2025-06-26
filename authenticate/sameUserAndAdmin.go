package authenticate

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func SameUserAndAdmin(c *gin.Context){
	//----> Get the user-id from param.
	userIdFromContext := c.Param("id")
	//----> Get user role from context.
	_, userId, isAdmin := GetUserAuthFromContext(c)

	//----> Check for same user.
	isUserSame := IsSameUser(userId, userIdFromContext)

	//----> Check for same user and admin privilege.
	if !isUserSame && isAdmin {
		//----> Invalid user.
		c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"status": "fail","message": "You are not permitted to access this page!", "statusCode": http.StatusForbidden})
		return
	}

	//----> Same user and admin are allowed to gain access.
	c.Next()
}

func IsSameUser(userId1, userId2 string)bool{
	return userId1 == userId2
}