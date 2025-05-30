package models

import (
	"errors"
	"time"

	"github.com/alfredamos/go-meal-api/initializers"
	"gorm.io/gorm"
)


type Pizza struct {
	ID        uint `gorm:"primaryKey" json:"id"`          
  CreatedAt time.Time
  UpdatedAt time.Time
  DeletedAt gorm.DeletedAt `gorm:"index"`
	Name        string `json:"name" binding:"required"`
	Topping     string `json:"topping" binding:"required"`
	Price       float64 `json:"price" binding:"required"`
	Quantity    float64 `json:"quantity" binding:"required"`
	Image       string `json:"image" binding:"required"`
	Description string `json:"description" binding:"required"`
	UserID      uint `gorm:"foreignKey:UserID;size:191" json:"userId" binding:"required"`
}

func (pizza *Pizza) CreatePizza() error{
	//---->  Create the pizza.
	err := initializers.DB.Create(&pizza).Error

	//----> Check for error.
	if err != nil{
		return errors.New("pizza creation fails")
	}

	//----> Send back the response.
	return nil
}

func (*Pizza) DeletePizzaById(id uint) error{
	//----> Get the pizza with the given id
	_, err := pizzaGetById(id)

	//----> Check for error.
	if err != nil {
		return errors.New("pizza cannot be retrieved")
	}

	//----> Delete pizza from the database.
	err = initializers.DB.Unscoped().Delete(&Pizza{}, id).Error

	//----> Check for error.
	if err != nil{
		return errors.New("pizza cannot be deleted")
	}

	//----> send back the response.

	return nil
}

func (pizza *Pizza) EditPizzaId(id uint) error{
	//----> Get the pizza with the given id
	_ , err := pizzaGetById(id)

	//----> Check for error.
	if err != nil {
		return errors.New("pizza cannot be retrieved")
	}

	//----> Update the cart-item in the database.
	err = initializers.DB.Model(&pizza).Updates(&pizza).Error

	//----> Check for error.
	if err != nil {
		return errors.New("pizza cannot be updated")
	}

	//----> send back the response.
	return nil
}

func (*Pizza) GetAllPizzas() ([]Pizza, error){
	pizzas := []Pizza{} //----> Pizza variable.

	//----> Retrieve pizzas from database.
	err := initializers.DB.Find(&pizzas).Error

	//----> Check for error.
	if err != nil {
		return []Pizza{}, errors.New("there are no pizzas to retrieve from database")
	}

	//----> Send back the response.
   return pizzas, nil   
}

func (*Pizza) GetPizzaById(id uint) (Pizza, error){
	//----> Get the pizza with the given id
	pizza, err := pizzaGetById(id)

	//----> Check for error.
	if err != nil {
		return Pizza{}, errors.New("pizza cannot be retrieved")
	}

	//----> Send back the response.
	return pizza, nil
}