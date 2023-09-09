package routes

import (
	"github.com/gin-gonic/gin"
)

func Router(router *gin.Engine) *gin.Engine {
	api := router.Group("/api")
	UserRouter(api)
	return router
}
