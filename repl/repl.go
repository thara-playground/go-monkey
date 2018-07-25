package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/thara/monkey/evaluator"
	"github.com/thara/monkey/lexer"
	"github.com/thara/monkey/parser"
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
		p := parser.New(l)

		program := p.ParseProgram()
		if len(p.Errors()) != 0 {
			printParseErrors(out, p.Errors())
			continue
		}

		evaluated := evaluator.Eval(program)
		if evaluated != nil {
			io.WriteString(out, evaluated.Inspect())
			io.WriteString(out, "\n")
		}
	}
}

func printParseErrors(out io.Writer, errors []string) {
	for _, msg := range errors {
		io.WriteString(out, "\t"+msg+"\n")
	}
}
