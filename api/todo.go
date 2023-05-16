package api

import (
	"pronesoft/server/controller"

	"github.com/gin-gonic/gin"
)

func SetTodoRoutes(router *gin.Engine) {

	router.GET("/todos", controller.GetTodos)
	router.GET("/todo/:id", controller.GetTodo)
	router.POST("/todo", controller.PostTodo)
	router.PUT("/todo/:id", controller.UpdateTodo)
	router.DELETE("/todo/:id", controller.DeleteTodo)
}
