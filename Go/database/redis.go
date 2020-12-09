package database

import (
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/sirupsen/logrus"
)

var RedisClient *redis.Client
var Ctx = context.Background()

func init() {
	RedisClient = redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		DB:   0, // use default DB
	})
	if err := RedisClient.Ping(Ctx).Err(); err != nil {
		logrus.Fatalf("Failed ping redis: %s", err.Error())
	}
}
