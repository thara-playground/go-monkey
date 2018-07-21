package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/thara/monkey/lexer"
	"github.com/thara/monkey/token"
)

const PROMPT = ">> "

func Start(in io.Reader, out io.Writer) {
	sc := bufio.NewScanner(in)

	for {
		fmt.Printf(PROMPT)
		if !sc.Scan() {
			return
		}

		line := sc.Text()
		l := lexer.New(line)

		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			fmt.Printf("%+v\n", tok)
		}
	}
}
