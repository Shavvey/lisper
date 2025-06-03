package lexer

type Atom struct {
	content string
	AtomType
}

type AtomType int

const (
	LeftParen AtomType = iota
)

func a() Atom {
	var a Atom
	a = Atom{"(", LeftParen}
	return a
}
