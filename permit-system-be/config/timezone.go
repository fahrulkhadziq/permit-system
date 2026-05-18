package config

import (
	"log"
	"time"
)

var AppLocation *time.Location

func InitTimezone() {
	loc, err := time.LoadLocation("Asia/Jakarta")

	if err != nil {
		log.Fatal(err)
	}
	AppLocation = loc

	time.Local = loc
}
