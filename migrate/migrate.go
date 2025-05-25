package main

import (
	"github.com/alfredamos/go-meal-api/initializers"
	"github.com/alfredamos/go-meal-api/models"
)

func init() {
	//----> Get all environment variables.
	initializers.LoadEnvVariable()

	//----> Connect to database.
	initializers.ConnectDB()
}

func main() {
	//----> Migrate the gorm models into mysql database.
	initializers.DB.AutoMigrate(&models.User{}, &models.Order{}, &models.Pizza{}, &models.CartItem{})
	
}
