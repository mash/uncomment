package main

import (
	"log"

	"github.com/mash/uncomment"
)

func main() {
	f, err := uncomment.NewFlags()
	if err != nil {
		log.Fatalf("uncomment: %s", err.Error())
	}
	defer f.Close()

	if err := uncomment.Uncomment(f.Reader, f.Writer); err != nil {
		log.Fatalf("uncomment: %s", err.Error())
	}
}
