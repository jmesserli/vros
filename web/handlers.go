package web

import (
	"net/http"

	"github.com/plimble/ace"
)

func Ping(c *ace.C) {
	c.JSON(http.StatusOK, map[string]string{
		"status":  "ok",
		"message": "pong",
	})
}
