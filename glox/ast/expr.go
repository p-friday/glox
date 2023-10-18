package ast

import "p-friday/glox/token"

type Expr interface {
	accept(Visitor) interface{}
}

type Visitor interface {
	visitBinaryExpr(*Binary) interface{}
	visitGroupingExpr(*Grouping) interface{}
	visitLiteralExpr(*Literal) interface{}
	visitUnaryExpr(*Unary) interface{}
}

type Binary struct {
	left Expr
	operator token.Token
	right Expr
}

func (binary *Binary) accept(visitor Visitor) interface{} {
	return visitor.visitBinaryExpr(binary)
}

type Grouping struct {
	exprssion Expr
}

func (grouping *Grouping) accept(visitor Visitor) interface{} {
	return visitor.visitGroupingExpr(grouping)
}

type Literal struct {
	value interface{}
}

func (literal *Literal) accept(visitor Visitor) interface{} {
	return visitor.visitLiteralExpr(literal)
}

type Unary struct {
	operator token.Token
	right Expr
}

func (unary *Unary) accept(visitor Visitor) interface{} {
	return visitor.visitUnaryExpr(unary)
}

