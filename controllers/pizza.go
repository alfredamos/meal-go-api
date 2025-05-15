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

 //----> Save the changed profiles into the database.
 pizza.CreatePizza()

 //----> Send back the response
 context.JSON(http.StatusCreated, gin.H{"message": "Pizza has been created successfully!"})

}

func DeletePizzaById(context *gin.Context) {
	//----> Get the type
 var pizza models.Pizza

 //----> Get the id from params.
 idd := context.Param("id")
 id, errId:= strconv.ParseUint(idd, 10, 32)

 //----> Check for error
 if errId != nil {
	context.JSON(http.StatusBadRequest, gin.H{"message": "Please provide a valid id!"})
	return
 }

 //----> Delete pizza with this id from the database.
 pizza.DeletePizzaById(uint(id))

 //----> Send back the response
 context.JSON(http.StatusOK, gin.H{"message": "Pizza has been deleted successfully!"})

}

func EditPizzaById(context *gin.Context) {
 //----> Get the type
 var pizza models.Pizza

 //----> Get the id from params.
 idd := context.Param("id")
 id, errId:= strconv.ParseUint(idd, 10, 32)

 //----> Check for error
 if errId != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Please provide a valid id!"})
		return
	}

 //----> Get the request payload
 err := context.ShouldBindJSON(&pizza)

 //----> Check for error.
 if err != nil {
	context.JSON(http.StatusBadRequest, gin.H{"message": "All values must be provided!"})
	return
 }

 //----> Save the changed profiles into the database.
 pizza.EditPizzaId(uint(id))

 //----> Send back the response
 context.JSON(http.StatusOK, gin.H{"message": "Pizza has been edited successfully!"})
}

func GetAllPizza(context *gin.Context) {
	//----> Get the type
	var pizza models.Pizza
	
	//----> Get all pizzas from database.
	pizza.User.GetAllUsers()

	//----> Send back the response
	context.JSON(http.StatusOK, gin.H{"message": "Pizzas are retrieved successfully!"})
}

func GetPizzaById(context *gin.Context) {
	//----> Get the type
	var pizza models.Pizza
	
	//----> Get the id from params.
	idd := context.Param("id")
	id, errId:= strconv.ParseUint(idd, 10, 32)

 //----> Check for error
 if errId != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Please provide a valid id!"})
		return
 }

 //----> Get pizza with this id from database.
 pizza.GetPizzaById(uint(id))

 //----> Send back the response
 context.JSON(http.StatusOK, gin.H{"message": "Pizza is retrieved successfully!"})

}