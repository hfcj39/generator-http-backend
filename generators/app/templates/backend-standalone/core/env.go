package core

import (
	"<%= displayName %>/global"
	"os"
	"strconv"
)

func Env() {

	// 通过环境变量设置 数据库 连接参数
	// 变量名： POSTGRES_DB、POSTGRES_USER、POSTGRES_PASSWORD、POSTGRES_HOST、POSTGRES_PORT
	if dbNameFromEnv := os.Getenv("POSTGRES_DB"); len(dbNameFromEnv) > 0 {
		global.CONFIG.Postgres.Dbname = dbNameFromEnv
	}
	if dbUserFromEnv := os.Getenv("POSTGRES_USER"); len(dbUserFromEnv) > 0 {
		global.CONFIG.Postgres.Username = dbUserFromEnv
	}
	if dbPasswordFromEnv := os.Getenv("POSTGRES_PASSWORD"); len(dbPasswordFromEnv) > 0 {
		global.CONFIG.Postgres.Password = dbPasswordFromEnv
	}
	if dbHostFromEnv := os.Getenv("POSTGRES_HOST"); len(dbHostFromEnv) > 0 {
		global.CONFIG.Postgres.Host = dbHostFromEnv
	}
	if dbPortFromEnv := os.Getenv("POSTGRES_PORT"); len(dbPortFromEnv) > 0 {
		global.CONFIG.Postgres.Port = dbPortFromEnv
	}

	// 通过环境变量设置 redis 连接参数
	// 变量名： REDIS_DB、REDIS_PASSWORD、REDIS_HOST
	if redisDbFromEnv := os.Getenv("REDIS_DB"); len(redisDbFromEnv) > 0 {
		db, err := strconv.Atoi(redisDbFromEnv)
		if err == nil {
			global.CONFIG.Redis.DB = db
		}
	}
	if redisPasswordFromEnv := os.Getenv("REDIS_PASSWORD"); len(redisPasswordFromEnv) > 0 {
		global.CONFIG.Redis.Password = redisPasswordFromEnv
	}
	if redisHostFromEnv := os.Getenv("REDIS_HOST"); len(redisHostFromEnv) > 0 {
		global.CONFIG.Redis.Host = redisHostFromEnv
	}

}
