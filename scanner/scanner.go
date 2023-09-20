package scanner

import (
	"log"
	"p-friday/glox/error"
	"p-friday/glox/token"
	"strconv"
)

var keywords = map[string]token.TokenType{
	"and":    token.AND,
	"class":  token.CLASS,
	"else":   token.ELSE,
	"false":  token.FALSE,
	"fun":    token.FUN,
	"for":    token.FOR,
	"if":     token.IF,
	"nil":    token.NIL,
	"or":     token.OR,
	"print":  token.PRINT,
	"return": token.RETURN,
	"super":  token.SUPER,
	"this":   token.THIS,
	"true":   token.TRUE,
	"var":    token.VAR,
	"while":  token.WHILE,
}

type Scanner struct {
	source  string
	tokens  []token.Token
	start   int
	current int
	line    int
}

func NewScanner(source string) *Scanner {
	return &Scanner{source: source, line: 1}
}

func (scanner *Scanner) ScanTokens() []token.Token {
	for !scanner.isAtEnd() {
		scanner.start = scanner.current
		scanner.scanToken()
	}

	scanner.tokens = append(scanner.tokens, *token.NewToken(token.EOF, "", nil, scanner.line))
	return scanner.tokens
}

func (scanner *Scanner) scanToken() {
	c := scanner.advance()
	switch c {
	case '(':
		scanner.addToken(token.LPAREN, nil)
	case ')':
		scanner.addToken(token.RPAREN, nil)
	case '{':
		scanner.addToken(token.LBRACE, nil)
	case '}':
		scanner.addToken(token.RBRACE, nil)
	case ',':
		scanner.addToken(token.COMMA, nil)
	case '.':
		scanner.addToken(token.DOT, nil)
	case '-':
		scanner.addToken(token.MINUS, nil)
	case '+':
		scanner.addToken(token.PLUS, nil)
	case ';':
		scanner.addToken(token.SEMICOLON, nil)
	case '*':
		scanner.addToken(token.STAR, nil)
	case '!':
		if scanner.match('=') {
			scanner.addToken(token.BANGEQUAL, nil)
		} else {
			scanner.addToken(token.BANG, nil)
		}
	case '=':
		if scanner.match('=') {
			scanner.addToken(token.EQUALEQUAL, nil)
		} else {
			scanner.addToken(token.EQUAL, nil)
		}
	case '>':
		if scanner.match('=') {
			scanner.addToken(token.GREATEREQUAL, nil)
		} else {
			scanner.addToken(token.GREATER, nil)
		}
	case '<':
		if scanner.match('=') {
			scanner.addToken(token.LESSEQUAL, nil)
		} else {
			scanner.addToken(token.LESS, nil)
		}
	case '/':
		if scanner.match('/') {
			for scanner.peek() != '\n' && !scanner.isAtEnd() {
				scanner.advance()
			}
		} else {
			scanner.addToken(token.SLASH, nil)
		}
	case ' ':
	case '\r':
	case '\t':
	case '\n':
		scanner.line += 1
	case '"':
		scanner.string()
	default:
		if isDigit(c) {
			scanner.number()
		} else if isAlpha(c) {
			scanner.identifier()
		} else {
			error.Error(scanner.line, "Unexpected character.")
		}
	}
}

func (scanner *Scanner) match(expected byte) bool {
	if scanner.isAtEnd() {
		return false
	} else if scanner.source[scanner.current] != expected {
		return false
	} else {
		scanner.current += 1
		return true
	}
}

func (scanner *Scanner) peek() byte {
	if scanner.isAtEnd() {
		return 0
	} else {
		return scanner.source[scanner.current]
	}
}

func (scanner *Scanner) peekNext() byte {
	if scanner.current+1 >= len(scanner.source) {
		return 0
	} else {
		return scanner.source[scanner.current+1]
	}
}

func (scanner *Scanner) string() {
	for scanner.peek() != '"' && !scanner.isAtEnd() {
		if scanner.peek() == '\n' {
			scanner.line += 1
		}
		scanner.advance()
	}

	if scanner.isAtEnd() {
		error.Error(scanner.line, "Unterminated string.")
		return
	}

	scanner.advance()

	value := scanner.source[scanner.start+1 : scanner.current-1]
	scanner.addToken(token.STRING, value)
}

func (scanner *Scanner) number() {
	for isDigit(scanner.peek()) {
		scanner.advance()
	}

	if scanner.peek() == '.' && isDigit(scanner.peekNext()) {
		scanner.advance()

		for isDigit(scanner.peek()) {
			scanner.advance()
		}
	}
	val, err := strconv.ParseFloat(scanner.source[scanner.start:scanner.current], 32)
	if err != nil {
		log.Fatal(err)
	}
	scanner.addToken(token.NUMBER, val)
}

func (scanner *Scanner) identifier() {
	for isAlphaNumeric(scanner.peek()) {
		scanner.advance()
	}

	text := scanner.source[scanner.start:scanner.current]
	ttype, ok := keywords[text]
	if !ok {
		ttype = token.IDENTIFIER
	}

	scanner.addToken(ttype, nil)
}

func (scanner *Scanner) isAtEnd() bool {
	return scanner.current >= len(scanner.source)
}

func (scanner *Scanner) advance() byte {
	scanner.current++
	return scanner.source[scanner.current-1]
}

// not sure if i want to make a Token with nil literal or just add nil
func (scanner *Scanner) addToken(ttype token.TokenType, literal interface{}) {
	text := scanner.source[scanner.start:scanner.current]
	scanner.tokens = append(scanner.tokens, *token.NewToken(ttype, text, literal, scanner.line))
}

func isDigit(c byte) bool {
	return c >= '0' && c <= '9'
}

func isAlpha(c byte) bool {
	return (c >= 'a' && c <= 'z') ||
		(c >= 'A' && c <= 'Z') ||
		c == '_'
}

func isAlphaNumeric(c byte) bool {
	return isAlpha(c) || isDigit(c)
}
