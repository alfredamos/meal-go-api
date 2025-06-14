package models

import (
	"database/sql"
	"errors"
	"time"

	"github.com/alfredamos/go-meal-api/initializers"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Status string

const (
	Delivered Status = "Delivered"
	Pending Status = "Pending"
	Shipped Status = "Shipped" 
)

type Order struct {
	ID        string `gorm:"primaryKey;type:varchar(255)" json:"id"`          
  CreatedAt time.Time
  UpdatedAt time.Time
  DeletedAt gorm.DeletedAt `gorm:"index"`
	PaymentId string `json:"paymentId"`
	OrderDate time.Time `json:"orderDate"`
	ShippingDate sql.NullTime `gorm:"type:TIMESTAMP NULL" json:"shippingDate"`
	DeliveryDate sql.NullTime `gorm:"type:TIMESTAMP NULL" json:"deliveryDate"`
	TotalQuantity float64 `json:"totalQuantity"`
	TotalPrice float64 `json:"totalPrice"`
	IsShipped bool `json:"isShipped"`
	IsPending bool `json:"isPending"`
	IsDelivered bool `json:"isDelivered"`
	Status Status `json:"status" binding:"required"`
	UserID string `gorm:"foreignKey:UserID;type:varchar(255)" json:"userId" binding:"required"`
	User User `json:"user"`
	CartItems []CartItem `gorm:"foreignKey:OrderID" json:"cartItems"`

}

// This functions are called before creating any Post
func (t *Order) BeforeCreate(tx *gorm.DB) (err error) {
	t.ID = uuid.New().String()
	return
}

func (order *Order) DeleteOrderById(id string) error{
	//----> Check to see if the order to be deleted is available in the database.
	err := initializers.DB.Model(&Order{}).Preload("CartItems").First(&order, "id = ?", id).Error
	
	//----> Check for error.
	if err != nil {
		return errors.New("order does not exist")
	}

	//----> Get the cart-items from order.
	carts := order.CartItems 

	//----> Delete all the cart-items attached to order with the given id.
	 err = deleteManyCartItems(carts)

	//----> Check for error.
	if err != nil {
		return errors.New("cart-items cannot be deleted")
	} 
 
	//----> Delete the order with given id.
	err = initializers.DB.Unscoped().Delete(&Order{}, "id = ?", id).Error

	//----> Check for error.
	if err != nil {
		return errors.New("order cannot be deleted")
	}

	return nil
}

func (*Order) DeleteOrderByUserId(userId string) error{
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

func (*Order) GetAllOrdersByUserId(userId string) ([]Order, error){
	orders := []Order{} //----> Orders variable.

	//----> Retrieve orders from database.
	err := initializers.DB.Preload("User").Preload("CartItems").Find(&orders, Order{UserID: userId}).Error
	
	//----> Check for error.
	if err != nil {
		return []Order{}, errors.New("orders are not available in the database")
	}

	//----> Send back response.
	return orders, nil
}

func (order *Order) GetOrderById(id string) (Order, error){
	//----> retrieve the order with the given id from database.
	err := initializers.DB.Model(&Order{}).Preload("User").Preload("CartItems").First(&order, "id = ?", id).Error

	//----> Check for error.
	if err != nil {
		return Order{}, errors.New("order is not available in the database ")
	}

	//----> Send back response.
	return *order, nil
}

func (order *Order) OrderDelivered(id string) (Order, error){
	//----> Retrieve the order.
	err := initializers.DB.First(&order, "id = ?", id).Error

	//----> Check for error.
	if err != nil {
		return Order{}, errors.New("order cannot be found")
	}

	//----> Update the shipping info.
	orderEdited , err := deliveryInfo(order)

	//----> Check for error.
	if err != nil{
		return Order{}, errors.New("delivery info cannot be changed")
	}

	//----> send back the response.
	return orderEdited, nil
}

func (order *Order) OrderShipped(id string) (Order, error){
	//----> Retrieve the order.
	err := initializers.DB.First(&order, "id = ?", id).Error

	//----> Check for error.
	if err != nil {
		return Order{}, errors.New("order cannot be found")
	}

	//----> Update the shipping info.
	orderEdited, err := shippingInfo(order)

	//----> Check for error.
	if err != nil{
		return Order{}, errors.New("shipping info cannot be changed")
	}

	//----> send back the response.
	return orderEdited, nil
}

type OrderPayload struct {
	UserId string `json:"userId"`
	PaymentId string `json:"paymentId"`
	CartItems []CartItem 
}

func (order *OrderPayload) CheckOutOrder() error{
	//----> Get the carts slice.
	carts := order.CartItems //----> Cart-items.
	userId := order.UserId //----> User-id
	paymentId := order.PaymentId //----> Payment-id

	//----> Make order struct.
	orderPayload := makeOrder(userId, carts, paymentId)

	//----> Insert order in the database.
	err := initializers.DB.Create(&orderPayload).Error

	//----> Check for error.
	if err != nil{
		return errors.New("order creation fails")
	}

	//----> Get the orderPayload-id
  orderPayloadId := orderPayload.ID

	//----> Make cart-items from cart-item struct.
	cartItems := makeCartItems(carts, orderPayloadId)

	//----> Insert all the cart-items with the given order-id in the database.
	err = initializers.DB.CreateInBatches(&cartItems, len(cartItems)).Error

	//----> Check for error.
	if err != nil{
		return errors.New("cartItems creation fails")
	}

	return nil
}

