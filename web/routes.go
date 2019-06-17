package web

import (
	"github.com/plimble/ace"
)

func RegisterAndRun() {
	a := ace.Default()

	a.GET("/ping", Ping)
	a.Run(":80")
}
