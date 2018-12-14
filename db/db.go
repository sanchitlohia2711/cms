package db

import (
	"fmt"

	"github.com/ko/cms-db/config"
	"github.com/sanchitlohia2711/db-mysql/db"
)

var (
	conf = config.New()
)

//Initialize : initialize the db
func Initialize() {
	params := &db.MySQLParamsStruct{}
	params.Username = conf.MySQL.Username
	params.Password = conf.MySQL.Password
	params.Hostname = conf.MySQL.Hostname
	params.Schema = conf.MySQL.Schema
	params.MaxOpenConns = conf.MySQL.MaxOpenConns
	err := db.Initialize(params)
	if err != nil {
		err = fmt.Errorf("dberror", err.Error())
		panic(err.Error())
	}
}
