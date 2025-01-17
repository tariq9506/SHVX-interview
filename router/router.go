package router

import (
	"shvx/controllers"

	"github.com/gin-gonic/gin"
)

func AddRouter(router *gin.RouterGroup) {
	router.POST("/shvx/register", controllers.UserSignUP)
	router.POST("/shvx/sign-in", controllers.UserSignIn)
}
func SetupRouter() *gin.Engine {
	router := gin.Default()
	AddRouter(&router.RouterGroup)
	return router

}
