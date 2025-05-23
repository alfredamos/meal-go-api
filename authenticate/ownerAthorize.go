package authenticate

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

func OwnerAuthorize(userId uint, c *gin.Context) error {
		//----> Get user id from context.
		userIdFromContext, exists  := c.Get("userId")

		//----> Check for existence of user id in context.
		if !exists {
			return errors.New("you are not permitted to access this page or perform this function")
		}
	
		//----> Convert userId to string
		userIdStr := fmt.Sprintf("%v", userIdFromContext)
		userIdInt, err := strconv.ParseUint(userIdStr, 10, 64)

		//----> Check for error.
		if err != nil {
			return errors.New("you are not permitted to access this page or perform this function")
		}

		fmt.Println("User argument, userId : ", userId)
		fmt.Println("User from context, userIdInt : ", userIdInt)
		
		//----> Convert to uint.
		userIdUint := uint(userIdInt)

		//----> Check for equality of userId.
		isSameUser := isSame(userIdUint, userId) 

		//----> Get the user role.
		role, err := getRoleFromContext(c)

		//----> Check for error.
		if err != nil{
			return errors.New("role cannot be retrieved from context")
		}

		//----> Check for passage criteria.
		isAdmin := role == "Admin"

		fmt.Println("isAdmin : ", isAdmin)
		fmt.Println("isSameUser : ", isSameUser)

		//----> Admin is allowed and same user is also allowed.
		if isAdmin || isSameUser {
			return nil
		}

		// You are not admin neither is same user, hence you are not allowed.
		return errors.New("you are not permitted to access this page or perform this function")

}


func isSame(numb1, numb2 uint) bool{
	return numb1 == numb2
}

