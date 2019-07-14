package main

import (
	"log"

	"github.com/indiependente/busygopher/busygopher"
)

func main() {
	tape, err := busygopher.NewStringTape("0000000000")
	if err != nil {
		log.Fatalf("Error while creating tape: %v", err)
	}
	bg := &busygopher.BusyGopher{
		Tape: tape,
	}
	log.Println(bg)
}
