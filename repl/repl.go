package repl

import (
	"bufio"
	"fmt"
	"io"
	"github.com/ChrysanthemumT/gointerpreter/evaluator"
	"github.com/ChrysanthemumT/gointerpreter/lexer"
	"github.com/ChrysanthemumT/gointerpreter/object"
	"github.com/ChrysanthemumT/gointerpreter/parser"
)

const PROMPT = ">>"

func Start(in io.Reader, out io.Writer) {
    scanner := bufio.NewScanner(in)
    env := object.NewEnvironment()

    for{
        fmt.Fprint(out, PROMPT)
        scanned := scanner.Scan()
        if !scanned {
            return 
        }

        line := scanner.Text()
        l := lexer.New(line)
        p := parser.New(l)

        program := p.ParseProgram()

        if len(p.Errors()) != 0 {
            printParserErorrs(out, p.Errors())
            continue
        }

        evaluated := evaluator.Eval(program, env)
        if evaluated != nil {
            io.WriteString(out, evaluated.Inspect())
            io.WriteString(out, "\n")
        }
    }
}

func printParserErorrs(out io.Writer, errors []string) {
    for _, msg := range errors {
        io.WriteString(out, "\t"+msg+"\n")
    }
}

