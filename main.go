package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/indiependente/busygopher/busygopher"
)

func main() {
	N := flag.Int("states", 2, "the number of states the machine can assume")
	flag.Parse()
	tape, err := busygopher.NewStringTape()
	if err != nil {
		log.Fatalf("Error while creating tape: %v", err)
	}
	deck := busygopher.NewDeck(*N)
	bg := &busygopher.BusyGopher{
		Tape:  tape,
		Cards: deck,
		State: 'a',
	}

	for !bg.Halted {
		fmt.Println(tape)
		bg.Step()
	}
	fmt.Println(tape)
}
