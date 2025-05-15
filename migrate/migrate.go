package main

import (
	"github.com/alfredamos/go-meal-api/initializers"
	"github.com/alfredamos/go-meal-api/models"
)

func init() {
	initializers.LoadEnvVariable()
	initializers.ConnectDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.User{}, &models.Order{}, &models.Pizza{}, &models.CartItem{})
	
}
