package lexer

import "fmt"

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

func AtomTypeToStr(t AtomType) string {
	switch t {
	case LeftParen:
		return "LEFT_PAREN"
	case RightParen:
		return "RIGHT_PAREN"
	case Keyword:
		return "KEYWORD"
	case Symbol:
		return "SYMBOL"
	default:
		panic(fmt.Errorf("Unknown type: %d", t))
	}
}

// atom is just a singular token (everything should resolve to this)
type Atom struct {
	content string
	AtomType
}
