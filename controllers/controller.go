package controllers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"

	"github.com/auyer/ampdb-api/db"
	"github.com/gorilla/mux"
	r "gopkg.in/rethinkdb/rethinkdb-go.v5"
)

// ErrorBody structure is used to improve error reporting in a JSON response body
type ErrorBody struct {
	Reason string `json:"reason"`
}

// AmpController is used to export the API handler functions
type AmpController struct {
	DB     *r.Session
	DBName string
} // THis is used to make functions callable from AmpController

// GetAmpByID handler returns the AMP with the provided ID
func (ctrl AmpController) GetAmpByID(w http.ResponseWriter, request *http.Request) {
	// id := c.Param("id")
	vars := mux.Vars(request)
	id := vars["id"]                             // URL parameter
	r, err := regexp.Compile(`\b[0-9A-Za-z]+\b`) // REGEX checking
	if !r.MatchString(id) {
		log.Println(err)
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(ErrorBody{
			Reason: err.Error(),
		})
		log.Print("[MUX] " + " | 400 | " + request.Method + "  " + request.URL.Path)
		return
	}
	amp, err := db.GetAMP(id, ctrl.DBName, ctrl.DB)
	if err != nil {
		log.Println(err)
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(ErrorBody{
			Reason: err.Error(),
		})
		log.Print("[MUX] " + " | 400 | " + request.Method + "  " + request.URL.Path)
		return
	}
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(&amp)
	log.Print("[MUX] " + " | 200 | " + request.Method + "  " + request.URL.Path)
	return
}

// GetAMPs handler returns every sing AMP
func (ctrl AmpController) GetAMPs(w http.ResponseWriter, request *http.Request) {
	amps, err := db.GetAMPs(ctrl.DBName, ctrl.DB)
	if err != nil {
		log.Println(err)
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(ErrorBody{
			Reason: err.Error(),
		})
		log.Print("[MUX] " + " | 400 | " + request.Method + "  " + request.URL.Path)
		return
	}
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(&amps)
	log.Print("[MUX] " + " | 200 | " + request.Method + "  " + request.URL.Path)
	return
}

// GetAMPIDs handler return all the IDs
func (ctrl AmpController) GetAMPIDs(w http.ResponseWriter, request *http.Request) {
	amps, err := db.GetAMPs(ctrl.DBName, ctrl.DB)
	var idlist []string
	for _, value := range amps {
		idlist = append(idlist, value.ID)
	}
	if err != nil {
		log.Println(err)
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(ErrorBody{
			Reason: err.Error(),
		})
		log.Print("[MUX] " + " | 400 | " + request.Method + "  " + request.URL.Path)
		return
	}
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(&amps)
	log.Print("[MUX] " + " | 200 | " + request.Method + "  " + request.URL.Path)
	return
}

// GetAMPFile Handler is used only for the first data insertion.
func (ctrl AmpController) GetAMPFile(w http.ResponseWriter, request *http.Request) {
	file, err := ioutil.ReadFile("./amp.fta.json")
	var amps []db.AMP

	if err != nil {
		log.Print(err.Error())
		log.Println(err)
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(ErrorBody{
			Reason: err.Error(),
		})
		log.Print("[MUX] " + " | 400 | " + request.Method + "  " + request.URL.Path)
		return
	}
	err = json.Unmarshal(file, &amps)
	if err != nil {
		log.Println(err)
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(ErrorBody{
			Reason: err.Error(),
		})
		log.Print("[MUX] " + " | 400 | " + request.Method + "  " + request.URL.Path)
		return
	}
	for _, element := range amps {
		i, err := db.InsertAMP(element, ctrl.DBName, ctrl.DB)
		if err != nil {
			log.Println(err)
			w.WriteHeader(400)
			json.NewEncoder(w).Encode(ErrorBody{
				Reason: err.Error(),
			})
			log.Print("[MUX] " + " | 400 | " + request.Method + "  " + request.URL.Path)
			return
		}
		log.Printf("Inserted %d", i)
	}
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(&amps)
	log.Print("[MUX] " + " | 200 | " + request.Method + "  " + request.URL.Path)
	return
}
