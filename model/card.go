package model

import (
	"fmt"
	"strconv"

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

const RegisterCodeCardSerial = "RE:GI:ST:ER:CO:DE"

func NewCardModel(config config.RedisConfig) CardModel {
	return CardModel{
		client: redis.NewClient(&redis.Options{Addr: fmt.Sprintf("%s:%v", config.Host, config.Port), Password: "", DB: 0}),
	}
}

func (m CardModel) mkKey(serial string) string {
	return fmt.Sprintf("card:%s", serial)
}

func (m CardModel) getRegisterCode() string {
	if !m.Exists(RegisterCodeCardSerial) {
		m.Save(Card{
			Serial:       RegisterCodeCardSerial,
			RegisterCode: "1",
		})
		return "1"
	}

	card := m.Get(RegisterCodeCardSerial)
	rc, err := strconv.Atoi(card.RegisterCode)
	if err != nil {
		panic(err)
	}

	rc++
	if rc >= 100 {
		rc = 1
	}
	if rc%10 == 0 {
		rc++
	}
	card.RegisterCode = strconv.Itoa(rc)
	m.Save(card)

	return card.RegisterCode
}

func (m CardModel) CreateWithRegisterCode(serial string) Card {
	if m.Exists(serial) {
		card := m.Get(serial)
		if len(card.RegisterCode) > 0 {
			return card
		}
	}

	rc := m.getRegisterCode()
	card := Card{
		Serial:       serial,
		RegisterCode: rc,
	}
	m.Save(card)

	return card
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

func (m CardModel) Save(card Card) {
	fieldsMap := make(map[string]interface{})
	fieldsMap["name"] = card.Name
	fieldsMap["email"] = card.Email
	fieldsMap["register_code"] = card.RegisterCode

	err := m.client.HMSet(m.mkKey(card.Serial), fieldsMap).Err()
	if err != nil {
		panic(err)
	}
}
