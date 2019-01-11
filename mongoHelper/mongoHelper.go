// mongoHelper project mongoHelper.go
package mongoHelper

import (
	"errors"
	"log"
	//	"strconv"
	"strings"
	"time"

	"github.com/globalsign/mgo"
)

type MongoHelper struct {
	connectionString string
}

func NewMongoHelper(connString string) *MongoHelper {
	this := new(MongoHelper)
	this.connectionString = connString

	return this
}

func (this *MongoHelper) Connect() *mgo.Session {
	session, err := mgo.Dial(this.connectionString)

	if err != nil {
		return nil
	}

	session.SetMode(mgo.Monotonic, true)
	session.SetPoolLimit(10)
	//	session.SetPoolTimeout(time.Second * 60 * 15)
	//	session.SetSocketTimeout(time.Second * 60 * 15)

	//	notimeout, _ := time.ParseDuration("0s")
	//	session.SetCursorTimeout(notimeout)

	return session
}

func (this *MongoHelper) QueryAll(dbName string, collectionName string, query interface{}, result interface{}) error {
	s := this.Connect()

	if s == nil {
		return errors.New("QueryAll mongo session is nil")
	}

	defer s.Close()

	db := s.DB(dbName)

	if db == nil {
		return errors.New("QueryAll mongo database is nil")
	}

	c := db.C(collectionName)

	if c == nil {
		return errors.New("QueryAll mongo collection is nil")
	}

	q := c.Find(query)

	if q == nil {
		return errors.New("QueryAll mongo query is nil")
	}

	err := q.All(result)

	if err != nil {
		result = nil
	}

	return nil
}

//统计查询
func (this *MongoHelper) Aggregate(dbName string, collectionName string, query interface{}, result interface{}) error {
	s := this.Connect()

	if s == nil {
		return errors.New("Aggregate mongo session is nil")
	}

	defer s.Close()

	db := s.DB(dbName)

	if db == nil {
		return errors.New("Aggregate mongo database is nil")
	}

	c := db.C(collectionName)

	if c == nil {
		return errors.New("Aggregate mongo collection is nil")
	}

	pipe := c.Pipe(query)

	err := pipe.All(result)

	if err != nil {
		log.Println("mongodb Aggregate pipe all error=", err)
		result = nil

		return errors.New(strings.Join([]string{"Aggregate mongo pipe is nil .", err.Error()}, ""))
	}

	return nil
}

//统计查询
func (this *MongoHelper) AggregateCtrl(session *mgo.Session, dbName string, collectionName string, query interface{}, result interface{}) error {
	db := session.DB(dbName)

	if db == nil {
		return errors.New("Aggregate mongo database is nil")
	}

	c := db.C(collectionName)

	if c == nil {
		return errors.New("Aggregate mongo collection is nil")
	}

	pipe := c.Pipe(query)

	timeout, _ := time.ParseDuration("1h")
	pipe.SetMaxTime(timeout)
	pipe.Batch(1000)

	if pipe != nil {
		err := pipe.All(result)

		if err != nil {
			log.Println("mongodb Aggregate pipe all error=", err)
			result = nil
			return errors.New(strings.Join([]string{"Aggregate mongo pipe is nil .", err.Error()}, ""))
		}
	}

	return nil
}

//查询首文档
func (this *MongoHelper) QueryFirst(dbName string, collectionName string, query interface{}, result interface{}) error {
	s := this.Connect()

	if s == nil {
		return errors.New("QueryFirst mongo session is nil")
	}

	defer s.Close()

	db := s.DB(dbName)

	if db == nil {
		return errors.New("QueryFirst mongo database is nil")
	}

	c := db.C(collectionName)

	if c == nil {
		return errors.New("QueryFirst mongo collection is nil")
	}

	q := c.Find(query)

	if q == nil {
		return errors.New("Query mongo query is nil")
	}

	err := q.One(result)

	if err != nil {
		return errors.New("QueryFirst mongo query is nil")
	}

	return nil
}

//插入文档
func (this *MongoHelper) Insert(dbName string, collectionName string, docs ...interface{}) error {
	s := this.Connect()

	if s == nil {
		return errors.New("Insert mongo session is nil")
	}

	defer s.Close()

	db := s.DB(dbName)

	if db == nil {
		return errors.New("Insert mongo database is nil")
	}

	c := db.C(collectionName)

	if c == nil {
		return errors.New("Insert mongo collection is nil")
	}

	err := c.Insert(docs...)

	if err != nil {
		log.Println("insert mongo collection error:\r\n", err)
	}

	return err
}

//插入单文档
func (this *MongoHelper) InsertSingle(dbName string, collectionName string, data interface{}) error {
	var (
		err error
		s   *mgo.Session
	)

	s = this.Connect()

	if s == nil {
		return errors.New("InsertSingle mongo session is nil")
	}

	defer s.Close()

	db := s.DB(dbName)

	if db == nil {
		return errors.New("InsertSingle mongo database is nil")
	}

	c := db.C(collectionName)

	if c == nil {
		return errors.New("InsertSingle mongo collection is nil")
	}

	err = c.Insert(data)

	return err
}

//delete data
func (this *MongoHelper) Delete(dbName string, collectionName string, seletor interface{}) error {
	s := this.Connect()

	if s == nil {
		return errors.New("Delete mongo session is nil")
	}

	defer s.Close()

	db := s.DB(dbName)

	if db == nil {
		return errors.New("Delete mongo database is nil")
	}

	c := db.C(collectionName)

	if c == nil {
		return errors.New("Delete mongo collection is nil")
	}

	_, err := c.RemoveAll(seletor)
	return err

	//	return c.Remove(seletor)
}

//按条件修改文档
func (this *MongoHelper) Update(dbName string, collectionName string, selector interface{}, update interface{}) error {
	s := this.Connect()

	if s == nil {
		return errors.New("Update mongo session is nil")
	}

	defer s.Close()

	db := s.DB(dbName)

	if db == nil {
		return errors.New("Update mongo database is nil")
	}

	c := db.C(collectionName)

	if c == nil {
		return errors.New("Update mongo collection is nil")
	}

	_, err := c.UpdateAll(selector, update)

	return err
}

//
func (this *MongoHelper) InsertSingleControl(session *mgo.Session, dbName string, collectionName string, data interface{}) error {
	if session == nil {
		return errors.New("InsertSingleControl mongo session is nil")
	}

	db := session.DB(dbName)

	if db == nil {
		return errors.New("InsertSingleControl mongo database is nil")
	}

	c := db.C(collectionName)

	if c == nil {
		return errors.New("InsertSingleControl mongo collection is nil")
	}

	err := c.Insert(data)

	return err
}

func (this *MongoHelper) UpdateControl(session *mgo.Session, dbName string, collectionName string, selector interface{}, update interface{}) error {
	if session == nil {
		return errors.New("UpdateControl session is nil")
	}

	defer session.Close()

	db := session.DB(dbName)

	if db == nil {
		return errors.New("UpdateControl mongo database is nil")
	}

	c := db.C(collectionName)

	if c == nil {
		return errors.New("UpdateControl mongo collection is nil")
	}

	err := c.Update(selector, update)

	return err
}
