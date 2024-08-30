package main

import (
	"example/bookstore/database"
	"example/bookstore/routes"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
	database.ConnectDatabase()
	r := gin.Default()
	routes.RegisterBookRoutes(r)
	r.Run(":8080")
}
