package helpers

import (
	"log"
	"os"
)

// get current dir
func Pwd() string {
	myDir, err := os.Getwd()

	if err != nil {
		log.Fatal(err)
	}

	return myDir
}
