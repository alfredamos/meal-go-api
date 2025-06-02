package authenticate

import "github.com/gin-gonic/gin"

func GetRoleFromContext(c *gin.Context) string {
	//----> Get user role from context.
	role := c.GetString("role")

	//----> Send back the role.
	return role

}

func GetUserIdFromContext(c *gin.Context) uint{
	//----> Get user-id from context.
	userId := c.GetUint("userId")

	//----> Send back the user-id.
	return userId
}