package controllers

import (
	"net/http"
	"strconv"
	"github.com/alfredamos/go-meal-api/models"
	"github.com/gin-gonic/gin"
)

func CreatePizza(context *gin.Context) {
 //----> Get the type
 pizza := models.Pizza{}

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
 context.JSON(http.StatusCreated, gin.H{"status": "success", "message": "Pizza has been created successfully!", "statusCode": http.StatusCreated})

}

func DeletePizzaById(context *gin.Context) {
	//----> Get the type
 pizza := models.Pizza{}

 //----> Get the pizza-id from params.
 id, err := strconv.Atoi(context.Param("id"))

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
 context.JSON(http.StatusNoContent, gin.H{"status": "success", "message": "Pizza has been deleted successfully!", "statusCode": http.StatusNoContent})

}

func EditPizzaById(context *gin.Context) {
 //----> Get the type
 pizza := models.Pizza{}

 //----> Get the pizza-id from params.
 id, err := strconv.Atoi(context.Param("id"))

 //----> Check for error.
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
 context.JSON(http.StatusNoContent, gin.H{"status": "success", "message": "Pizza has been edited successfully!", "statusCode": http.StatusNoContent})
}

func GetAllPizza(context *gin.Context) {
	//----> Get the type
	pizza := models.Pizza{}
	
	//----> Get all pizzas from database.
	pizzas, err := pizza.GetAllPizzas()

	//----> Check for error.
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"message": "pizzas are not available in the database!"})
		return
	}
	
	//----> Send back the response
	context.JSON(http.StatusOK, pizzas)
}

func GetPizzaById(context *gin.Context) {
	//----> Get the type
	pizza := models.Pizza{}
	
	//----> Get the pizza-id from params.
	id, err := strconv.Atoi(context.Param("id"))

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
 context.JSON(http.StatusOK, pizza)

}