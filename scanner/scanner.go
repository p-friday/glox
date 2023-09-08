package scanner

import (
	"p-friday/glox/token"
)

type Scanner struct {
	source string
	tokens []token.Token
	start int
	current int
	line int
}

func NewScanner(source string) *Scanner {
	return &Scanner{source: source, line: 1}
}

func (scanner *Scanner) scanTokens() []token.Token {
	for !scanner.isAtEnd() {
		start := scanner.current
		scanToken()
	}

	scanner.tokens = append(scanner.tokens, *token.NewToken(token.EOF, "", nil, scanner.line))	
	return scanner.tokens
}

func (scanner *Scanner) isAtEnd() bool {
	return scanner.current >= len(scanner.source)
}
