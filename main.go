package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/pranjalch99/ecommerce-golang/controllers"
	"github.com/pranjalch99/ecommerce-golang/database"
	"github.com/pranjalch99/ecommerce-golang/middleware"
	"github.com/pranjalch99/ecommerce-golang/routes"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	app := controllers.NewApplication(database.ProductData(database.Client, "Products"), database.UserData(database.Client, "Users"))

	//router is created without any middleware by default
	router := gin.New()

	//Global middleware
	router.Use(gin.Logger())

	//Recovery middleware recovers from any panics and returns 500 if there was one.
	router.Use(gin.Recovery())

	routes.UserRoutes(router)
	router.Use(middleware.Authentication())

	router.GET("/addtocart", app.AddToCart())
	router.GET("/removeitem", app.RemoveItem())
	router.GET("/cartcheckout", app.BuyFromCart())
	router.GET("/instantbuy", app.InstantBuy())

	log.Fatal(router.Run(":" + port))
}
