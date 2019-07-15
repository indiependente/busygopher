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

func (st StringTape) String() string {
	return st.data
}

func NewStringTape() (*StringTape, error) {
	data := "0"
	return &StringTape{
		head: len(data) / 2,
		data: data,
	}, nil
}

func (st *StringTape) ReadHead() rune {
	return []rune(st.data)[st.head]
}

func (st *StringTape) WriteHead(c rune) {
	copy := []rune(st.data)
	copy[st.head] = c
	st.data = string(copy)
}

func (st *StringTape) MoveHead(c rune) {
	if c == 'l' { // left
		if st.head == 0 {
			st.data = zeroes(st.data) + st.data
			st.head = len(st.data)
		}
		st.head--
		return
	}
	if c == 'r' { // right
		if st.head == len(st.data)-1 {
			st.data = st.data + zeroes(st.data)
		}
		st.head++
		return
	}
}

func zeroes(s string) string {
	copy := ""
	for i := 0; i < len(s); i++ {
		copy += "0"
	}
	return copy
}

func validInput(data string) error {
	for i, c := range data {
		if int(c) != '0' && int(c) != '1' {
			return errors.Wrapf(ErrInvalid, "Invalid character %c in position %d", c, i)
		}
	}
	return nil
}
