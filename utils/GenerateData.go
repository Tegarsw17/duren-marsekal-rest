package utils

import (
	"log"
	"time"
)

func GenerateDate(value string) *time.Time {

	parsedTime, err := time.Parse(time.RFC3339, value)
	if err != nil {
		log.Fatal(err)
	}
	return &parsedTime
}
