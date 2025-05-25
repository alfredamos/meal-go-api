package authenticate

import (
	"errors"
	"fmt"
	"net/http"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm/utils"
)

func RolePermission(roles ...string) gin.HandlerFunc{
	return func(c *gin.Context){
		//----> Get user role from context.
		role, err := getRoleFromContext(c)

		//----> Check for error.
		if err != nil{
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"status": "fail","message": "You are not permitted to access this page or perform this function!", "statusCode": http.StatusForbidden})
			return
		}

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

func getRoleFromContext(c *gin.Context) (string, error){
	//----> Get user role from context.
	role, exists  := c.Get("role")
		
	//----> Check for existence of role.
	if !exists {
		return string(""), errors.New("you are not permitted to access this page or perform this function")
	}

	//----> Convert role to string
	roleToString := fmt.Sprintf("%v", role)

	//----> Send back the role.
	return roleToString, nil

}


