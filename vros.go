package main

import (
	"fmt"
	"time"

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

	stamp := sModel.Get(card.Serial, time.Now())
	sModel.AddTime(stamp, time.Now())
	stamp = sModel.Get(card.Serial, time.Now())
	stamps := sModel.GetAllForDay(time.Now())

	fmt.Println(stamp)
	fmt.Println(stamps)
}
