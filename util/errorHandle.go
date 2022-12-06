package util

import (
	"log"
)

func LogError(err error) string {

	log.Println("error", err.Error())
	errorText := err.Error()
	return errorText
}
