package routes

import (
	user_controller "github.com/api-rest-go/controllers/user.controller"
	"github.com/gin-gonic/gin"
)

func UserRouter(router *gin.RouterGroup) {
	user := router.Group("/user")
	user.GET("", user_controller.GetUsers)
	user.POST("", user_controller.PostUser)
	user.GET("/:id", user_controller.GetUserById)
	user.PUT("/:id", user_controller.Update)
	user.DELETE("/:id", user_controller.Delete)
}
