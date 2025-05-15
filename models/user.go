package models

import (
	"errors"
	"fmt"

	"github.com/alfredamos/go-meal-api/initializers"
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

func (user *User) GetAllUsers() ([]User, error) {
	//----> Declare slice of users.
	var users []User
	
	//----> Retrieve the users from the database.
	result := initializers.DB.Find(&users)
	
	fmt.Println("error : ", result.Error)
	fmt.Println("rowAffected : ", result.RowsAffected)
	//----> Check for empty slice of user.
	if result.RowsAffected == 0 {
		return []User{}, errors.New("there are no users to retrieve from database")
	}
	 fmt.Println("users : ", users)
	//----> Send back the response.
   return users, nil    
}

func (*User)GetUserById(id uint) {
	fmt.Println("Get user by id : ", id)
}

func (*User)DeleteUserById(id uint){
	fmt.Println("Delete user by id : ", id)
}