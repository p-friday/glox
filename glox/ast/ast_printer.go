package ast

import (
	"fmt"
	"p-friday/glox/token"
	s "strings"
)

type AstPrinter struct{}

func (printer *AstPrinter) print(expr Expr) interface{} {
	return expr.accept(printer)
}

func (printer *AstPrinter) visitBinaryExpr(expr *Binary) interface{} {
	return printer.parenthesize(expr.operator.Lexeme, expr.left, expr.right)
}

func (printer *AstPrinter) visitGroupingExpr(expr *Grouping) interface{} {
	return printer.parenthesize("group", expr.exprssion)
}

func (printer *AstPrinter) visitLiteralExpr(expr *Literal) interface{} {
	if expr.value == nil {
		return "nil"
	} else {
		return expr.value
	}
}

func (printer *AstPrinter) visitUnaryExpr(expr *Unary) interface{} {
	return printer.parenthesize(expr.operator.Lexeme, expr.right)
}

func (ap *AstPrinter) parenthesize(name string, exprs ...Expr) string {
	var builder s.Builder
	builder.WriteByte('(')
	builder.WriteString(name)
	for _, expr := range exprs {
		builder.WriteByte(' ')
		builder.WriteString(fmt.Sprint(expr.accept(ap)))
	}
	builder.WriteByte(')')

	return builder.String()
}

func AstPrinterTest() {
	expression := &Binary{
		left: &Unary{
			*token.NewToken(token.MINUS, "-", nil, 1),
			&Literal{123},
		},
		operator: *token.NewToken(token.STAR, "*", nil, 1),
		right:    &Grouping{&Literal{45.67}},
	}

	var ap AstPrinter
	fmt.Println(ap.print(expression))
}
