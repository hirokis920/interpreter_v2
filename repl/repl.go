package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/hirokis920/interpreter_v2/lexer"
	"github.com/hirokis920/interpreter_v2/token"
)

const PROMPT = ">> "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Println(PROMPT)
		scaned := scanner.Scan()
		if !scaned {
			return
		}
		line := scanner.Text()
		l := lexer.New(line)
		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			fmt.Printf("%+v\n", tok)
		}
	}
}
