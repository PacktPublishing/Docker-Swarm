package models

import (
	"gopkg.in/mgo.v2"
	"os"
)

type Datastore interface {
	InsertName(data Person) error
	GetNames() ([]Person, error)
}

type DB struct {
	*mgo.Session
}

func NewDB() (*DB, error) {
	db := os.Getenv("DB")
	if len(db) == 0 {
		db = "localhost"
	}
	session, err := mgo.Dial(db)
	if err != nil {
		return nil, err
	}

	session.SetMode(mgo.Monotonic, true)

	return &DB{session}, nil
}
