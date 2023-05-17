package main

import (
	"log"
	"pronesoft/server/api"
	"pronesoft/server/middlewares"
	"pronesoft/server/model"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db, err := model.Database()
	if err != nil {
		log.Println(err)
	}
	db.DB()

	router := gin.Default()
	v1 := router.Group("/api/v1")

	protected := v1.Group("/")
	protected.Use(middlewares.JwtAuthMiddleware())

	api.SetTodoRoutes(v1)
	api.SetUserRoutes(v1, protected)
	api.SetAuthRoutes(v1, protected)

	log.Fatal(router.Run(":8080"))
}
