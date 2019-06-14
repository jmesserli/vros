package main

import (
	"fmt"

	"github.com/jmesserli/vros/config"
	"github.com/jmesserli/vros/model"
)

func main() {
	config := config.ReadConfig("./config.json")
	cModel := model.NewCardModel(config.Redis)
	fmt.Println(cModel)

	card := model.Card{
		Serial:       "te:st:te:st:te:st",
		Name:         "Std Beispiel Hans",
		Email:        "hans@beispiel.ch",
		RegisterCode: "asdf",
	}

	cModel.New(card)
	fmt.Println(cModel.Exists(card.Serial))
	fmt.Println(cModel.Get(card.Serial))
}
