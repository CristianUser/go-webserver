package main

import (
	"log"
	"pronesoft/server/api"
	"pronesoft/server/model"

	"github.com/gin-gonic/gin"
)

func main() {

	db, err := model.Database()
	if err != nil {
		log.Println(err)
	}
	db.DB()

	router := gin.Default()
	api.SetTodoRoutes(router)

	log.Fatal(router.Run(":8080"))
}
