package main

import (
	"b2life/internal/db"
	"b2life/middlewares"
	"b2life/routes"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: No .env file found")
	}

	database := db.InitDB()
	db.RunMigrations(database)

	router := gin.Default()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Use(middlewares.CorsMiddleware())

	routes.ApplyRoutes(router)

	addr := db.GetEnv("SERVER_PORT")
	log.Printf("Server starting on %s", addr)
	if err := router.Run(":" + addr); err != nil {
		log.Fatal("Failed to start server: ", err)
	}
	return
}
