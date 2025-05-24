package routes

import(
	controller "github.com/akmmbh/golang-authentication/controllers"
	"github.com/gin-gonic/gin"
)
func AuthRoutes(incommingRoutes *gin.Engine){
incommingRoutes.POST("users/signup",controller.Signup)
incommingRoutes.POST("users/logic",controller.Login())

}