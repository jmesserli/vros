package config

import (
	"io/ioutil"

	"gopkg.in/yaml.v3"
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
	Host string
	Port int
}

type Verlesung struct {
	Name        string
	Time        string
	Days        []Weekday
	SignupStart string
}

type Config struct {
	Redis       RedisConfig
	Verlesungen []Verlesung
}

func ReadConfig(path string) Config {
	fileContent, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}

	config := Config{}
	err = yaml.Unmarshal(fileContent, &config)
	if err != nil {
		panic(err)
	}

	return config
}
