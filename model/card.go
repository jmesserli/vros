package model

import (
	"fmt"

	"github.com/go-redis/redis"
	"github.com/jmesserli/vros/config"
)

type Card struct {
	Serial       string
	Name         string
	Email        string
	RegisterCode string
}

type CardModel struct {
	client *redis.Client
}

func NewCardModel(config config.RedisConfig) CardModel {
	return CardModel{
		client: redis.NewClient(&redis.Options{Addr: fmt.Sprintf("%s:%v", config.Host, config.Port), Password: "", DB: 0}),
	}
}

func (m CardModel) mkKey(serial string) string {
	return fmt.Sprintf("card:%s", serial)
}

func (m CardModel) Exists(serial string) bool {
	keys, err := m.client.HKeys(m.mkKey(serial)).Result()
	if err != nil {
		panic(err)
	}

	return len(keys) > 0
}

func (m CardModel) Get(serial string) Card {
	resultMap, err := m.client.HGetAll(m.mkKey(serial)).Result()
	if err != nil {
		panic(err)
	}

	return Card{
		Serial:       serial,
		Name:         resultMap["name"],
		Email:        resultMap["email"],
		RegisterCode: resultMap["register_code"],
	}
}

func (m CardModel) New(card Card) {
	fieldsMap := make(map[string]interface{})
	fieldsMap["name"] = card.Name
	fieldsMap["email"] = card.Email
	fieldsMap["register_code"] = card.RegisterCode

	err := m.client.HMSet(m.mkKey(card.Serial), fieldsMap).Err()
	if err != nil {
		panic(err)
	}
}
