package main

import (
	"log"

	"github.com/indiependente/busygopher/busygopher"
)

func main() {
	N := 1
	tape, err := busygopher.NewStringTape()
	if err != nil {
		log.Fatalf("Error while creating tape: %v", err)
	}
	deck := busygopher.NewDeck(N)
	bg := &busygopher.BusyGopher{
		Tape:  tape,
		Cards: deck,
		State: 'a',
	}

	for !bg.Halted {
		log.Println(bg)
		log.Println(tape)
		bg.Step()
	}
	log.Println(bg)
	log.Println(tape)
}
