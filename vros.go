package main

import (
	"github.com/jmesserli/vros/web"
)

func main() {
	// config := config.ReadConfig("./config.json")

	// cModel := model.NewCardModel(config.Redis)
	// sModel := model.NewStampModel(config.Redis, cModel)

	web.RegisterAndRun()
}
