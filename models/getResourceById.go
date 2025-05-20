package models

import (
	"errors"
	"fmt"

	"github.com/alfredamos/go-meal-api/initializers"
)

func cartItemGetById(id uint) (CartItem, error) {
	var cartItem CartItem //----> Declaration.

	//----> Retrieve the cart-item with given id from database.
	result := initializers.DB.First(&cartItem, id)

	//----> Check for error.
	if result.RowsAffected == 0 {
		return CartItem{}, errors.New("the cart-item with the given id is not found")
	}

	//----> Send back the response.
	return cartItem, nil
}

func pizzaGetById(id uint) (Pizza, error) {
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

func userGetById(id uint) (User, error) {
	var user User //----> User variable.
	//----> Retrieve the user with the given id from the database.
	result := initializers.DB.Omit("Password").First(&user, id)

	//----> Check for non existent user.
	if result.RowsAffected == 0 {
		return User{}, errors.New("there is no user with the given id to retrieve from database")
	}

	//----> Send back the response.
	return user, nil
}

type CartL struct {
	ID uint
}

func getAllCartItemsIds(carts []CartItem)[]CartItem{
	cartItems := make([]CartItem, 0) //----> Slice of cart-ids
	//cartItem := CartL{}
	for _, cart := range carts {
		cartItem := CartItem{ ID: cart.ID}

		cartItems = append(cartItems, cartItem)

	}

	fmt.Println("cartsIds : ", cartItems)

	return cartItems
}
