package ast

import (
	"github.com/hababisha/mon/token"
	"bytes"
)

type Node interface{
	TokenLiteral() string //will be used for debugging and testing
	String() string
}

type Statement interface{
	Node
	statementNode()
}

type Expression interface{
	Node
	expressionNode()
}

//root node
type Program struct {
	Statements []Statement
}


func (p *Program) String() string{
	var out bytes.Buffer
	for _, s := range p.Statements{
		out.WriteString(s.String())
	}

	return out.String()
}

func (ls *LetStatement) String() string{
	var out bytes.Buffer

	out.WriteString(ls.TokenLiteral() + " ")
	out.WriteString(ls.Name.String())
	out.WriteString(" = ")
	
	if ls.Value != nil {
		out.WriteString(ls.Value.String())
	}
	out.WriteString(";")

	return out.String()
}

func (rs *ReturnStatement) String() string{
	var out bytes.Buffer
	out.WriteString(rs.TokenLiteral() + " ")
	if rs.ReturnValue != nil {
		out.WriteString(rs.ReturnValue.String())
	}
	out.WriteString(";")
	return out.String()
}

func (es *ExpressionStatement) String() string{
	if es.Expression != nil {
		return es.Expression.String()
	}

	return ""
}


func (i *Identifier) String() string{
	return i.Value
}
func (p *Program) TokenLiteral() string{
	if len(p.Statements) > 0{
		return p.Statements[0].TokenLiteral()
	}else{
		return ""
	}
}

//Let statement
type LetStatement struct {
	Token token.Token // the token.Let token
	Name *Identifier
	Value Expression
}

//return statement
type ReturnStatement struct{
	Token token.Token // the return Token
	ReturnValue Expression
}

type ExpressionStatement struct {
	Token token.Token //the first token literal
	Expression Expression
}
type Identifier struct {
	Token token.Token // the token.IDENT token
	Value string
}

func (ls *LetStatement) statementNode() {}
func (ls *LetStatement) TokenLiteral() string { return ls.Token.Literal }

func (rs *ReturnStatement) statementNode() {}
func (rs *ReturnStatement) TokenLiteral() string { return rs.Token.Literal}

func (es *ExpressionStatement) expressionNode()
func (es *ExpressionStatement) TokenLiteral() string { return es.Token.Literal}


func (i *Identifier) expressionNode() {}
func (i *Identifier) TokenLiteral() string { return i.Token.Literal}

