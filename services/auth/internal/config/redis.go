package config

import (
	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
)

func NewRedis(configuration *viper.Viper) *redis.Client {
	store := redis.NewClient(&redis.Options{
		Addr:     configuration.GetString("redis.host"),
		Password: configuration.GetString("redis.password"), // no password set
		DB:       configuration.GetInt("redis.db"),          // use default DB
	})
	return store
}
