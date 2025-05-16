package routes

import (
	"github.com/alfredamos/go-meal-api/authenticate"
	"github.com/alfredamos/go-meal-api/controllers"
	"github.com/gin-gonic/gin"
)


func RegisteredRoutes(server *gin.Engine){
	//----> Role related permitted argument.
	list := make([]string, 0)

	//----> Unprotected routes.
	//----> Auth-routes.
	server.POST("/auth/signup", controllers.SignupController)
	server.POST("/auth/login",controllers.LoginController)
	
	//----> Pizza-routes.
	server.GET("/pizzas", controllers.GetAllPizza)

	//----> Apply middleware for protected routes
	r := server.Use(authenticate.VerifyToken)

	//----> Protected routes.
	//----> Auth routes.
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
	r.GET("/orders/orders-by-user-id/:userId", controllers.GetAllOrderByUserId)
	r.DELETE("/orders/delete-all-orders-by-user-id/:userId", controllers.DeleteOrderByUserId)
	r.PATCH("/orders/:id", controllers.EditOrderById)
	r.GET("/orders/:id", controllers.GetOrderById)
	r.DELETE("/orders/:id", controllers.DeleteOrderById)
	
	//----> Pizza-routes.
	r.GET("/pizzas/:id", controllers.GetPizzaById)

	//----> Admin role permitted routes middleware.
	list = append(list, "Admin")
	p := server.Use(authenticate.VerifyToken, authenticate.RolePermission(list))

	//----> Orders routes.
	p.GET("/orders", controllers.GetAllOrders)
	p.POST("/orders", controllers.CreateOrder)
	p.PATCH("/orders/:id/delivered", controllers.OrderDelivered)
	p.PATCH("/orders/:id/shipped", controllers.OrderShipped)

	//----> Pizza routes.
	p.POST("/pizzas", controllers.CreatePizza)
	p.DELETE("/pizzas/:id", controllers.DeletePizzaById)
	p.PATCH("/pizzas/:id", controllers.EditPizzaById)
	//----> User routes.
	p.GET("/users", controllers.GetAllUsers)
	p.GET("/users/:id", controllers.GetUserById)
	p.DELETE("/users/:id", controllers.DeleteUserById)
}