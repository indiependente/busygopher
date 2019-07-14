package busygopher

// Op is an operation that the machine can perform.
type Op struct {
	Write int
	Move  int
	Stop  int
	Next  int
}

// Tape where the input
type Tape interface {
	ReadHead() int
	WriteHead(int)
	MoveHead(int)
}

type Deck interface {
	Card(int) [2]Op
}
type BusyGopher struct {
	Cards  Deck
	Tape   Tape
	State  int
	Halted bool
}

func (bg *BusyGopher) Step(c int) {
	op := bg.Cards.Card(bg.State)[c]
	if op.Stop == 1 {
		return // halt
	}
	bg.Tape.WriteHead(op.Write)
	bg.Tape.MoveHead(op.Move)
	bg.State = op.Next
}
