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
		result := utils.Contains(roles, roleToString)

		//----> Check for valid role.
		if  !result {
			//----> Invalid role.
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"status": "fail","message": "You are not permitted to access this page or perform this function!", "statusCode": http.StatusForbidden})
			return
		}else if result{
			//----> Valid role.
			c.Next()
		}
	}
}


