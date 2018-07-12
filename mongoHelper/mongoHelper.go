// mongoHelper project mongoHelper.go
package mongoHelper

import (
	//	"fmt"
	"errors"

	"gopkg.in/mgo.v2"
	//	"gopkg.in/mgo.v2/bson"
)

type MongoHelper struct {
	connectString string
}

func NewMongoHelper(connString string) *MongoHelper {
	helper := new(MongoHelper)
	helper.connectString = connString

	return helper
}

func (this *MongoHelper) Connect() *mgo.Session {
	session, err := mgo.Dial(this.connectString)

	if err != nil {
		return nil
	}

	session.SetMode(mgo.Monotonic, true)

	return session
}

func (this *MongoHelper) QueryAll(dbName string, collectionName string, query interface{}, result interface{}) error {
	session := this.Connect()

	if session == nil {
		return errors.New("QueryFirst mongo session is nil")
	}

	defer session.Close()

	db := session.DB(dbName)

	if db == nil {
		return errors.New("QueryFirst mongo database is nil")
	}

	c := db.C(collectionName)

	if c == nil {
		return errors.New("QueryFirst mongo collection is nil")
	}

	q := c.Find(query)

	if q == nil {
		return errors.New("QueryFirst mongo query is nil")
	}

	err := q.All(result)

	if err != nil {
		result = nil
	}

	return nil
}

func (this *MongoHelper) QueryFirst(dbName string, collectionName string, query interface{}, result interface{}) error {
	session := this.Connect()

	if session == nil {
		return errors.New("QueryFirst mongo session is nil")
	}

	defer session.Close()

	db := session.DB(dbName)

	if db == nil {
		return errors.New("QueryFirst mongo database is nil")
	}

	c := db.C(collectionName)

	if c == nil {
		return errors.New("QueryFirst mongo collection is nil")
	}

	q := c.Find(query)

	if q == nil {
		return errors.New("QueryFirst mongo query is nil")
	}

	err := q.One(result)

	if err != nil {
		result = nil
	}

	return nil
}

//技术验证本方法不可用2018-05-18
func (this *MongoHelper) Insert(dbName string, collectionName string, docs ...interface{}) error {
	session := this.Connect()

	if session == nil {
		return errors.New("QueryFirst mongo session is nil")
	}

	defer session.Close()

	db := session.DB(dbName)

	if db == nil {
		return errors.New("Insert mongo database is nil")
	}

	c := db.C(collectionName)

	if c == nil {
		return errors.New("Insert mongo collection is nil")
	}

	err := c.Insert(docs...)

	return err
}

func (this *MongoHelper) InsertSingle(dbName string, collectionName string, data interface{}) error {
	session := this.Connect()

	if session == nil {
		return errors.New("QueryFirst mongo session is nil")
	}

	defer session.Close()

	db := session.DB(dbName)

	if db == nil {
		return errors.New("InsertSingle mongo database is nil")
	}

	c := db.C(collectionName)

	if c == nil {
		return errors.New("InsertSingle mongo collection is nil")
	}

	err := c.Insert(data)

	return err
}

func (this *MongoHelper) Update(dbName string, collectionName string, selector interface{}, update interface{}) error {
	session := this.Connect()

	if session == nil {
		return errors.New("QueryFirst mongo session is nil")
	}

	defer session.Close()

	db := session.DB(dbName)

	if db == nil {
		return errors.New("Update mongo database is nil")
	}

	c := db.C(collectionName)

	if c == nil {
		return errors.New("Update mongo collection is nil")
	}

	err := c.Update(selector, update)

	return err
}

func (this *MongoHelper) InsertSingleControl(session *mgo.Session, dbName string, collectionName string, data interface{}) error {
	if session == nil {
		return errors.New("QueryFirst mongo session is nil")
	}

	db := session.DB(dbName)

	if db == nil {
		return errors.New("InsertSingle mongo database is nil")
	}

	c := db.C(collectionName)

	if c == nil {
		return errors.New("InsertSingle mongo collection is nil")
	}

	err := c.Insert(data)

	return err
}

func (this *MongoHelper) UpdateControl(session *mgo.Session, dbName string, collectionName string, selector interface{}, update interface{}) error {
	if session == nil {
		return errors.New("mongo session is nil")
	}

	defer session.Close()

	db := session.DB(dbName)

	if db == nil {
		return errors.New("Update mongo database is nil")
	}

	c := db.C(collectionName)

	if c == nil {
		return errors.New("Update mongo collection is nil")
	}

	err := c.Update(selector, update)

	return err
}
