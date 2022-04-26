package dao

import (
	"log"
	"time"

	"gtank/internal/conf"

	"github.com/go-redis/redis/v8"
)

var (
	Rdb         *redis.Client
	CachePrefix string
)

func InitRedis() error {
	c := conf.AppConf
	CachePrefix = c.GetString("redis.prefix")
	return ConnRedis(&redis.Options{
		Addr:        c.GetString("redis.addr"),
		DB:          c.GetInt("redis.db"),
		Password:    c.GetString("redis.pwd"),
		PoolSize:    c.GetInt("redis.pool_size"),
		MaxRetries:  c.GetInt("redis.max_reties"),
		IdleTimeout: c.GetDuration("redis.idle_timeout") * time.Millisecond,
	})
}

func ConnRedis(opt *redis.Options) error {
	Rdb = redis.NewClient(opt)
	_, err := Rdb.Ping(Rdb.Context()).Result()
	if err != nil {
		log.Printf("[dao] redis fail, err=%s", err)
		return err
	}
	return nil
}
