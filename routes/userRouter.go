package routes

import(
	controller "github.com/akmmbh/golang-authentication/controllers"
	"github.com/akmmbh/golang-authentication/middleware"
	"github.com/gin-gonic/gin"
)

func UserRoutes(incommingRoutes *gin.Engine){
	//useing middleware to check that both these routes are protected routes

	incommingRoutes.Use(middleware.Authenticate())
	incommingRoutes.GET("/users",controller.GetUsers())
	incommingRoutes.GET("/users/:user_id",controller.GetUser())
}