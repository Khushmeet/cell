package lexer

import (
	"cell/lexer/token"
	"testing"
)

type lexerTest struct {
	expectedType    token.TokenType
	expectedLiteral string
}

// TODO: Loops
func TestLexer(t *testing.T) {
	input := `
	(let x 14)

	(fn fact [x]
		(if (= x 1)
			1
			(* x (fact (- x 1)))
		)
	)

	(+ 1 2 3 4)

	(fn testCond [t]
		(cond
			((= t 1)
				(print "A")
			)
			((= t 4)
				(print "B")
			)
		)
	)

	(if true
		(+ 1 4)
		(- 1 4)
	)

	`

	tests := []lexerTest{
		{token.LEFT_PAREN, "("},
		{token.LET, "let"},
		{token.IDENTIFIER, "x"},
		{token.INT, "14"},
		{token.RIGHT_PAREN, ")"},
		{token.LEFT_PAREN, "("},
		{token.FUNCTION_DEF, "fn"},
		{token.IDENTIFIER, "fact"},
		{token.LEFT_BRACKET, "["},
		{token.IDENTIFIER, "x"},
		{token.RIGHT_BRACKET, "]"},
		{token.LEFT_PAREN, "("},
		{token.IF, "if"},
		{token.LEFT_PAREN, "("},
		{token.EQUAL, "="},
		{token.IDENTIFIER, "x"},
		{token.INT, "1"},
		{token.RIGHT_PAREN, ")"},
		{token.INT, "1"},
		{token.LEFT_PAREN, "("},
		{token.MULTIPLICATION, "*"},
		{token.IDENTIFIER, "x"},
		{token.LEFT_PAREN, "("},
		{token.IDENTIFIER, "fact"},
		{token.LEFT_PAREN, "("},
		{token.MINUS, "-"},
		{token.IDENTIFIER, "x"},
		{token.INT, "1"},
		{token.RIGHT_PAREN, ")"},
		{token.RIGHT_PAREN, ")"},
		{token.RIGHT_PAREN, ")"},
		{token.RIGHT_PAREN, ")"},
		{token.RIGHT_PAREN, ")"},
		{token.LEFT_PAREN, "("},
		{token.PLUS, "+"},
		{token.INT, "1"},
		{token.INT, "2"},
		{token.INT, "3"},
		{token.INT, "4"},
		{token.RIGHT_PAREN, ")"},
		{token.LEFT_PAREN, "("},
		{token.FUNCTION_DEF, "fn"},
		{token.IDENTIFIER, "testCond"},
		{token.LEFT_BRACKET, "["},
		{token.IDENTIFIER, "t"},
		{token.RIGHT_BRACKET, "]"},
		{token.LEFT_PAREN, "("},
		{token.CONDITIONAL, "cond"},
		{token.LEFT_PAREN, "("},
		{token.LEFT_PAREN, "("},
		{token.EQUAL, "="},
		{token.IDENTIFIER, "t"},
		{token.IDENTIFIER, "1"},
		{token.RIGHT_PAREN, ")"},
		{token.LEFT_PAREN, "("},
		{token.IDENTIFIER, "print"},
		{token.STRING, "\"A\""},
		{token.RIGHT_PAREN, ")"},
		{token.RIGHT_PAREN, ")"},
		{token.LEFT_PAREN, "("},
		{token.LEFT_PAREN, "("},
		{token.EQUAL, "="},
		{token.IDENTIFIER, "t"},
		{token.IDENTIFIER, "4"},
		{token.RIGHT_PAREN, ")"},
		{token.LEFT_PAREN, "("},
		{token.IDENTIFIER, "print"},
		{token.STRING, "\"B\""},
		{token.RIGHT_PAREN, ")"},
		{token.RIGHT_PAREN, ")"},
		{token.RIGHT_PAREN, ")"},
		{token.RIGHT_PAREN, ")"},
		{token.LEFT_PAREN, "("},
		{token.IF, "if"},
		{token.TRUE, "true"},
		{token.LEFT_PAREN, "("},
		{token.PLUS, "+"},
		{token.INT, "1"},
		{token.INT, "4"},
		{token.RIGHT_PAREN, ")"},
		{token.LEFT_PAREN, "("},
		{token.MINUS, "-"},
		{token.INT, "1"},
		{token.INT, "4"},
		{token.RIGHT_PAREN, ")"},
		{token.RIGHT_PAREN, ")"},
	}

	l := NewLexer(input)

	for i, tt := range tests {
		nextToken := l.NextToken()

		if nextToken.Type != tt.expectedType {
			t.Fatalf(`Test [%d] - TokenType wrong.
					  Expected = %v, got = %v`, i, tt.expectedType, nextToken.Type)
		}

		if nextToken.Literal != tt.expectedLiteral {
			t.Fatalf(`Test [%d] - TokenLiteral wrong.
					  Expected = %v, got = %v`, i, tt.expectedLiteral, nextToken.Literal)
		}
	}
}
