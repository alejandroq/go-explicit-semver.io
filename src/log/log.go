package log

import (
	"flag"
	"os"

	"github.com/sirupsen/logrus"
)

const directoryName = ".semver"

var verbose = flag.Bool("v", false, "verbose logging")

// vlog verbosely logs to stdout
var vlog = logrus.New()

// elog logs events to events.log
var elog = logrus.New()

func init() {
	// parse for verbose
	flag.Parse()

	vlog.Out = os.Stdout
	elog.Out = os.Stderr
	elog.Formatter = &logrus.JSONFormatter{}

	file, err := os.OpenFile("./"+directoryName+"/events.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err == nil {
		elog.Out = file
	}
}

func isVerbose() bool {
	return *verbose
}

// Log contents
func Log(information string, fields map[string]interface{}) {
	if isVerbose() {
		vlog.WithFields(fields).Info(information)
	}
	elog.WithFields(fields).Info(information)
}
