package dao

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/zfd81/rooster/rsql"
)

var db *rsql.DB

func Conf(dialect string, location string, port int, user string, pwd string, dbName string) (err error) {
	var driverName, dsn string
	if dialect == "mysql" {
		driverName = "mysql"
		dsn = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=true&loc=Local", user, pwd, location, port, dbName)
	}
	db, err = rsql.Open(driverName, dsn)
	return
}
