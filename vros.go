package main

import (
	"github.com/jmesserli/vros/config"
	"github.com/jmesserli/vros/model"
)

func main() {
	config := config.ReadConfig("./config.json")

	cModel := model.NewCardModel(config.Redis)
	sModel := model.NewStampModel(config.Redis, cModel)

	card := model.Card{
		Serial:       "te:st:te:st:te:st",
		Name:         "Std Beispiel Hans",
		Email:        "hans@beispiel.ch",
		RegisterCode: "asdf",
	}
}
