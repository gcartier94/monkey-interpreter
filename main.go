package main

import (
	"fmt"
	"os"
	"os/user"

	"github.com/gcartier94/monkey-interpreter/repl"
)

func main() {
    user, err := user.Current()
    if err != nil {
        panic(err)
    }
    fmt.Printf("Hell %s! This is the Monkey programming language!\n", user.Username)
    fmt.Printf("Fell free to type in commands\n")
    repl.Start(os.Stdin, os.Stdout)
}
