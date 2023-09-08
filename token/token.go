package token

import "fmt"

type TokenType int

const (
	// single character tokens
	LPAREN TokenType = iota
	RPAREN
	LBRACE
	RBRACE
	COMMA
	DOT
	MINUS
	PLUS
	SEMICOLON
	SLASH
	STAR

	// one or two character tokens
	BANG
	BANGEQUAL
	EQUAL
	EQUALEQUAL
	GREATER
	GREATEREQUAL
	LESS
	LESSEQUAL

	// literals
	IDENTIFIER
	STRING
	NUMBER

	// keywords
	AND
	CLASS
	ELSE
	FALSE
	FUN
	FOR
	IF
	NIL
	OR
	PRINT
	RETURN
	SUPER
	THIS
	TRUE
	VAR
	WHILE

	EOF
)

type Token struct {
	Type    TokenType
	Lexeme  string
	Literal interface{}
	Line    int
}

func NewToken(tokentype TokenType, lexeme string, literal interface{}, line int) *Token {
	return &Token{
		Type:    tokentype,
		Lexeme:  lexeme,
		Literal: literal,
		Line:    line,
	}
}

func (token *Token) String() string {
	return fmt.Sprintf("%v %v %v", token.Type, token.Lexeme, token.Literal)
}
