package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jpuseche/ecommerce/controllers"
	"github.com/jpuseche/ecommerce/database"
	"github.com/jpuseche/ecommerce/middleware"
	"github.com/jpuseche/ecommerce/routes"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	app := controllers.NewApplication(database.CollectionData(database.Client, "Products"), database.CollectionData(database.Client, "Users"))

	router := gin.New()
	router.Use(gin.Logger())

	routes.UserRoutes(router)
	router.Use(middleware.Authentication())

	router.GET("/addproducttocart", app.AddProductToCart())
	router.GET("/removeproduct", app.RemoveProduct())
	router.GET("/cartcheckout", app.CartCheckout())
	router.GET("/instantbuy", app.InstantBuy())

	log.Fatal(router.Run(":" + port))
}
