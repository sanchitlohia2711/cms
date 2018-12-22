package db

import (
	"time"

	"github.com/jinzhu/gorm"
	"github.com/ko/cms-db/config"
	"github.com/sanchitlohia2711/db-mysql/db"
	"github.com/sanchitlohia2711/go-extended-error/exerr"
)

var (
	conf   = config.New()
	gormDB = &gorm.DB{}
)

//Initialize : initialize the db
func init2() {
	params := &db.MySQLParamsStruct{}
	params.Username = conf.MySQL.Username
	params.Password = conf.MySQL.Password
	params.Hostname = conf.MySQL.Hostname
	params.Schema = conf.MySQL.Schema
	params.MaxOpenConns = conf.MySQL.MaxOpenConns
	gormDB, err := db.Initialize(params)
	if err != nil {
		err = exerr.NewExtendedError("SQL_CONNECTION_ERROR", err.Error())
		panic(err.Error())
	}

	//gormDB.DB().SetConnMaxLifetime(time.Minute * 5)
	gormDB.DB().SetMaxIdleConns(conf.MySQL.MaxOpenConns)
	gormDB.DB().SetMaxOpenConns(conf.MySQL.MaxOpenConns)
}

//Initialize : initialize the db
func init() {
	params := &db.MySQLParamsStruct{}
	params.Username = conf.MySQL.Username
	params.Password = conf.MySQL.Password
	params.Hostname = conf.MySQL.Hostname
	params.Schema = conf.MySQL.Schema
	params.MaxOpenConns = conf.MySQL.MaxOpenConns
	var err error
	gormDB, err = db.Initialize(params)
	if err != nil {
		err = exerr.NewExtendedError("SQL_CONNECTION_ERROR", err.Error())
		panic(err.Error())
	}

	gormDB.DB().SetConnMaxLifetime(time.Minute * 5)
	gormDB.DB().SetMaxIdleConns(conf.MySQL.MaxOpenConns)
	gormDB.DB().SetMaxOpenConns(conf.MySQL.MaxOpenConns)
}

//GormDB return gorm db
func GormDB() *gorm.DB {
	return gormDB
}
