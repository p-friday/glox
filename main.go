package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func main() {
	args := os.Args[1:]
	if len(args) > 1 {
		fmt.Println("Usage: glox [script]")
		os.Exit(64)
	} else if len(args) == 1 {
		RunFile(args[0])
	} else {
		RunPrompt()
	}
}

func RunFile(path string) {
	file, err := filepath.Abs(path)
	if err != nil {
		panic(err)
	}

	bytes, err := os.ReadFile(file)
	if err != nil {
		panic(err)
	}
	run(string(bytes))
}

func RunPrompt() {
	input := io.Reader(os.Stdin)
	reader := bufio.NewReader(input)

	for {
		fmt.Print("> ")
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			os.Exit(0)
		} else if err != nil {
			panic(err)
		}

		if len(line) == 0 {
			break
		}

		run(string(line))
	}
}

func run(source string) {
	scanner := NewScanner() //lexer
	tokens := scanner.ScanTokens()

	for i, token := range tokens {
		fmt.Println(token)
	}
}

func error(line int, message string) {
	report(line, "", message)
}

func report(line int, where string, message string) {
	fmt.Println("[line " + string(line) + "] Error" + where + ": " + message)
	//hadError := true
}
