package authenticate

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm/utils"
)

func RolePermission(roles ...string) gin.HandlerFunc{
	return func(c *gin.Context){
		//----> Get user role from context.
		role := GetRoleFromContext(c)

		//----> Check for role in roles slice.
		isValidRole := utils.Contains(roles, role)

		//----> Check for invalid role.
		if  !isValidRole {
			//----> Invalid role.
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"status": "fail","message": "You are not permitted to access this page or perform this function!", "statusCode": http.StatusForbidden})
			return
		}
			
		//----> The role is valid, user is authorized.
		c.Next()
	
	}
}




