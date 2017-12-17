package lexer

import "cell/lexer/token"

// Lexer holds lexer's internal state while reading source code.
type Lexer struct {
	input        string
	position     int  // current position in input
	readPosition int  // current reading position in input
	ch           byte // current character
}

// NewLexer generates the pointer to lexer with the given input.
func NewLexer(input string) *Lexer {
	return &Lexer{input: input}
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition++
}

// NextToken returns the next token struct from the input
func (l *Lexer) NextToken() token.Token {
	var tok token.Token
	l.readChar()

	switch l.ch {
	case '=':
		tok = token.NewToken(token.EQUAL, l.ch)
	case '+':
		tok = token.NewToken(token.PLUS, l.ch)
	case '-':
		tok = token.NewToken(token.MINUS, l.ch)
	case '*':
		tok = token.NewToken(token.MULTIPLICATION, l.ch)
	case '/':
		tok = token.NewToken(token.DIVISION, l.ch)
	case '(':
		tok = token.NewToken(token.LEFT_PAREN, l.ch)
	case ')':
		tok = token.NewToken(token.RIGHT_PAREN, l.ch)
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	}
	return tok
}
