package models

import (
	"fmt"

	"gorm.io/gorm"
	//"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string
	Email    string `gorm:"unique"`
	Phone    string
	Image    string
	Gender   Gender
	Password string
	Role     Role
	Address  string
	Orders []Order `gorm:"foreignKey:UserID"`
	Pizzas []Pizza `gorm:"foreignKey:UserID"`

}

func (*User) GetAllUsers() {
	fmt.Println("Get all Users ")
}

func (*User)GetUserById(id uint) {
	fmt.Println("Get user by id : ", id)
}

func (*User)DeleteUserById(id uint){
	fmt.Println("Delete user by id : ", id)
}