package busygopher

const (
	HALT = 'h'
)

// Op is an operation that the machine can perform.
type Op struct {
	Next  rune
	Write rune
	Move  rune
}

// Tape where the input
type Tape interface {
	ReadHead() rune
	WriteHead(rune)
	MoveHead(rune)
}

type Deck interface {
	Card(rune, rune) *Op
}
type BusyGopher struct {
	Cards  Deck
	Tape   Tape
	State  rune
	Halted bool
}

func (bg *BusyGopher) Step() {
	op := bg.Cards.Card(bg.State, bg.Tape.ReadHead())
	bg.Tape.WriteHead(op.Write)
	bg.Tape.MoveHead(op.Move)
	bg.State = op.Next
	if op.Next == HALT {
		bg.Halted = true
		return // stop execution
	}

}

func (bg *BusyGopher) String() string {
	return "state = " + string(bg.State)
}
