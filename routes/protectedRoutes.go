package routes

import (
	"github.com/alfredamos/go-meal-api/controllers"
	"github.com/gin-gonic/gin"
)

func protectedRoutes(r gin.IRoutes){
	//----> Auth routes.
	r.GET("/auth/current-user", controllers.GetCurrentUserController)
	r.PATCH("/auth/change-password", controllers.ChangePasswordController)
	r.PATCH("/auth/edit-profile", controllers.EditProfileController)
	r.POST("/auth/logout",controllers.LogoutController)
	
	//----> Cart-item routes.
	r.GET("/cart-items", controllers.GetAllCartItems)
	r.POST("/cart-items", controllers.CreateCartItem)
	r.DELETE("/cart-items/:id", controllers.DeleteCartItemById)
	r.GET("/cart-items/:id", controllers.GetCartItemById)
	r.PATCH("/cart-items/:id", controllers.EditCartItemById)

	//----> Order routes.
	r.PATCH("/orders/checkout", controllers.CheckOutOrder)
	r.GET("/orders/orders-by-user-id/:userId", controllers.GetAllOrderByUserId)
	r.DELETE("/orders/delete-all-orders-by-user-id/:userId", controllers.DeleteOrderByUserId)
	
	r.GET("/orders/:id", controllers.GetOrderById)
	r.DELETE("/orders/:id", controllers.DeleteOrderById)
	
	//----> Pizza-routes.
	r.GET("/pizzas/:id", controllers.GetPizzaById)

	//----> User-route
	r.GET("/users/:id", controllers.GetUserById)

}