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

		isEqual := isSame(userIdUint, userId)
		
		//----> Check for existence of role.
		if !isEqual {
			return errors.New("you are not permitted to access this page or perform this function")
		}

		//----> User is the same you can continue.
		return nil

}


func isSame(numb1, numb2 uint) bool{
	return numb1 == numb2
}

