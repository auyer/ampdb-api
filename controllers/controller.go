package controllers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"regexp"

	"github.com/auyer/ampdb-api/db"
	"github.com/labstack/echo"
)

//ErrorBody structure is used to improve error reporting in a JSON response body
type ErrorBody struct {
	Reason string `json:"reason"`
}

//AmpController is used to export the API handler functions
type AmpController struct{} // THis is used to make functions callable from AmpController

func (AmpController) GetAMPFile(c echo.Context) error {
	// id := c.Param("id") // URL parameter
	// Data security checking to be insterted here
	// r, err := regexp.Compile(`\b[0-9A-Za-z]+\b`)
	file, err := ioutil.ReadFile("./amp.fta.json")
	var amps []db.AMP

	if err != nil {
		log.Print(err.Error())
		return err
	}

	err = json.Unmarshal(file, &amps)

	if err != nil {
		log.Print(err.Error())
		return err
	}
	for _, element := range amps {
		i, err := db.InsertAMP(element)
		if err != nil {
			log.Print(err.Error())
			return err
		}
		log.Printf("Inserted %d", i)
	}
	// amp, err := db.GetAMP(id)
	// if err != nil {
	// 	log.Println(err)
	// 	return c.JSON(400, ErrorBody{
	// 		Reason: err.Error(),
	// 	})
	// }
	return c.String(200, "")
}

//GetServidorMat funtion returns the "servidor" matching a given id
func (AmpController) GetAmpByID(c echo.Context) error {
	id := c.Param("id") // URL parameter
	// Data security checking to be insterted here
	r, err := regexp.Compile(`\b[0-9A-Za-z]+\b`)
	if !r.MatchString(id) {
		log.Println(err)
		return c.JSON(404, ErrorBody{
			Reason: err.Error(),
		})
	}
	amp, err := db.GetAMP(id)
	if err != nil {
		log.Println(err)
		return c.JSON(400, ErrorBody{
			Reason: err.Error(),
		})
	}
	return c.JSON(200, amp)
}

func (AmpController) GetAMPs(c echo.Context) error {
	amps, err := db.GetAMPs()
	if err != nil {
		log.Println(err)
		return c.JSON(400, ErrorBody{
			Reason: err.Error(),
		})
	}
	return c.JSON(200, amps)
}
