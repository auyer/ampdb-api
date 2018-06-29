package db

import (
	"log"

	"github.com/auyer/ampdb-api/config"
	r "gopkg.in/gorethink/gorethink.v4"
)

//DB ...
// type DB struct {
// 	db *r.Session
// }

var db *r.Session

//Init ...
func Init() {

	connectDB(r.ConnectOpts{
		Address: config.ConfigParams.DbHost,
	})

}
func Close() {
	db.Close()
}

//ConnectDB ...

func connectDB(opts r.ConnectOpts) {
	var err error

	db, err = r.Connect(opts)
	if err != nil {
		log.Fatalln(err.Error())
	}
}

//GetDB ...
func GetDB() *r.Session {
	return db
}
