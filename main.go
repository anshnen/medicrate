package main

import (
	"log"
	"os"

	"github.com/anshnen/medicrate/controllers"
	"github.com/anshnen/medicrate/database"
	"github.com/anshnen/medicrate/middleware"
	"github.com/anshnen/medicrate/routes"
	"github.com/gin-gonic/gin"
)

func main(){
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}
	app := controllers.NewApplication(database.ProductData(database.Client, "Products"), database.UserData(database.Client, "Users"))

	router := gin.New()
	router.Use(gin.Logger())
	routes.UserRoutes(router)
	router.Use(middleware.Authentication())
	router.GET("/removeitem", app.RemoveFromCart())
	router.GET("/addtocart", app.AddToCart())
	router.GET("/listcart", controllers.GetItemFromCart())
	router.POST("/addaddress", controllers.AddAddress())
	router.PUT("/edithomeaddress", controllers.EditHomeAddress())
	router.PUT("/editworkaddress", controllers.EditWorkAddress())
	router.GET("/deleteaddresses", controllers.DeleteAddress())
	router.GET("/cartcheckout", app.BuyFromCart())
	router.GET("/instantbuy", app.InstantBuy())
	log.Fatal(router.Run(":" + port))
}