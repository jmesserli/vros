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
	formatted := date.Format("2006-01-02")
	return fmt.Sprintf("stamps:%s", formatted)
}

func (m StampModel) GetAllForDay(date time.Time) []Stamp {
	keys, err := m.client.HKeys(m.mkKey(date)).Result()
	if err != nil {
		panic(err)
	}

	var stamps = []Stamp{}
	for _, key := range keys {
		stamps = append(stamps, m.Get(key, date))
	}

	return stamps
}

func (m StampModel) Get(cardSerial string, date time.Time) Stamp {
	timeString, err := m.client.HGet(m.mkKey(date), cardSerial).Result()
	if err == redis.Nil {
		return Stamp{
			Card:  m.card.Get(cardSerial),
			Times: []time.Time{},
		}
	} else if err != nil {
		panic(err)
	}

	timeStrings := strings.Split(timeString, ",")
	times := []time.Time{}
	for _, timeString := range timeStrings {
		parsed, _ := time.Parse(time.RFC3339, timeString)
		times = append(times, parsed)
	}

	return Stamp{
		Card:  m.card.Get(cardSerial),
		Times: times,
	}
}

func (m StampModel) save(date time.Time, stamp Stamp) {
	timeStrings := []string{}
	for _, date := range stamp.Times {
		timeStrings = append(timeStrings, date.Format(time.RFC3339))
	}

	times := strings.Join(timeStrings, ",")

	err := m.client.HSet(m.mkKey(date), stamp.Card.Serial, times).Err()
	if err != nil {
		panic(err)
	}
}

func (m StampModel) AddTime(stamp Stamp, date time.Time) {
	stamp = m.Get(stamp.Card.Serial, date)
	stamp.Times = append(stamp.Times, date)

	m.save(date, stamp)
}
