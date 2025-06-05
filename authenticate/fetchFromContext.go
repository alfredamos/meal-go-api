package authenticate

import (
	"github.com/gin-gonic/gin"
)

func GetRoleFromContext(c *gin.Context) (string, bool) {
	//----> Get user role from context.
	role := c.GetString("role")

	
		//----> Check for admin role.
		isAdmin := role == "Admin"


	//----> Send back the role.
	return role, isAdmin

}

func GetUserIdFromContext(c *gin.Context) string{
	//----> Get user-id from context.
	userId := c.GetString("userId")

	//----> Send back the user-id.
	return userId
}