package main

import (
	"github.com/jmesserli/vros/config"
	"github.com/jmesserli/vros/model"
	"github.com/jmesserli/vros/scheduling"
	"github.com/jmesserli/vros/web"
)

func main() {
	config := config.ReadConfig("./config.json")
	cModel := model.NewCardModel(config.Redis)
	sModel := model.NewStampModel(config.Redis, cModel)

	scheduling.ScheduleJobs(&config)
	web.RegisterAndRun(&config, &cModel, &sModel)
}
