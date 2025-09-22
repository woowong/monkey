package main

import (
	"fmt"
	"monkey/repl"
	"os"
	"os/user"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Helelo %s! This is the Monkey programming launguage!\n", user.Username)
	fmt.Printf("Feel free type in commands\n")
	repl.Start(os.Stdin, os.Stdout)
}
