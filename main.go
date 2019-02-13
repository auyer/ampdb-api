package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/auyer/ampdb-api/config"
	"github.com/auyer/ampdb-api/controllers"
	"github.com/auyer/ampdb-api/db"
	"github.com/gorilla/mux"
)

func main() {
	longConfigFile := flag.String("config", "", "Use to indicate the configuration file location")
	shortConfigFile := flag.String("c", "", "Use to indicate the configuration file location")
	flag.Parse()
	var conf config.ConfigurationStruct
	var err error
	if *longConfigFile != "" || *shortConfigFile != "" {
		conf, err = config.ReadFromFile(*longConfigFile + *shortConfigFile)
		if err != nil {
			log.Print(err.Error())
			return
		}
	} else {
		conf = config.ReadFromEnv()
	}

	fmt.Println("Starting GORILLA MUX API")

	log.SetOutput(conf.LogFile)
	dbPointer, err := db.ConnectDB(conf.DBAddress)
	if err != nil {
		log.Print(err.Error())
		return
	}
	defer dbPointer.Close()
	router := mux.NewRouter() // Initializes Router

	controller := &controllers.AmpController{DB: dbPointer, DBName: conf.DBName}
	// controller := new(controllers.AmpController)
	// controller.DB = DB //Controller instance

	router.HandleFunc("/api/amp/", controller.GetAMPs).Methods("GET")       // Get all AMPS
	router.HandleFunc("/api/amp/id/", controller.GetAMPIDs).Methods("POST") // Get all the IDs
	router.HandleFunc("/api/amp/:id", controller.GetAmpByID).Methods("GET") // Get the AMP with matching ID
	// router.HandleFunc("/api/amp/do/", controllers.GetAMPFile()).Methods("GET") // Administration Route, used only to load the first data into the Database

	err = http.ListenAndServe(":"+conf.HTTPPort, router)
	if err != nil {
		fmt.Println(err.Error())
		log.Fatal(err)
		return
	}
}
