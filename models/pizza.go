package models

import (
	"fmt"

	"gorm.io/gorm"
)

type Pizza struct {
	gorm.Model
	Name        string
	Topping     string
	Price       float64
	Quantity    float64
	Image       string
	Description string
	UserID      uint
	User User 
}


func (pizza *Pizza) CreatePizza(){
	fmt.Println("Pizza : ", pizza)
}

func (*Pizza) DeletePizzaById(id uint){
	fmt.Println("Id : ", id)
}

func (pizza *Pizza) EditPizzaId(id uint){
	fmt.Println("Pizza : ", pizza, " id : ", id)
}

func (*Pizza) GetAllPizzas(){
	fmt.Println("Get all pizzas! ")
}

func (*Pizza) GetPizzaById(id uint){
	fmt.Println("Get pizza by Id : ", id)
}