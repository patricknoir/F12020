package errors

import (
	"log"
	"os"
)

func LogError(err error) bool {
	if err!=nil {
		log.Println(err)
	}
	return err != nil
}

func ExitOnError(err error, code int) {
	if err!=nil {
		log.Println(err)
		os.Exit(code)
	}
}
