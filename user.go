package gosh

import (
	"os/user"

	log "github.com/sirupsen/logrus"
)

// Get the username of the current user
func UserName() string {
	u, err := user.Current()
	if err != nil {
		log.Fatalf("UserName: ", err)

	}
	return u.Username
}
