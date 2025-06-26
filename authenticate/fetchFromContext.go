package authenticate

import (
	"github.com/gin-gonic/gin"
)

func GetUserAuthFromContext(c *gin.Context) (string, string, bool) {
	//----> Get user role from context.
	role := c.GetString("role")

	//----> Get the user-id from context.
	userId := c.GetString("userId")
	
	//----> Check for admin role.
	isAdmin := role == "Admin"

	//----> Send back the role.
	return role, userId, isAdmin

}

func GetUserIdFromContext(c *gin.Context) string{
	//----> Get user-id from context.
	userId := c.GetString("userId")

	//----> Send back the user-id.
	return userId
}