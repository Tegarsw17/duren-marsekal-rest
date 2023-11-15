package utils

import "log"

func ErrorNotNill(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
