package main

import (
	"log"

	"github.com/mash/uncomment"
)

func main() {
	f := uncomment.ParseFlags()
	r, w, err := uncomment.Session(f)
	if err != nil {
		log.Fatalf("uncomment: %s", err)
	}
	defer r.Close()
	defer w.Close()

	if err := uncomment.Uncomment(r, w, f.Options()...); err != nil {
		log.Fatalf("uncomment: %s", err)
	}
}
