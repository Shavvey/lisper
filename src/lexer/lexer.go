package lexer

import "io"

type Lexer struct {
	file  io.Reader
	pos   int
	bytes int
}
