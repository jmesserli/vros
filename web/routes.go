package web

import (
	"github.com/plimble/ace"

	"github.com/jmesserli/vros/config"
	"github.com/jmesserli/vros/model"
)

func RegisterAndRun(config *config.Config, cModel *model.CardModel, sModel *model.StampModel) {
	handlers := handlers{
		Config:     config,
		CardModel:  cModel,
		StampModel: sModel,
	}

	a := ace.Default()

	a.GET("/ping", handlers.Ping)
	a.POST("/echo", handlers.Echo)
	a.POST("/stamp", handlers.StampCard)
	a.Run(":80")
}
