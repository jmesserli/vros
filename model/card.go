package model

import (
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/go-redis/redis"
	"github.com/jmesserli/vros/config"
)

type Card struct {
	Serial       string
	Name         string
	Mobile       string
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

func (m CardModel) getRegisterCode() (string, error) {
	createCode := func() string {
		first, second := rand.Intn(6)+1, rand.Intn(6)+1
		return fmt.Sprintf("%v%v", first, second)
	}

	existingCodes := m.getExistingRegisterCodes()
	newCode := ""

	found := false
	for count := 55; count > 0; count-- {
		newCode = createCode()

		exists := false
		for _, existingCode := range existingCodes {
			if existingCode == newCode {
				exists = true
				break
			}
		}

		if !exists {
			found = true
			break
		}
	}

	if !found {
		return "", fmt.Errorf("Could not create a new register code.")
	}

	return newCode, nil
}

func (m CardModel) CreateWithRegisterCode(serial string) (Card, error) {
	if m.Exists(serial) {
		card := m.Get(serial)
		if len(card.RegisterCode) > 0 {
			return card, nil
		}
	}

	rc, err := m.getRegisterCode()
	if err != nil {
		return Card{}, err
	}

	card := Card{
		Serial:       serial,
		RegisterCode: rc,
	}
	ttl, _ := time.ParseDuration("10m")
	m.save(card, ttl)

	return card, nil
}

func (m CardModel) Exists(serial string) bool {
	keys, err := m.client.HKeys(m.mkKey(serial)).Result()
	if err != nil {
		panic(err)
	}

	return len(keys) > 0
}

func (m CardModel) getExistingRegisterCodes() []string {
	cards := m.getAll()
	var codes []string

	for _, card := range cards {
		if len(card.RegisterCode) > 0 {
			codes = append(codes, card.RegisterCode)
		}
	}

	return codes
}

func (m CardModel) getAll() []Card {
	keys, err := m.client.Keys("card:*").Result()
	if err != nil {
		panic(err)
	}

	cards := make([]Card, 0, len(keys))
	for _, key := range keys {
		idx := strings.Index(key, ":")

		cards = append(cards, m.Get(key[idx+1:]))
	}

	return cards
}

func (m CardModel) Get(serial string) Card {
	resultMap, err := m.client.HGetAll(m.mkKey(serial)).Result()
	if err != nil {
		panic(err)
	}

	return Card{
		Serial:       serial,
		Name:         resultMap["name"],
		Mobile:       resultMap["mobile"],
		RegisterCode: resultMap["register_code"],
	}
}

func (m CardModel) Save(card Card) {
	m.save(card, -1)
}

func (m CardModel) save(card Card, ttl time.Duration) {
	fieldsMap := make(map[string]interface{})
	fieldsMap["name"] = card.Name
	fieldsMap["mobile"] = card.Mobile
	fieldsMap["register_code"] = card.RegisterCode

	key := m.mkKey(card.Serial)
	err := m.client.HMSet(key, fieldsMap).Err()
	if err != nil {
		panic(err)
	}

	if ttl == -1 {
		err = m.client.Persist(key).Err()
	} else if ttl > 0 {
		m.client.Expire(key, ttl)
	}
}
