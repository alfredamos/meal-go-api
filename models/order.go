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
	PaymentId string
	OrderDate time.Time
	ShippingDate time.Time
	DeliveryDate time.Time
	TotalQuantity float64
	TotalPrice float64
	IsShipped bool
	IsPending bool
	IsDelivered bool
	Status Status
	UserID uint
	User User 
	CartItems []CartItem `gorm:"foreignKey:OrderID"`

}

func (order *Order) CreateOrder(){
	fmt.Println("Order : ", order)
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

