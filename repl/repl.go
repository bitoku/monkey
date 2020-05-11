package repl

import (
	"bufio"
	"fmt"
	"io"
	"monkey/lexer"
	"monkey/parser"
	"os"
)

const PROMPT = ">> "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

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

		if _, err := io.WriteString(out, program.String()); err != nil {
			fmt.Print(err)
			os.Exit(1)
		}
		if _, err := io.WriteString(out, "\n"); err != nil {
			fmt.Print(err)
			os.Exit(1)
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
