package ast

import "github.com/Shri333/golox/scanner"

type Expr interface {
	Accept(v ExprVisitor) interface{}
}

type BinaryExpr struct {
	Left     Expr
	Operator *scanner.Token
	Right    Expr
}

func (b *BinaryExpr) Accept(v ExprVisitor) interface{} {
	return v.VisitBinaryExpr(b)
}

type GroupingExpr struct {
	Expression Expr
}

func (g *GroupingExpr) Accept(v ExprVisitor) interface{} {
	return v.VisitGroupingExpr(g)
}

type LiteralExpr struct {
	Value interface{}
}

func (l *LiteralExpr) Accept(v ExprVisitor) interface{} {
	return v.VisitLiteralExpr(l)
}

type UnaryExpr struct {
	Operator *scanner.Token
	Right    Expr
}

func (u *UnaryExpr) Accept(v ExprVisitor) interface{} {
	return v.VisitUnaryExpr(u)
}
