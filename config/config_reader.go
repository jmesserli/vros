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

type VerlesungConfig struct {
	Entries               []VerlesungEntry `json:"entries"`
	ReminderMinutesBefore int              `json:"reminder_minutes_before"`
}

type VerlesungEntry struct {
	Name        string    `json:"name"`
	Time        string    `json:"time"`
	Days        []Weekday `json:"days"`
	SignupStart string    `json:"signup_start"`
}

type Config struct {
	Redis     RedisConfig     `json:"redis"`
	Verlesung VerlesungConfig `json:"verlesung"`
	Twilio    TwilioConfig    `json:"twilio"`
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
