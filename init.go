package main

import (
	"github.com/go-redis/redis"
	_ "github.com/go-sql-driver/mysql"
	"notes/logging"
	"notes/storage/util"
	confutil "notes/util"
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

const (
	RedisPath = "./config/redis.json"
	MySQLPath = "./config/database.json"
)

type RedisConfig struct {
	Addr     string `json:"addr"`
	Password string `json:"password"`
	DB       int    `json:"db"`
}

func initCache() error {

	var config RedisConfig
	if err := confutil.LoadConfigJSON(RedisPath, &config); err != nil {
		logging.Fatal("[initDB] load config failed")
	}

	client := redis.NewClient(&redis.Options{
		Addr:     config.Addr,
		Password: config.Password,
		DB:       config.DB,
	})

	util.CacheClient = client

	return nil
}

type MySQLConfig struct {
	Database        string `json:"database"`
	Dsn             string `json:"dsn"`
	DbDriver        string `json:"dbdriver"`
	MaxOpenConn     int    `json:"maxopenconn"`
	MaxIdleConn     int    `json:"maxidleconn"`
	ConnMaxLifetime int    `json:"connmaxlifetime"`
}

func initDB() error {

	var config MySQLConfig

	if err := confutil.LoadConfigJSON(MySQLPath, &config); err != nil {
		logging.Fatal("[initDB] load config failed")
	}

	setting, err := mysql.ParseURL(config.Dsn)
	if err != nil {
		logging.Fatal("[initDB] parse config failed")
		return err
	}
	db, err := mysql.Open(setting)
	if err != nil {
		logging.Fatal("[initDB] open setting failed")
		return err
	}

	util.DBConnector = db

	util.DBConnector.SetMaxOpenConns(config.MaxOpenConn)
	util.DBConnector.SetMaxIdleConns(config.MaxIdleConn)

	return nil
}
