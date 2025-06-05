package models

import (
	"errors"
	"time"

	"github.com/alfredamos/go-meal-api/initializers"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CartItem struct {
	ID        string `gorm:"primaryKey" json:"id"`          
  CreatedAt time.Time
  UpdatedAt time.Time
  DeletedAt gorm.DeletedAt `gorm:"index"`
	Name     string `json:"name" binding:"required"`
	Price    float64 `json:"price" binding:"required"`
	Quantity float64 `json:"quantity" binding:"required"`
	Image    string `json:"image" binding:"required"`
	OrderID  string `gorm:"foreignKey:OrderID" json:"orderId"`
	Order Order 
	PizzaID  string `gorm:"foreignKey:PizzaID" json:"pizzaId" binding:"required"`
	Pizza Pizza 
}

// This functions are called before creating any Post
func (t *CartItem) BeforeCreate(tx *gorm.DB) (err error) {
	t.ID = uuid.New().String()
	return
}

func (cartItem *CartItem) CreateCartItem() error{
	//----> Insert the cart-item into the database.
	err := initializers.DB.Create(&cartItem).Error 

	//----> Check for error.
	if err != nil {
		return errors.New("cart-item is not created")
	}

	//----> Send back response
	return nil
}

func (*CartItem) DeleteCartItemById(id string) error{
	//----> Retrieve the cart-item with the given id.
	_, err := cartItemGetById(id)

	//----> Check for error.
	if err != nil {
		return errors.New("cart-item is not found")
	}

	//----> Delete the cart-item with the given id from database.
	err = initializers.DB.Delete(&CartItem{}, id).Error
	
	//----> Check for error.
	if err != nil {
		return errors.New("cart-item cannot be deleted")
	}

	//----> Send back the response.
	return nil
}

func (cartItem *CartItem) EditCartItemId(id string) error{
	//----> Retrieve the cart-item with the given id.
	_, err := cartItemGetById(id)

	//----> Check for error.
	if err != nil {
		return errors.New("cart-item is not found")
	}

	//----> Update the cart-item in the database.
	err = initializers.DB.Model(&cartItem).Updates(&cartItem).Error

	//----> Check for error.
	if err != nil {
		return errors.New("cart-item cannot be updated")
	}

	//----> Send back the response.
	return nil
}

func (*CartItem) GetAllCartItems() ([]CartItem, error){
	cartItems := []CartItem{} //----> Declaration.

	//----> Retrieve the cart-items from the database.
	err := initializers.DB.Find(&cartItems).Error

	//----> Check for error.
	if err != nil {
		return []CartItem{}, errors.New("cart-items are not found")
	}

	//----> send back the response.
	return cartItems, nil
}

func (*CartItem) GetCartItemById(id string) (CartItem, error){
	//----> Retrieve the cart-item with the given id.
	cartItem, err := cartItemGetById(id)

	//----> Check for error.
	if err != nil {
		return CartItem{}, errors.New("cart-item is not found")
	}

	//----> send back the response.
	return cartItem, nil

}
