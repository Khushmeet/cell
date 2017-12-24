package lexer

import "cell/lexer/token"

// Lexer holds lexer's internal state while reading source code.
type Lexer struct {
	input        string
	position     int  // current position in input
	readPosition int  // current reading position in input
	ch           byte // current character
}

// NextToken returns the next token struct from the input
func (l *Lexer) NextToken() token.Token {
	var tok token.Token
	l.skipWhiteSpace()

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
	case '[':
		tok = token.NewToken(token.LEFT_BRACKET, l.ch)
	case ']':
		tok = token.NewToken(token.RIGHT_BRACKET, l.ch)
	case '"':
		l.readChar()
		for isLetter(l.ch) {
			tok.Literal = l.readIdentifier()
		}
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	default:
		if isLetter(l.ch) {
			tok.Literal = l.readIdentifier()
			tok.Type = token.CheckForKeyword(tok.Literal)
			return tok
		} else if isDigit(l.ch) {
			tok.Literal = l.readNumber()
			tok.Type = token.INT
			return tok
		}
		tok = token.NewToken(token.ILLEGAL, l.ch)
	}
	l.readChar()
	return tok
}

// NewLexer generates the pointer to lexer with the given input.
func NewLexer(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
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

func (l *Lexer) readIdentifier() string {
	position := l.position

	for isLetter(l.ch) {
		l.readChar()
	}

	return l.input[position:l.position]
}

func (l *Lexer) readNumber() string {
	position := l.position
	for isDigit(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

func (l *Lexer) skipWhiteSpace() {
	for l.ch == ' ' || l.ch == '\n' || l.ch == '\t' || l.ch == '\r' {
		l.readChar()
	}
}

func isLetter(ch byte) bool {
	return 'a' <= ch && 'z' >= ch || 'A' <= ch && 'Z' >= ch
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}
