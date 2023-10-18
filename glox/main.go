package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"p-friday/glox/ast"
	"p-friday/glox/scanner"
	"path/filepath"
)

func main() {
	ast.AstPrinterTest()
	// args := os.Args[1:]
	// if len(args) > 1 {
	// 	fmt.Println("Usage: glox [script]")
	// 	os.Exit(64)
	// } else if len(args) == 1 {
	// 	RunFile(args[0])
	// } else {
	// 	RunPrompt()
	// }
}

func RunFile(path string) {
	file, err := filepath.Abs(path)
	if err != nil {
		log.Fatal(err)
	}

	bytes, err := os.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}
	run(string(bytes))
}

func RunPrompt() {
	input := io.Reader(os.Stdin)
	reader := bufio.NewReader(input)

	for {
		fmt.Print(">> ")
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			os.Exit(0)
		} else if err != nil {
			log.Fatal(err)
		}

		if len(line) == 0 {
			break
		}

		run(string(line))
	}
}

func run(source string) {
	lexer := scanner.NewScanner(source)
	tokens := lexer.ScanTokens()

	for _, token := range tokens {
		fmt.Println(token)
	}
}
