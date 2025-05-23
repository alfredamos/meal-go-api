package models

import (
	"database/sql"
	"errors"
	"time"
	"github.com/alfredamos/go-meal-api/initializers"
	"gorm.io/gorm"
)

type Status string

const (
	Delivered Status = "Delivered"
	Pending Status = "Pending"
	Shipped Status = "Shipped" 
)

type Order struct {
	ID        uint           `gorm:"primaryKey"`
  CreatedAt time.Time
  UpdatedAt time.Time
  DeletedAt gorm.DeletedAt `gorm:"index"`
	PaymentId string `json:"paymentId"`
	OrderDate time.Time `json:"orderDate"`
	ShippingDate sql.NullTime `gorm:"type:TIMESTAMP NULL"`
	DeliveryDate sql.NullTime `gorm:"type:TIMESTAMP NULL"`
	TotalQuantity float64 `json:"totalQuantity"`
	TotalPrice float64 `json:"totalPrice"`
	IsShipped bool `json:"isShipped"`
	IsPending bool `json:"isPending"`
	IsDelivered bool `json:"isDelivered"`
	Status Status `json:"status" binding:"required"`
	UserID uint `json:"userId" binding:"required"`
	User User 
	CartItems []CartItem `gorm:"foreignKey:OrderID"`

}

func (order *Order) DeleteOrderById(id uint) error{
	//----> Check to see if the order to be deleted is available in the database.
	 err := initializers.DB.Model(&Order{}).Preload("CartItems").First(&order, id).Error

	//----> Check for error.
	if err != nil {
		return errors.New("order does not exist")
	}

	//----> Get the cart-items from order.
	carts := order.CartItems 

	//----> Delete all the cart-items attached to order with the given id.
	 err = deleteManyCartItems(carts, id)

	//----> Check for error.
	if err != nil {
		return errors.New("cart-items cannot be deleted")
	} 
 
	//----> Delete the order with given id.
	err = initializers.DB.Unscoped().Delete(&Order{}, id).Error

	//----> Check for error.
	if err != nil {
		return errors.New("order cannot be deleted")
	}

	return nil
}

func (*Order) DeleteOrderByUserId(userId uint) error{
	orders := []Order{} //----> Orders variable.

	//----> Retrieve orders from database.
	err := initializers.DB.Preload("CartItems").Find(&orders, Order{UserID: userId}).Error
	//----> Check for error.
	if err != nil {
		return errors.New("orders are not available in the database")
	}
	
	//----> Delete all orders and associated cart-items connected to this user-id.
	err = deleteManyOrders(orders)

	//----> Check for error.
	if err != nil {
		return errors.New("orders cannot be deleted")
	}

	return nil
}

func (*Order) DeleteAllOrders() error{
	orders := []Order{} //----> Orders variable.

	//----> Retrieve orders from database.
	err := initializers.DB.Preload("CartItems").Find(&orders).Error
	//----> Check for error.
	if err != nil {
		return errors.New("orders are not available in the database")
	}
	
	//----> Delete all orders and associated cart-items connected to this user-id.
	err = deleteManyOrders(orders)

	//----> Check for error.
	if err != nil {
		return errors.New("orders cannot be deleted")
	}

	return nil
}

func (*Order) GetAllOrders() ([]Order, error){
	orders := []Order{} //----> Orders variable.

	//----> Retrieve orders from database.
	err := initializers.DB.Model(&Order{}).Preload("User").Preload("CartItems").Find(&orders).Error

	//----> Check for error.
	if err != nil {
		return []Order{}, errors.New("orders are not available in the database")
	}

	//----> Send back response.
	return orders, nil
}

func (*Order) GetAllOrdersByUserId(userId uint) ([]Order, error){
	orders := []Order{} //----> Orders variable.

	//----> Retrieve orders from database.
	err := initializers.DB.Preload("CartItems").Find(&orders, Order{UserID: userId}).Error
	
	//----> Check for error.
	if err != nil {
		return []Order{}, errors.New("orders are not available in the database")
	}

	//----> Send back response.
	return orders, nil
}

func (order *Order) GetOrderById(id uint) (Order, error){
	//----> retrieve the order with the given id from database.
	err := initializers.DB.Model(&Order{}).Preload("CartItems").First(&order, id).Error

	//----> Check for error.
	if err != nil {
		return Order{}, errors.New("order is not available in the database ")
	}

	//----> Send back response.
	return *order, nil
}

func (order *Order) OrderDelivered(id uint) error{
	//----> Retrieve the order.
	err := initializers.DB.First(&order, id).Error

	//----> Check for error.
	if err != nil {
		return errors.New("order cannot be found")
	}

	//----> Update the shipping info.
	err = deliveryInfo(order)

	//----> Check for error.
	if err != nil{
		return errors.New("delivery info cannot be changed")
	}

	//----> send back the response.
	return nil
}

func (order *Order) OrderShipped(id uint) error{
	//----> Retrieve the order.
	err := initializers.DB.First(&order, id).Error

	//----> Check for error.
	if err != nil {
		return errors.New("order cannot be found")
	}

	//----> Update the shipping info.
	err = shippingInfo(order)

	//----> Check for error.
	if err != nil{
		return errors.New("shipping info cannot be changed")
	}

	//----> send back the response.
	return nil
}

type Cart struct{
	//ID uint
	Name     string `json:"name"`
	Price    float64 `json:"price"`
	Quantity float64 `json:"quantity"`
	Image    string `json:"image"`
	PizzaId  uint `json:"pizzaId"`
	OrderId  uint `json:"orderId"`
}

type Carts []Cart

type OrderPayload struct {
	UserId uint `json:"userId"`
	Carts
}

func (order *OrderPayload) CheckOutOrder() error{
	//----> Get the carts slice.
	carts := order.Carts

	//----> Make order struct.
	orderPayload := makeOrder(carts, order.UserId)

	//----> Insert order in the database.
	err := initializers.DB.Create(&orderPayload).Error

	//----> Check for error.
	if err != nil{
		return errors.New("order creation fails")
	}

	//----> Make cart-items from cart-item struct.
	cartItems := makeCart(carts, orderPayload.ID)

	//----> Insert all the cart-items with the given order-id in the database.
	err = initializers.DB.CreateInBatches(&cartItems, len(cartItems)).Error

	//----> Check for error.
	if err != nil{
		return errors.New("cartItems creation fails")
	}

	return nil
}

