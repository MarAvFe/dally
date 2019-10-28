package dally

import mgo "gopkg.in/mgo.v2"

// MongoDAL is an implementation of DataAccessLayer for MongoDB
type MongoDAL struct {
	session *mgo.Session
	dbName  string
}

// c is a helper method to get a collection from the session
func (m *MongoDAL) c(collection string) *mgo.Collection {
	return m.session.DB(m.dbName).C(collection)
}

// Insert stores documents in mongo
func (m *MongoDAL) Insert(collection string, docs ...interface{}) error {
	e := []error{}
	for _, doc := range docs {
		e = append(e, m.c(collection).Insert(doc))
	}
	return e[0]
}

// NewMongoDAL creates a MongoDAL
func NewMongoDAL(dbURI string, dbName string) (DataAccessLayer, error) {
	session, err := mgo.Dial(dbURI)
	mongo := &MongoDAL{
		session: session,
		dbName:  dbName,
	}
	return mongo, err
}
