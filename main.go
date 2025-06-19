package main

import (
	"fmt"
	"os"
	"os/user"

	"github.com/hirokis920/interpreter_v2/repl"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Hello, %s!\n", user.Username)
	fmt.Println("Free free to type in commands")
	repl.Start(os.Stdin, os.Stdout)
}
