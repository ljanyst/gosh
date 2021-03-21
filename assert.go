package gosh

import (
	log "github.com/sirupsen/logrus"
)

// If error is not nil, print it and exit
func Must(err error) {
	if err != nil {
		log.Fatalf("[ERROR] %s\n", err)
	}
}
