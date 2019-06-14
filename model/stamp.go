package model

import (
	"fmt"
	"strings"
	"time"

	"github.com/go-redis/redis"
	"github.com/jmesserli/vros/config"
)

type StampModel struct {
	client *redis.Client
	card   CardModel
}

type Stamp struct {
	Card  Card
	Times []time.Time
}

func NewStampModel(config config.RedisConfig, cardModel CardModel) StampModel {
	return StampModel{
		client: redis.NewClient(&redis.Options{Addr: fmt.Sprintf("%s:%v", config.Host, config.Port), Password: "", DB: 0}),
		card:   cardModel,
	}
}

func (m StampModel) mkKey(date time.Time) string {
	formatted := date.Format("20060102")
	return fmt.Sprintf("stamps:%s", formatted)
}

func (m StampModel) GetAllForDay(date time.Time) []Stamp {
	resultMap, err := m.client.HGetAll(m.mkKey(date)).Result()
	if err != nil {
		panic(err)
	}

	var stamps = []Stamp{}
	for key, value := range resultMap {
		splitDates := strings.Split(value, ",")

		times := make([]time.Time, len(splitDates))
		for i, str := range splitDates {
			time, err := time.Parse(time.RFC3339, str)
			if err != nil {
				panic(err)
			}
			times[i] = time
		}

		stamps = append(stamps, Stamp{
			Card:  m.card.Get(key),
			Times: times,
		})
	}

	return stamps
}
