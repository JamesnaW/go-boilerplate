package klogl

import (
	c "go-boilerplate/config"
	"os"

	"github.com/sirupsen/logrus"
)

var (
	Log = logrus.New()
)

func init() {
	// Exported Log
	Log.Out = os.Stdout

	// Set Log Level
	Log.SetLevel(logrus.DebugLevel)

	if c.AppConfig.Environment == "production" {
		Log.SetLevel(logrus.ErrorLevel)
	}
}
