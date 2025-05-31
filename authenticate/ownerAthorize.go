package authenticate

import (
	"errors"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm/utils"
)

func OwnerAuthorize(userId uint, c *gin.Context) error {
		//----> Get user id from context.
		userIdFromContext, exists  := c.Get("userId")

		//----> Check for existence of user id in context.
		if !exists {
			return errors.New("you are not permitted to access this page or perform this function")
		}

		//----> Convert user-id from context to string and then to int.
		userIdInt, err := strconv.Atoi(utils.ToString(userIdFromContext))
		//----> Check for parsing error.
		if err != nil {
			return errors.New("user-id could not be parsed")
		}
	
		//----> Check for equality of userId.
		isSameUser := isSame(uint(userIdInt), userId) 

		//----> Same user allowed.
		if isSameUser {
			return nil
		}

		//----> Get the user role.
		role, err := getRoleFromContext(c)

		//----> Check for error.
		if err != nil{
			return errors.New("role cannot be retrieved from context")
		}

		//----> Check for admin role.
		isAdmin := role == "Admin"

		//----> Admin is allowed.
		if isAdmin {
			return nil
		}

		// You are not admin neither is same user, hence you are not allowed.
		return errors.New("you are not permitted to access this page or perform this function")

}

//----> Check for checking for same user.
func isSame(userIdFromContext, userIdFromParam uint) bool{
	return userIdFromContext == userIdFromParam
}

