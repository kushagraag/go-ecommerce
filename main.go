package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	controller "github.com/kushagraag/golang-ecommerce-cart/controllers"
	database "github.com/kushagraag/golang-ecommerce-cart/database"
	middleware "github.com/kushagraag/golang-ecommerce-cart/middleware"
	routes "github.com/kushagraag/golang-ecommerce-cart/routes"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	// a new controller instance is created with db access and stored in app var so app can be used to
	// access all funcs of controller pacakge which has address, cart, controllers go file
	app := controller.NewApplication(database.ProductData(database.Client, "Products"), database.UserData(database.Client, "Users"))

	router := gin.New()
	router.Use(gin.Logger())

	routes.UserRoutes(router)
	router.Use(middleware.Authentication())

	router.GET("/addtocart", app.AddToCart())
	router.GET("/removeitem", app.RemoveItem())
	router.GET("/cartcheckout", app.BuyFromCart())
	router.GET("/instantbuy", app.InstantBuy())

	log.Fatal(router.Run(":" + port))

}
