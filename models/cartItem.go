package models

import (
	"errors"
	"time"

	"github.com/alfredamos/go-meal-api/initializers"
	"gorm.io/gorm"
)



type CartItem struct {
	ID        uint           `gorm:"primaryKey"`
  CreatedAt time.Time
  UpdatedAt time.Time
  DeletedAt gorm.DeletedAt `gorm:"index"`
	Name     string `json:"name" binding:"required"`
	Price    float64 `json:"price" binding:"required"`
	Quantity float64 `json:"quantity" binding:"required"`
	Image    string `json:"image" binding:"required"`
	OrderID  uint `json:"orderId"`
	Order Order 
	PizzaID  uint `json:"pizzaId" binding:"required"`
	Pizza Pizza 
}

func (cartItem *CartItem) CreateCartItem() error{
	//----> Insert the cart-item into the database.
	result := initializers.DB.Create(&cartItem) 

	//----> Check for error.
	if result.RowsAffected == 0 {
		return errors.New("cart-item is not created")
	}

	//----> Send back response
	return nil
}

func (*CartItem) DeleteCartItemById(id uint) error{
	//----> Retrieve the cart-item with the given id.
	_, err := cartItemGetById(id)

	//----> Check for error.
	if err != nil {
		return errors.New("cart-item is not found")
	}

	//----> Delete the cart-item with the given id from database.
	result := initializers.DB.Delete(&CartItem{}, id)
	
	//----> Check for error.
	if result.RowsAffected == 0 {
		return errors.New("cart-item cannot be deleted")
	}

	//----> Send back the response.
	return nil
}

func (cartItem *CartItem) EditCartItemId(id uint) error{
	//----> Retrieve the cart-item with the given id.
	_, err := cartItemGetById(id)

	//----> Check for error.
	if err != nil {
		return errors.New("cart-item is not found")
	}

	//----> Update the cart-item in the database.
	result := initializers.DB.Model(&cartItem).Updates(&cartItem)

	//----> Check for error.
	if result.RowsAffected == 0 {
		return errors.New("cart-item cannot be updated")
	}

	//----> Send back the response.
	return nil
}

func (*CartItem) GetAllCartItems() ([]CartItem, error){
	var cartItems []CartItem //----> Declaration.

	//----> Retrieve the cart-items from the database.
	result := initializers.DB.Find(&cartItems)

	//----> Check for error.
	if result.RowsAffected == 0 {
		return []CartItem{}, errors.New("cart-items are not found")
	}

	//----> send back the response.
	return cartItems, nil
}

func (*CartItem) GetCartItemById(id uint) (CartItem, error){
	//----> Retrieve the cart-item with the given id.
	cartItem, err := cartItemGetById(id)

	//----> Check for error.
	if err != nil {
		return CartItem{}, errors.New("cart-item is not found")
	}

	//----> send back the response.
	return cartItem, nil

}