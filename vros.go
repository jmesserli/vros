package main

import (
	"github.com/jmesserli/vros/config"
	"github.com/jmesserli/vros/model"
	"github.com/jmesserli/vros/scheduling"
	"github.com/jmesserli/vros/web"
)

func main() {
	cfg := config.ReadConfig("./config.json")
	cModel := model.NewCardModel(cfg.Redis)
	sModel := model.NewStampModel(cfg.Redis, cModel)

	scheduling.ScheduleJobs(&cfg)
	web.RegisterAndRun(&cfg, &cModel, &sModel)
}
