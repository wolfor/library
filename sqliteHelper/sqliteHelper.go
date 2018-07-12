// dba_sqlite project dba_sqlite.go
package sqliteHelper

import (
	"errors"

	"github.com/gohouse/gorose"
	_ "github.com/mattn/go-sqlite3"
)

type SqliteHelper struct {
	DbConfig     map[string]interface{}
	DbDriverName string
}

func NewSqliteHelper(dbConfig map[string]interface{}, dbDriverName string) *SqliteHelper {
	helper := new(SqliteHelper)
	helper.DbConfig = dbConfig
	helper.DbDriverName = dbDriverName

	return helper
}

func (this *SqliteHelper) connect() (*gorose.Connection, error) {
	conn, err := gorose.Open(this.DbConfig, this.DbDriverName)

	return &conn, err
}

//关闭数据库连接
func (this *SqliteHelper) Close(conn gorose.Connection) {
	conn.Close()
}

//根据sql脚本获取数据
func (this *SqliteHelper) Query(sqlContext string) ([]map[string]interface{}, error) {
	conn, err := this.connect()

	if err != nil {
		return nil, errors.New("sqlite connect instance is null.")
	}

	defer conn.Close()

	return conn.Query(sqlContext)
}

//根据表名获取数据，注意：需要调用者关闭连接
func (this *SqliteHelper) Table(tableName string) (*gorose.Database, *gorose.Connection, error) {
	conn, err := this.connect()

	if err != nil {
		return nil, nil, err //errors.New("sqlite connect instance is null.")
	}

	//defer conn.Close()

	var db *gorose.Database

	db = conn.Table(tableName)

	return db, conn, nil
}

//执行sql脚本
func (this *SqliteHelper) Execute(sqlContext string) (int64, error) {
	conn, err := this.connect()

	if err != nil {
		return 0, errors.New("sqlite connect instance is null.")
	}

	defer conn.Close()

	return conn.Execute(sqlContext)
}
