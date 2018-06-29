package db

import (
	"github.com/auyer/ampdb-api/config"
	r "gopkg.in/gorethink/gorethink.v4"
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

func InsertAMP(doc AMP) (int, error) {
	_, err := r.DB(config.ConfigParams.DbName).Table("amp").Insert(doc).Run(db)
	if err != nil {
		return 0, err
	}
	// result.All()
	return 1, nil
}

func GetAMP(id string) ([]AMP, error) {
	res, err := r.DB(config.ConfigParams.DbName).Table("amp").Get(id).Run(db)
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

func GetAMPs() ([]AMP, error) {
	res, err := r.DB(config.ConfigParams.DbName).Table("amp").Run(db)
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
