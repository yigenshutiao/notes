package main

import (
	_ "github.com/go-sql-driver/mysql"
	"notes/storage/util"
	"time"
	"upper.io/db.v3/mysql"
)

func initLOGO() {
	println(`
             _            
 _ __   ___ | |_ ___  ___ 
| '_ \ / _ \| __/ _ \/ __|
| | | | (_) | ||  __/\__ \
|_| |_|\___/ \__\___||___/

`)
}

func initDB() error {
	setteing, err := mysql.ParseURL("root:root@/notes")
	if err != nil {
		return err
	}
	db, err := mysql.Open(setteing)
	if err != nil {
		return err
	}

	util.DBConnector = db

	util.DBConnector.SetConnMaxLifetime(time.Minute * 3)
	util.DBConnector.SetMaxOpenConns(100)
	util.DBConnector.SetMaxIdleConns(80)

	return nil
}
