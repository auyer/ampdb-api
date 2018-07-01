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
	err := config.ReadConfig()
	if err != nil {
		fmt.Print("Error reading configuration file")
		log.Print(err.Error())
		return
	}

	//log.SetOutput(config.LogFile)
	if config.ConfigParams.Debug != "true" {
	}
	// BEGIN HTTPS
	db.Init()
	defer db.Close()
	httpsRouter := echo.New()

	httpsRouter.Use(middleware.Logger())
	httpsRouter.Use(middleware.Recover())

	controller := new(controllers.AmpController) //Controller instance

	httpsRouter.GET("/api/amp/", controller.GetAMPs)       //Simple route
	httpsRouter.GET("/api/amp/:id", controller.GetAmpByID) //Route with URL parameter
	// httpsRouter.GET("/api/amp/do/", controller.GetAMPFile) //Route with URL parameter
	httpsRouter.GET("/api/amp/id/", controller.GetAMPIDs) //Route with URL parameter

	err = httpsRouter.Start(":" + config.ConfigParams.HttpPort) // (":"+config.ConfigParams.HttpsPort, config.ConfigParams.TLSCertLocation, config.ConfigParams.TLSKeyLocation) // listen and serve on 0.0.0.0:8080
	// err = httpsRouter.StartTLS(":"+config.ConfigParams.HttpsPort, config.ConfigParams.TLSCertLocation, config.ConfigParams.TLSKeyLocation) // listen and serve on 0.0.0.0:8080
	if err != nil {
		fmt.Println(err.Error())
		log.Fatal(err)
		return
	}
}
