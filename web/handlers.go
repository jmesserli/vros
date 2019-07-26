package web

import (
	"net/http"
	"time"

	"github.com/jmesserli/vros/config"
	"github.com/jmesserli/vros/model"

	"github.com/plimble/ace"
)

type handlers struct {
	Config     *config.Config
	StampModel *model.StampModel
	CardModel  *model.CardModel
}

type status struct {
	statusText string
	statusCode int
}

var ok = status{statusText: "ok", statusCode: http.StatusOK}
var error4 = status{statusText: "error", statusCode: http.StatusBadRequest}
var error5 = status{statusText: "error", statusCode: http.StatusInternalServerError}

func returnStatus(status status, message string, c *ace.C) {
	c.JSON(status.statusCode, map[string]string{
		"status":  status.statusText,
		"message": message,
	})
}

func userError(message string, c *ace.C) {
	returnStatus(error4, message, c)
}

func serverError(message string, c *ace.C) {
	returnStatus(error5, message, c)
}

func (h handlers) Ping(c *ace.C) {
	returnStatus(ok, "Pong", c)
}

func (h handlers) Echo(c *ace.C) {
	var json struct {
		UID string `json:"UID"`
	}
	c.ParseJSON(&json)

	c.JSON(http.StatusOK, map[string]string{
		"status": "ok",
		"uid":    json.UID,
	})
}

func (h handlers) StampCard(c *ace.C) {
	var json struct {
		Serial string `json:"serial"`
	}
	c.ParseJSON(&json)

	if len(json.Serial) < 1 {
		userError("serial must be longer than 0 chars", c)
		return
	}

	cardExists := h.CardModel.Exists(json.Serial)
	if !cardExists {
		card, err := h.CardModel.CreateWithRegisterCode(json.Serial)

		if err != nil {
			serverError(err.Error(), c)
			return
		}

		c.JSON(http.StatusOK, map[string]string{
			"status":        "ok",
			"action":        "register",
			"register_code": card.RegisterCode,
		})
		return
	}

	card := h.CardModel.Get(json.Serial)

	if len(card.RegisterCode) > 0 {
		c.JSON(http.StatusOK, map[string]string{
			"status":        "ok",
			"action":        "register",
			"register_code": card.RegisterCode,
		})
		return
	}

	stamp := h.StampModel.Get(json.Serial, time.Now())
	h.StampModel.AddTime(stamp, time.Now())

	c.JSON(http.StatusOK, map[string]string{
		"status": "ok",
		"action": "stamp",
	})
}
