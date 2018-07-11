package log

import (
	"os"

	log "github.com/sirupsen/logrus"
)

// Vlog verbosely logs to stdout
var Vlog = log.New()

// Elog logs events to explicit-events.log
var Elog = log.New()

func init() {
	Vlog.Out = os.Stdout
	Elog.Out = os.Stdout
	Elog.Formatter = &log.JSONFormatter{}

	file, err := os.OpenFile("./.v/events.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err == nil {

		Elog.Out = file
	}
}
