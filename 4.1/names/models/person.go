package models

import (
	"gopkg.in/mgo.v2/bson"
)

type Person struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

func (db *DB) InsertName(data Person) error {
	dbsession := db.Copy()
	defer dbsession.Close()

	c := dbsession.DB("test").C("people")
	return c.Insert(&data)
}

func (db *DB) GetNames() ([]Person, error) {
	var results []Person

	dbsession := db.Copy()
	defer dbsession.Close()

	c := dbsession.DB("test").C("people")
	err := c.Find(bson.M{}).All(&results)

	if err != nil {
		return nil, err
	}

	return results, nil
}
