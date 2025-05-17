package models

import (
	"fmt"
	"time"

	"gorm.io/gorm"
)

type Status string

const (
	Delivered Status = "Delivered"
	Pending Status = "Pending"
	Shipped Status = "Shipped"
)

type Order struct {
	gorm.Model
	PaymentId string `gorm:"null" json:"paymentId"`
	OrderDate time.Time `gorm:"null" json:"orderDate"`
	ShippingDate time.Time `gorm:"null" json:"shippingDate"`
	DeliveryDate time.Time `gorm:"null" json:"deliveryDate"`
	TotalQuantity float64 `gorm:"null" json:"totalQuantity"`
	TotalPrice float64 `gorm:"null" json:"totalPrice"`
	IsShipped bool `gorm:"null" json:"isShipped"`
	IsPending bool `gorm:"null" json:"isPending"`
	IsDelivered bool `gorm:"null" json:"isDelivered"`
	Status Status `json:"status" binding:"required"`
	UserID uint `json:"userId" binding:"required"`
	User User 
	CartItems []CartItem `gorm:"foreignKey:OrderID"`

}

type OrderPayload struct{
	TotalQuantity float64
	TotalPrice float64
	UserID uint
	CartItems []CartItem
}

func (order *OrderPayload) CreateOrder(){
	//var cartItems []CartItem
	cartItems := order.CartItems

	for _, cartItem := range cartItems {
		fmt.Printf("%+v", cartItem)
	}
	
}

func (*Order) DeleteOrderById(id uint){
	fmt.Println("Id : ", id)
}

func (*Order) DeleteOrderByUserId(userId uint){
	fmt.Println("Id : ", userId)
}

func (order *Order) EditOrderId(id uint){
	fmt.Println("Order : ", order, " id : ", id)
}

func (*Order) GetAllOrders(){
	fmt.Println("Get all orders")
}

func (*Order) GetAllOrdersByUserId(userId uint){
	fmt.Println("Get orders by userId : ", userId)
}

func (*Order) GetOrderById(id uint){
	fmt.Println("Get order by id : ", id)
	fmt.Println("Id : ", id)
}

func (*Order) OrderDelivered(id uint){
	fmt.Println("Order is delivered, id : ", id)
}

func (*Order)OrderShipped(id uint){
	fmt.Println("Order is shipped, id : ", id)
}

