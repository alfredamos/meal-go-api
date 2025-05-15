package main

import (
	"github.com/alfredamos/go-meal-api/initializers"
	"github.com/alfredamos/go-meal-api/routes"
	"github.com/gin-gonic/gin"
)

func init(){
	initializers.LoadEnvVariable()
	initializers.ConnectDB()
}


func main(){
	//----> Set the gin server.
	server := gin.Default()

	//---->Get the end-points
	routes.RegisterAllRoutes(server)

	server.Run()
}