package controllers

import (
	"net/http"
	"strconv"

	"github.com/alfredamos/go-meal-api/models"
	"github.com/gin-gonic/gin"
)

func CreatePizza(context *gin.Context) {
 //----> Get the type
 var pizza models.Pizza

 //----> Get the request payload
 err := context.ShouldBindJSON(&pizza)

 //----> Check for error.
 if err != nil {
	context.JSON(http.StatusBadRequest, gin.H{"message": "All values must be provided!"})
	return
 }

 //----> Save the new pizza into the database.
 err = pizza.CreatePizza()

 //----> Check for error.
 if err != nil {
	context.JSON(http.StatusBadRequest, gin.H{"message": "Pizza cannot be created!"})
	return
 }

 //----> Send back the response
 context.JSON(http.StatusCreated, gin.H{"message": "Pizza has been created successfully!"})

}

func DeletePizzaById(context *gin.Context) {
	//----> Get the type
 var pizza models.Pizza

 //----> Get the id from params.
 idd := context.Param("id")
 id, err := strconv.ParseUint(idd, 10, 32)

 //----> Check for error
 if err != nil {
	context.JSON(http.StatusBadRequest, gin.H{"message": "Please provide a valid id!"})
	return
 }

 //----> Delete pizza with this id from the database.
 err = pizza.DeletePizzaById(uint(id))

 //----> Check for error.
 if err != nil {
	context.JSON(http.StatusNotFound, gin.H{"message": "Pizza cannot be deleted!"})
	return
 }

 //----> Send back the response
 context.JSON(http.StatusOK, gin.H{"message": "Pizza has been deleted successfully!"})

}

func EditPizzaById(context *gin.Context) {
 //----> Get the type
 var pizza models.Pizza

 //----> Get the id from params.
 idd := context.Param("id")
 id, err := strconv.ParseUint(idd, 10, 32)

 //----> Check for error
 if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Please provide a valid id!"})
		return
	}

 //----> Get the request payload
 err = context.ShouldBindJSON(&pizza)

 //----> Check for error.
 if err != nil {
	context.JSON(http.StatusBadRequest, gin.H{"message": "All values must be provided!"})
	return
 }

 //----> Save the edited pizza into the database.
 err = pizza.EditPizzaId(uint(id))

 if err != nil {
	context.JSON(http.StatusNotFound, gin.H{"message": "Pizza cannot be updated!"})
	return
}

 //----> Send back the response
 context.JSON(http.StatusOK, gin.H{"message": "Pizza has been edited successfully!"})
}

func GetAllPizza(context *gin.Context) {
	//----> Get the type
	var pizza models.Pizza
	
	//----> Get all pizzas from database.
	pizzas, err := pizza.GetAllPizzas()

	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"message": "Pizzas cannot be retrieved from database!"})
		return
	}
	
	//----> Send back the response
	context.JSON(http.StatusOK, gin.H{"message": "Pizzas are retrieved successfully!", "pizzas": pizzas})
}

func GetPizzaById(context *gin.Context) {
	//----> Get the type
	var pizza models.Pizza
	
	//----> Get the id from params.
	idd := context.Param("id")
	id, err:= strconv.ParseUint(idd, 10, 32)

 //----> Check for error
 if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Please provide a valid id!"})
		return
 }

 //----> Get pizza with this id from database.
 pizza, err = pizza.GetPizzaById(uint(id))

 //----> Check for error.
 if err != nil {
	context.JSON(http.StatusNotFound, gin.H{"message": "Pizza cannot be retrieved from database!"})
	return
 }

 //----> Send back the response
 context.JSON(http.StatusOK, gin.H{"message": "Pizza is retrieved successfully!", "pizza" : pizza})

}