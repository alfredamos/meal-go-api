package authenticate

import (
	"errors"
	"github.com/gin-gonic/gin"
)

func OwnerAuthorize(userId uint, c *gin.Context) error {
		//----> Get user id from context.
		userIdInt := GetUserIdFromContext(c)

		//----> Check for equality of userId.
		userIsSame := isSameUser(userIdInt, userId) 
		
		//----> Same user allowed.
		if userIsSame {

			return nil
		}

		//----> Get the user role.
		role := GetRoleFromContext(c)

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
func isSameUser(userId1, userId2 uint) bool{
	return userId1 == userId2
}

