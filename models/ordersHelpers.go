package models

import (
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/alfredamos/go-meal-api/initializers"
)

func CalTotalPriceAndQuantity(carts Carts) (float64, float64) {
	totalQuantity := 0.0
	totalPrice := 0.0
	for _, value := range carts {
		totalQuantity += value.Quantity
		totalPrice += value.Quantity * value.Price
	}

	return totalQuantity, totalPrice
}

func makeOrder(carts []Cart, userId uint) Order {
	//----> Get the total quantity and total price.
	totalQuantity, totalPrice := CalTotalPriceAndQuantity(carts)

	//----> Make the order.
	order := Order{
		UserID:        userId,
		PaymentId:     "2edklugr",
		OrderDate:     time.Now(),
		TotalQuantity: totalQuantity,
		TotalPrice:    totalPrice,
		IsDelivered:   false,
		IsPending:     false,
		IsShipped:     false,
		Status:        "Pending",
	}

	return order
}

func makeCart(carts []Cart, orderId uint) []CartItem {
	newCarts := []CartItem{} //----> Cart variable.

	//----> Make the cart-items
	for _, value := range carts {
		newCart := CartItem{
			Name:     value.Name,
			Price:    value.Price,
			Quantity: value.Quantity,
			Image:    value.Image,
			OrderID:  orderId,
			PizzaID:  value.PizzaId,
		}
		//----> Append newCart to newCarts.
		newCarts = append(newCarts, newCart)
	}

	return newCarts
}

func deleteManyCartItems(carts []CartItem, id uint) error{
	//----> Get all the ids of the cart-items to be deleted.
	cartItems := getAllCartItemsIds(carts)
	fmt.Printf("%+v", cartItems)
	//----> Delete all cart-items.
	err := initializers.DB.Unscoped().Delete(&cartItems).Error
	//result := initializers.DB.Unscoped().Delete(&cartItems)

	//----> Check for error.
	if err != nil {
		return errors.New("cart-items cannot be deleted")
	}

	return nil
}

func deleteManyOrders(orders []Order) error{
	allOrders := make([]Order,0) //----> orders - ids.
	for _, order := range orders{
		oneOrder := Order{ID: order.ID}//----> Order-id.
		allOrders  = append(allOrders , oneOrder) //----> orders-ids.

		//----> Delete all cart-items associated with this order.
		err := deleteManyCartItems(order.CartItems, order.ID)

		if err != nil{
			return errors.New("cart-items cannot be deleted")
		}
	}

	//----> Delete all orders.
	result := initializers.DB.Unscoped().Delete(&allOrders)

	//----> Check for error.
	if result.RowsAffected == 0 {
		return errors.New("orders cannot be deleted")
	}

	return nil
}

func shippingInfo(order Order) error{
	//----> Update the order shipping info.
	order.IsShipped = true //----> Order shipped.
	order.IsPending = false //----> Order no longer pending.
	order.ShippingDate = sql.NullTime{Time: time.Now()} //----> Order shipping date.
	order.Status = "Shipped" //----> Order status.

	//----> Update order in the database.
	result := initializers.DB.Save(&order)

	//----> Check for error.
	if result.RowsAffected == 0 {
		return errors.New("order shipping info cannot be saved")
	}

	//----> Send back response.
	return nil
}

func deliveryInfo(order Order) error{
	//----> Update the order delivery info.
	order.IsDelivered = true; //----> Order shipped.
	order.DeliveryDate = sql.NullTime{Time: time.Now()} //----> Order shipping date.
	order.Status = "Delivered" //----> Order status.
	
	//----> Update order in the database.
	result := initializers.DB.Save(&order)

	//----> Check for error.
	if result.RowsAffected == 0 {
		return errors.New("order shipping info cannot be saved")
	}

	//----> Send back response.
	return nil
}