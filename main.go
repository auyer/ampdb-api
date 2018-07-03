package main

import (
	"fmt"
	"log"

	"github.com/auyer/ampdb-api/config"
	"github.com/auyer/ampdb-api/controllers"
	"github.com/auyer/ampdb-api/db"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	fmt.Println("Starting Echo API")
	err := config.ReadConfig() // Reads configuration File
	if err != nil {
		fmt.Print("Error reading configuration file")
		log.Print(err.Error())
		return
	}

	if config.ConfigParams.Debug != "true" {
	}
	// BEGIN HTTPS
	db.Init() // Opens database Connection
	defer db.Close()
	router := echo.New() // Initializes Router

	router.Use(middleware.Logger())
	router.Use(middleware.Recover())

	controller := new(controllers.AmpController) //Controller instance

	router.GET("/api/amp/", controller.GetAMPs)       // Get all AMPS
	router.GET("/api/amp/id/", controller.GetAMPIDs)  // Get all the IDs
	router.GET("/api/amp/:id", controller.GetAmpByID) // Get the AMP with matching ID
	// router.GET("/api/amp/do/", controller.GetAMPFile) // Administration Route, used only to load the first data into the Database

	err = router.Start(":" + config.ConfigParams.HttpPort) // (":"+config.ConfigParams.HttpsPort, config.ConfigParams.TLSCertLocation, config.ConfigParams.TLSKeyLocation) // listen and serve on 0.0.0.0:8080
	// err = router.StartTLS(":"+config.ConfigParams.HttpsPort, config.ConfigParams.TLSCertLocation, config.ConfigParams.TLSKeyLocation) // listen and serve on 0.0.0.0:8080
	if err != nil {
		fmt.Println(err.Error())
		log.Fatal(err)
		return
	}
}
