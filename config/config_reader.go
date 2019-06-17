package config

import (
	"encoding/json"
	"io/ioutil"
)

type Weekday int

const (
	Monday Weekday = iota + 1
	_
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
	Sunday
)

type RedisConfig struct {
	Host string `json:"host"`
	Port int    `json:"port"`
}

type TwilioConfig struct {
	SID    string `json:"sid"`
	Secret string `json:"secret"`
}

type Verlesung struct {
	Name        string    `json:"name"`
	Time        string    `json:"time"`
	Days        []Weekday `json:"days"`
	SignupStart string    `json:"signup_start"`
}

type Config struct {
	Redis       RedisConfig  `json:"redis"`
	Verlesungen []Verlesung  `json:"verlesungen"`
	Twilio      TwilioConfig `json:"twilio"`
}

func ReadConfig(path string) Config {
	fileContent, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}

	config := Config{}
	err = json.Unmarshal(fileContent, &config)
	if err != nil {
		panic(err)
	}

	return config
}
