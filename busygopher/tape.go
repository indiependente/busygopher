package busygopher

import (
	"github.com/pkg/errors"
)

const (
	ErrInvalid = tapeError("Invalid character")
)

type tapeError string

func (te tapeError) Error() string {
	return string(te)
}

type StringTape struct {
	head int
	data string
}

func NewStringTape(data string) (*StringTape, error) {
	err := validInput(data)
	if err != nil {
		return nil, errors.Wrap(err, "Could not instantiate StringTape")
	}
	return &StringTape{
		head: len(data) / 2,
		data: data,
	}, nil
}

func (st *StringTape) ReadHead() int {
	return int(st.data[st.head])
}

func (st *StringTape) WriteHead(c int) {
	copy := []rune(st.data)
	copy[st.head] = rune(c)
	st.data = string(copy)
}
func (st *StringTape) MoveHead(c int) {
	if c == 0 {
		st.head--
		return
	}
	if c == 1 {
		st.head++
		return
	}
}

func validInput(data string) error {
	for i, c := range data {
		if int(c)-'0' != 0 && int(c)-'0' != 1 {
			return errors.Wrapf(ErrInvalid, "Invalid character %c in position %d", c, i)
		}
	}
	return nil
}
