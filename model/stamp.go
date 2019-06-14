package model

import (
	"fmt"
	"time"

	"github.com/go-redis/redis"
	"github.com/jmesserli/vros/config"
)

type StampModel struct {
	client *redis.Client
}

type Stamp struct {
	Card  Card
	Times []time.Time
}

func NewStampModel(config config.RedisConfig) StampModel {
	return StampModel{
		client: redis.NewClient(&redis.Options{Addr: fmt.Sprintf("%s:%v", config.Host, config.Port), Password: "", DB: 0}),
	}
}
