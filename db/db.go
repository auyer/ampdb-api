package db

import (
	r "gopkg.in/rethinkdb/rethinkdb-go.v5"
)

// AMP structure is used to store data used by this API
type AMP struct {
	ID                   string `gorethink:"id"`
	Species              string `gorethink:"species"`
	Title                string `gorethink:"title"`
	Header               string `gorethink:"header"`
	HidrofobicStructures []struct {
		AvgHfobicity   string   `gorethink:"avgHfobicity"`
		Charge         string   `gorethink:"charge"`
		CrgNat         string   `gorethink:"crgNat"`
		DipNat         string   `gorethink:"dipNat"`
		DipoleMomentum string   `gorethink:"dipoleMomentum"`
		HfobicSequence string   `gorethink:"hfobicSequence"`
		OcArea         []string `gorethink:"ocArea"`
		SubID          string   `gorethink:"sub_id"`
	} `gorethink:"hidrofobicStructures"`
}

// config.ConfigParams.DbName
// GetAMP Makes the Querry for a specific AMP
func GetAMP(id string, dbName string, db *r.Session) ([]AMP, error) {
	res, err := r.DB(dbName).Table("amp").Get(id).Run(db)
	defer res.Close()
	var a []AMP
	if err != nil {
		// log.Output(err)
		return a, err
	}
	err = res.All(&a)
	if err != nil {
		return a, err
		// log.Output(err)
	}
	return a, nil
}

// GetAMPs Makes the Querry for all AMPs
func GetAMPs(dbName string, db *r.Session) ([]AMP, error) {
	res, err := r.DB(dbName).Table("amp").Run(db)
	defer res.Close()
	var a []AMP
	if err != nil {
		// log.Output(err)
		return a, err
	}
	err = res.All(&a)
	if err != nil {
		return a, err
		// log.Output(err)
	}
	return a, nil
}

// GetAMPIDs Makes the Querry for all IDs
func GetAMPIDs(dbName string, db *r.Session) ([]AMP, error) {
	res, err := r.DB(dbName).Table("amp").Pluck("id").Run(db)
	defer res.Close()
	var a []AMP
	if err != nil {
		// log.Output(err)
		return a, err
	}
	err = res.All(&a)
	if err != nil {
		return a, err
		// log.Output(err)
	}
	return a, nil
}

// InsertAMP is a sime inserting handler. This should be used only for testing.
func InsertAMP(doc AMP, dbName string, db *r.Session) (int, error) {
	_, err := r.DB(dbName).Table("amp").Insert(doc).Run(db)
	if err != nil {
		return 0, err
	}
	// result.All()
	return 1, nil
}
