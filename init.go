package main

import (
	"github.com/go-redis/redis"
	"time"

	"notes/logging"
	"notes/storage/util"

	_ "github.com/go-sql-driver/mysql"
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

func initCache() error {
	client := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "",
		DB:       0,
	})

	util.CacheClient = client

	return nil
}

func initDB() error {
	setting, err := mysql.ParseURL("root:root@/notes")
	if err != nil {
		logging.Logger.Panicln("[initDB] parse config failed")
		return err
	}
	db, err := mysql.Open(setting)
	if err != nil {
		logging.Logger.Panicln("[initDB] open setting failed")
		return err
	}

	util.DBConnector = db

	util.DBConnector.SetConnMaxLifetime(time.Minute * 3)
	util.DBConnector.SetMaxOpenConns(100)
	util.DBConnector.SetMaxIdleConns(80)

	return nil
}
