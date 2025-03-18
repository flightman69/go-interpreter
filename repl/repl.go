package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/flightman69/go-interpreter/lexer"
	"github.com/flightman69/go-interpreter/token"
)

const PROMPT = ">> "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Fprintf(out, "%s", PROMPT)
		scanned := scanner.Scan()

		if !scanned {
			return
		}

		line := scanner.Text()
		l := lexer.New(line) // creating new lexer

		tok := l.NextToken() // initializing token
		for {                // iterating through all tokens
			if tok.Type == token.EOF {
				return
			}

			fmt.Fprintf(out, "%+v\n", tok)
			tok = l.NextToken()
		}
	}
}
