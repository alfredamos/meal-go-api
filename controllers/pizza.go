package controllers

import (
	"fmt"
	"net/http"
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
	context.JSON(http.StatusBadRequest, gin.H{"status": "failed!", "message": fmt.Sprintf("%v", err)})
	return
 }

 //----> Save the new pizza into the database.
 err = pizza.CreatePizza()

 //----> Check for error.
 if err != nil {
	context.JSON(http.StatusBadRequest, gin.H{"status": "failed!", "message": fmt.Sprintf("%v", err)})
	return
 }

 //----> Send back the response
 context.JSON(http.StatusCreated, gin.H{"status": "success", "message": "Pizza has been created successfully!", "statusCode": http.StatusCreated})

}

func DeletePizzaById(context *gin.Context) {
	//----> Get the type
 pizza := models.Pizza{}

 //----> Get the pizza-id from params.
 id:= context.Param("id")
 
 //----> Delete pizza with this id from the database.
 err := pizza.DeletePizzaById(id)

 //----> Check for error.
 if err != nil {
	context.JSON(http.StatusNotFound, gin.H{"status": "failed!", "message": fmt.Sprintf("%v", err)})
	return
 }

 //----> Send back the response
 context.JSON(http.StatusNoContent, gin.H{"status": "success", "message": "Pizza has been deleted successfully!", "statusCode": http.StatusNoContent})

}

func EditPizzaById(context *gin.Context) {
 //----> Get the type
 pizza := models.Pizza{}

 //----> Get the pizza-id from params.
 id := context.Param("id")
 
 //----> Get the request payload
 err := context.ShouldBindJSON(&pizza)

 //----> Check for error.
 if err != nil {
	context.JSON(http.StatusBadRequest, gin.H{"status": "failed!", "message": fmt.Sprintf("%v", err)})
	return
 }

 //----> Save the edited pizza into the database.
 err = pizza.EditPizzaId(id)

 if err != nil {
	context.JSON(http.StatusNotFound, gin.H{"status": "failed!", "message": fmt.Sprintf("%v", err)})
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
		context.JSON(http.StatusNotFound, gin.H{"status": "failed!", "message": fmt.Sprintf("%v", err)})
		return
	}
	
	//----> Send back the response
	context.JSON(http.StatusOK, pizzas)
}

func GetPizzaById(context *gin.Context) {
	//----> Get the type
	pizza := models.Pizza{}
	
	//----> Get the pizza-id from params.
	id := context.Param("id")
 
 //----> Get pizza with this id from database.
 pizza, err := pizza.GetPizzaById(id)

 //----> Check for error.
 if err != nil {
	context.JSON(http.StatusNotFound, gin.H{"status": "failed!", "message": fmt.Sprintf("%v", err)})
	return
 }

 //----> Send back the response
 context.JSON(http.StatusOK, pizza)

}