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

	// operators
	EQUAL
	PLUS
	MULTIPLICATION
	MINUS
	DIVISION
	LEFT_PAREN
	RIGHT_PAREN
	LEFT_BRACKET
	RIGHT_BRACKET

	// keywords
	FUNCTION_DEF
	LET
	CONDITIONAL
	IF
	LOOP
	TRUE
	FALSE

	IDENTIFIER

	INT
	STRING // Unused
)

var keywords = map[string]TokenType{
	"fn":    FUNCTION_DEF,
	"let":   LET,
	"if":    IF,
	"cond":  CONDITIONAL,
	"loop":  LOOP,
	"true":  TRUE,
	"false": FALSE,
}

// NewToken generates token upon seeing a character from input
func NewToken(tokenType TokenType, ch byte) Token {
	return Token{Type: tokenType, Literal: string(ch)}
}

// CheckForKeyword if the identifier is a keyword, if yes then return's its TokenType
func CheckForKeyword(ch string) TokenType {
	if tokenType, ok := keywords[ch]; ok {
		return tokenType
	}
	return IDENTIFIER
}
