package models

import (
	"errors"

	"github.com/alfredamos/go-meal-api/initializers"
)

func cartItemGetById(id string) (CartItem, error) {
	cartItem := CartItem{} //----> Declaration.

	//----> Retrieve the cart-item with given id from database.
	err := initializers.DB.First(&cartItem, "id = ?", id).Error

	//----> Check for error.
	if err != nil {
		return CartItem{}, errors.New("the cart-item with the given id is not found")
	}

	//----> Send back the response.
	return cartItem, nil
}

func pizzaGetById(id string) (Pizza, error) {
	pizza := Pizza{} //----> Pizza variable.

	//----> Retrieve the pizza with the given id from the database.
	err := initializers.DB.First(&pizza, "id = ?", id).Error

	//----> Check for non existent pizza.
	if err != nil {
		return Pizza{}, errors.New("the pizza with the given id is not found")
	}

	//----> Send back the response.
	return pizza, nil
}

func userGetById(id string) (User, error) {
	user := User{} //----> User variable.
	
	//----> Retrieve the user with the given id from the database.
	err := initializers.DB.Omit("Password").First(&user, "id = ?", id).Error

	//----> Check for non existent user.
	if err != nil {
		return User{}, errors.New("there is no user with the given id to retrieve from database")
	}

	//----> Send back the response.
	return user, nil
}

func getAllCartItemsIds(carts []CartItem)[]CartItem{
	cartItems := make([]CartItem, 0) //----> Slice of cart-ids
	
	//----> Get all the cart-items ids.
	for _, cart := range carts {
		//----> Compose the id from the cart-item struct.
		cartItem := CartItem{ ID: cart.ID}

		//----> Append all the ids together to have a slice of cart-item ids.
		cartItems = append(cartItems, cartItem)

	}

	//----> Send back the response
	return cartItems
}
