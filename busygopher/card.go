package busygopher

type Card struct {
	Ops [2]Op
}

type MapDeck map[int]Card

func (md MapDeck) Card(i int) [2]Op {
	return md[i].Ops
}
func NewDeck(n int) MapDeck {
	m := [7]map[string]string{
		{},

		{"a0": "h1r"},

		{"a0": "b1r", "a1": "b1l",
			"b0": "a1l", "b1": "h1r"},

		{"a0": "b1r", "a1": "h1r",
			"b0": "c0r", "b1": "b1r",
			"c0": "c1l", "c1": "a1l"},

		{"a0": "b1r", "a1": "b1l",
			"b0": "a1l", "b1": "c0l",
			"c0": "h1r", "c1": "d1l",
			"d0": "d1r", "d1": "a0r"},

		{"a0": "b1l", "a1": "a1l",
			"b0": "c1r", "b1": "b1r",
			"c0": "a1l", "c1": "d1r",
			"d0": "a1l", "d1": "e1r",
			"e0": "h1r", "e1": "c0r"},

		{"a0": "b1r", "a1": "e0l",
			"b0": "c1l", "b1": "a0r",
			"c0": "d1l", "c1": "c0r",
			"d0": "e1l", "d1": "f0l",
			"e0": "a1l", "e1": "c1l",
			"f0": "e1l", "f1": "h1r"},
	}

}