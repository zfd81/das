package meta

import (
	"fmt"

	"github.com/zfd81/rooster/rsql"
)

type Connection struct {
	Driver       string   `json:"driver"`
	Address      string   `json:"address"`
	Port         int      `json:"port"`
	UserName     string   `json:"userName"`
	Password     string   `json:"password"`
	DatabaseName string   `json:"db"`
	Project      *Project `json:"-"`
	db           *rsql.DB `json:"-"`
}

func (c *Connection) Connect() (err error) {
	var driverName, dsn string
	if c.db == nil {
		if c.Driver == "mysql" {
			driverName = "mysql"
			dsn = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=true&loc=Local", c.UserName, c.Password, c.Address, c.Port, c.DatabaseName)
		}
		c.db, err = rsql.Open(driverName, dsn)
	}
	err = c.db.Ping()
	return
}
