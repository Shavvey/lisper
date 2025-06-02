package lexer

// main lexer type interface
type Token interface {
	toString() string
}

// list is just a collection of other tokens
type List struct {
	items []Token
}

// types of atoms we can have in our lang
type AtomType int

// NOTE: expand this later
const (
	LeftParen AtomType = iota
	RightParen
	Keyword
	Symbol
)

// atom is just a singular token (everything should resolve to this)
type Atom struct {
	content string
	AtomType
}
