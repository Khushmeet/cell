package repl

import (
	"bufio"
	"cell/lexer"
	"cell/lexer/token"
	"fmt"
	"io"
)

const prompt = "λ>"

// StartREPL initiates the REPL
func StartREPL(input io.Reader, output io.Writer) {
	scanner := bufio.NewScanner(input)

	for {
		fmt.Printf(prompt)
		scanned := scanner.Scan()
		if !scanned {
			return
		}
		code := scanner.Text()
		l := lexer.NewLexer(code)

		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			fmt.Printf("%+v\n", tok)
		}
	}
}
