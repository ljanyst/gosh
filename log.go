package gosh

import (
	log "github.com/sirupsen/logrus"
	prefixed "github.com/x-cray/logrus-prefixed-formatter"
)

// Set logging up, valid levels are: Trace, Debug, Info, Warn, Error, Fatal, Panic
func SetupLogging(level string) {
	log.SetFormatter(&prefixed.TextFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
		FullTimestamp:   true,
		ForceFormatting: true,
	})

	lvl := log.InfoLevel
	if level != "" {
		l, err := log.ParseLevel(level)
		if err != nil {
			log.Fatal(err)
		}
		lvl = l
	}
	log.SetLevel(lvl)
}
