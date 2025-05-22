package authenticate

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm/utils"
)

func RolePermission(roles []string) gin.HandlerFunc{
	return func(c *gin.Context){
		//----> Get user role from context.
		role, exists  := c.Get("role")
		
		//----> Check for existence of role.
		if !exists {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"status": "fail","message": "You are not permitted to access this page or perform this function!", "statusCode": http.StatusForbidden})
			return
		}

		//----> Convert role to string
		roleToString := fmt.Sprintf("%v", role)

		//----> Check for role in roles slice.
		isValidRole := utils.Contains(roles, roleToString)

		//----> Check for valid role.
		if  !isValidRole {
			//----> Invalid role.
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"status": "fail","message": "You are not permitted to access this page or perform this function!", "statusCode": http.StatusForbidden})
			return
		}
			
		//----> The role is valid, user is authorized.
		c.Next()
	
	}
}


