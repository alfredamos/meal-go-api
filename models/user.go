package models

import (
	"errors"
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
	
	//----> Check for empty slice of user.
	if result.RowsAffected == 0 {
		return []User{}, errors.New("there are no users to retrieve from database")
	}

	//----> Send back the response.
   return users, nil    
}

func (*User)GetUserById(id uint) (User, error) {
	var user User //----> User variable.
	//----> Retrieve the user with the given id from the database.
	result := initializers.DB.First(&user, id)
	
	//----> Check for non existent user.
	if result.RowsAffected == 0 {
		return User{}, errors.New("there is no user with the given id to retrieve from database")
	}
	
	//----> Send back the response.
   return user, nil
}

func (*User)DeleteUserById(id uint) error{
	var user User //----> User variable.
	//----> Retrieve the user with the given id from the database.
	result := initializers.DB.First(&user, id)
	
	//----> Check for non existent user.
	if result.RowsAffected == 0 {
		return errors.New("there is no user with the given id to retrieve from database")
	}

	//----> Delete the user.
	result = initializers.DB.Unscoped().Delete(&User{}, id)

	//----> Check for error.
	if result.RowsAffected == 0 {
		return errors.New("this user cannot be deleted")
	}
	
	return nil
}