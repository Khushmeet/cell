package main

import (
	"cell/repl"
	"fmt"
	"os"
	"os/user"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Hello %s, Welcome to Cell REPL.\n", user.Username)
	repl.StartREPL(os.Stdin, os.Stdout)
}
