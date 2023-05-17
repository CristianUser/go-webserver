package api

import (
	"pronesoft/server/controller"

	"github.com/gin-gonic/gin"
)

func SetUserRoutes(router, protected *gin.RouterGroup) {

	router.GET("/users", controller.GetUsers)
	router.GET("/user/:id", controller.GetUser)
	router.POST("/user", controller.PostUser)
	router.PUT("/user/:id", controller.UpdateUser)
	router.DELETE("/user/:id", controller.DeleteUser)
	protected.GET("/me", controller.CurrentUser)
}
