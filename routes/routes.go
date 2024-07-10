package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jpuseche/ecommerce/controllers"
)

func UserRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.POST("/users/signup", controllers.SignUp())
	incomingRoutes.POST("/users/login", controllers.Login())
	incomingRoutes.POST("/admin/addproduct", controllers.AddProductToCart())
	incomingRoutes.GET("/users/viewproduct", controllers.GetProductFromCart())
	incomingRoutes.GET("/users/searchproduct", controllers.SearchProductByQuery())
}
