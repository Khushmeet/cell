package lexer

import (
	"cell/lexer/token"
	"testing"
)

type lexerTest struct {
	expectedType    token.TokenType
	expectedLiteral string
}

func TestLexer(t *testing.T) {
	input := "(+/*)"

	tests := []lexerTest{
		{token.LEFT_PAREN, "("},
		{token.PLUS, "+"},
		{token.DIVISION, "/"},
		{token.MULTIPLICATION, "*"},
		{token.RIGHT_PAREN, ")"},
		{token.EOF, ""},
	}

	l := NewLexer(input)

	for i, tt := range tests {
		nextToken := l.NextToken()

		if nextToken.Type != tt.expectedType {
			t.Fatalf(`Test [%d] - TokenType wrong.
					  Expected = %q, got = %q`, i, tt.expectedType, nextToken.Type)
		}

		if nextToken.Literal != tt.expectedLiteral {
			t.Fatalf(`Test [%d] - TokenLiteral wrong.
					  Expected = %q, got = %q`, i, tt.expectedLiteral, nextToken.Literal)
		}

	}
}
