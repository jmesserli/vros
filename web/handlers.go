package web

import (
	"net/http"

	"github.com/jmesserli/vros/config"
	"github.com/jmesserli/vros/model"

	"github.com/plimble/ace"
)

type handlers struct {
	Config     *config.Config
	StampModel *model.StampModel
	CardModel  *model.CardModel
}

func (h handlers) Ping(c *ace.C) {
	c.JSON(http.StatusOK, map[string]string{
		"status":  "ok",
		"message": "pong",
	})
}
