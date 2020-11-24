package util

import (
	"github.com/go-redis/redis"
	"upper.io/db.v3/lib/sqlbuilder"
)

var (
	DBConnector sqlbuilder.Database
	CacheClient *redis.Client
)
