package api

import (
	"pronesoft/server/controller"

	"github.com/gin-gonic/gin"
)

func SetAuthRoutes(router, protected *gin.RouterGroup) {
	router.POST("/login", controller.Login)
	router.POST("/register", controller.PostUser)
	protected.POST("/logout", controller.Logout)
	// protected.POST("/refresh", controller.Refresh)
}
