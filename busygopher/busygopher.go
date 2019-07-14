package busygopher

import "io"

type Op struct {
	write rune
	move  rune
	stop  rune
}
type Tape interface {
	io.ReadWriteSeeker
}

type Deck interface {
	Card(int) Op
}
type BusyGopher struct {
	Cards  Deck
	Halted bool
}

func (bg *BusyGopher) Step(rune c) {}
