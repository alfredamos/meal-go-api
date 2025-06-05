package models

import (
	"errors"
	"time"

	"github.com/alfredamos/go-meal-api/initializers"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID        string `gorm:"primaryKey" json:"id"`          
  CreatedAt time.Time
  UpdatedAt time.Time
  DeletedAt gorm.DeletedAt `gorm:"index"`
	Name     string `json:"name" binding:"required"`
	Email    string `gorm:"unique" json:"email" binding:"required"`
	Phone    string `json:"phone" binding:"required"`
	Image    string `json:"image" binding:"required"`
	Gender   Gender `json:"gender" binding:"required"`
	Password string `json:"password" binding:"required"`
	Role     Role `json:"role"`
	Address  string `json:"address" binding:"required"`
	Orders []Order `gorm:"foreignKey:UserID"`
	Pizzas []Pizza `gorm:"foreignKey:UserID"`

}

// This functions are called before creating any Post
func (t *User) BeforeCreate(tx *gorm.DB) (err error) {
	t.ID = uuid.New().String()
	return
}

func (user *User) GetAllUsers() ([]User, error) {
	//----> Declare slice of users.
	users := []User{}
	
	//----> Retrieve the users from the database.
	err := initializers.DB.Omit("Password").Find(&users).Error
	
	//----> Check for empty slice of user.
	if err != nil {
		return []User{}, errors.New("there are no users to retrieve from database")
	}

	//----> Send back the response.
   return users, nil    
}

func (*User) GetUserById(id string) (User, error) {
	//----> Get user with the given id.
	user, err := userGetById(id)

	//----> Check for error.
	if err != nil {
		return User{}, errors.New("user cannot be retrieved")
	}

	//----> Send back the response.
	 return user, nil
}

func (*User) DeleteUserById(id string) error{
	//----> Get user with the given id.
	_, err := userGetById(id)

	//----> Check for error.
	if err != nil {
		return errors.New("pizza cannot be retrieved")
	}

	//----> Delete the user.
	err = initializers.DB.Unscoped().Delete(&User{}, id).Error

	//----> Check for error.
	if err != nil {
		return errors.New("this user cannot be deleted")
	}
	
	return nil
}