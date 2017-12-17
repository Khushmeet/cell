package token

// TokenType represents the type of token
// whether it is an identifier or operator or etc.
type TokenType int

// Token represents an emited token from the lexer
type Token struct {
	Type    TokenType
	Literal string
}

const (
	ILLEGAL TokenType = iota
	EOF
	IDENTIFIER
	EQUAL
	PLUS
	MULTIPLICATION
	MINUS
	DIVISION
	LEFT_PAREN
	RIGHT_PAREN
)

func NewToken(tokenType TokenType, ch byte) Token {
	return Token{Type: tokenType, Literal: string(ch)}
}
