package repl

import (
	"bufio"
	"fmt"
	"io"
	"monkey/evaluator"
	"monkey/lexer"
	"monkey/object"
	"monkey/parser"
	"os"
)

const PROMPT = ">> "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	env := object.NewEnvironment()

	for {
		fmt.Printf(PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		l := lexer.New(line)
		p := parser.New(l)

		program := p.ParseProgram()
		if len(p.Errors()) != 0 {
			printParserErrors(out, p.Errors())
			continue
		}

		evaluated := evaluator.Eval(program, env)

		if evaluated != nil {
			if _, err := io.WriteString(out, evaluated.Inspect()); err != nil {
				fmt.Print(err)
				os.Exit(1)
			}
			if _, err := io.WriteString(out, "\n"); err != nil {
				fmt.Print(err)
				os.Exit(1)
			}
		}
	}
}

func printParserErrors(out io.Writer, errors []string) {
	for _, msg := range errors {
		if _, err := io.WriteString(out, "  "+msg+"\n"); err != nil {
			fmt.Print(err)
			os.Exit(1)
		}
	}
}
