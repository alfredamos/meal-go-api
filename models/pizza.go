package models

import (
	"errors"
	"github.com/alfredamos/go-meal-api/initializers"
	"gorm.io/gorm"
)

func pizzaGetById(id uint) (Pizza, error){
	var pizza Pizza //----> Pizza variable.
	
	//----> Retrieve the pizza with the given id from the database.
	result := initializers.DB.First(&pizza, id)
	
	//----> Check for non existent pizza.
	if result.RowsAffected == 0 {
		return Pizza{}, errors.New("the pizza with the given id is not found")
	}
	
	//----> Send back the response.
   return pizza, nil
}


type Pizza struct {
	gorm.Model
	Name        string `json:"name" binding:"required"`
	Topping     string `json:"topping" binding:"required"`
	Price       float64 `json:"price" binding:"required"`
	Quantity    float64 `json:"quantity" binding:"required"`
	Image       string `json:"image" binding:"required"`
	Description string `json:"description" binding:"required"`
	UserID      uint `json:"userId" binding:"required"`
	User User 
}

func (pizza *Pizza) CreatePizza() error{
	//---->  Create the pizza.
	result := initializers.DB.Create(&pizza)

	//----> Check for error.
	if result.RowsAffected == 0{
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
	result := initializers.DB.Unscoped().Delete(&Pizza{}, id)

	//----> Check for error.
	if result.RowsAffected == 0{
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
	result := initializers.DB.Model(&pizza).Updates(&pizza)

	//----> Check for error.
	if result.RowsAffected == 0 {
		return errors.New("pizza cannot be updated")
	}

	//----> send back the response.
	return nil
}

func (*Pizza) GetAllPizzas() ([]Pizza, error){
	var pizzas []Pizza //----> Pizza variable.

	//----> Retrieve pizzas from database.
	result := initializers.DB.Find(&pizzas)

	//----> Check for error.
	if result.RowsAffected == 0 {
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