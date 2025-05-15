package models

import (
	"fmt"

	"gorm.io/gorm"
)

type CartItem struct {
	gorm.Model
	Name     string
	Price    float64
	Quantity float64
	Image    string
	OrderID  uint
	Order Order 
	PizzaID  uint
	Pizza Pizza 
}

func (cartItem *CartItem) CreateCartItem(){
	fmt.Println("CartItem : ", cartItem)
}

func (*CartItem) DeleteCartItemById(id uint){
	fmt.Println("Id : ", id)
}

func (cartItem *CartItem) EditCartItemId(id uint){
	fmt.Println("CartItem : ", cartItem, " id : ", id)
}

func (*CartItem) GetAllCartItems(){
	fmt.Println("Get all cart-items")
}

func (*CartItem) GetCartItemById(id uint){
	fmt.Println("Id : ", id)
}