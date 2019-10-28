package dally

import (
	"log"
	"time"
)

// https://medium.com/@harrygogonis/testing-go-mocking-third-party-dependancies-4ab4e1c9bd3f

// Thing ...
type Thing struct {
	Hora int64
}

// InsertTime stores Now seconds in foo collection
func InsertTime(data DataAccessLayer) Thing {
	doc := Thing{time.Now().Unix()}

	e := data.Insert("foo", doc)
	if e != nil {
		log.Println("insert failed", e.Error())
	}
	return doc
}

// DataAccessLayer defines what methods we need from the database
type DataAccessLayer interface {
	Insert(collectionName string, docs ...interface{}) error
}
