package main

import (
	"log"

	"github.com/indiependente/busygopher/busygopher"
)

func main() {
	tape, err := busygopher.NewStringTape()
	if err != nil {
		log.Fatalf("Error while creating tape: %v", err)
	}
	bg := &busygopher.BusyGopher{
		Tape: tape,
	}
	log.Println(bg)
}
