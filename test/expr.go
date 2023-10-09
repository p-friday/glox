package ast

import "p-friday/token"

type Expr interface {}
type Binary struct {
	left Expr
	operator token.Token
	right Expr
}
type Grouping struct {
	exprssion Expr
}
type Literal struct {
	value interface{}
}
type Unary struct {
	operator token.Token
	right Expr
}
