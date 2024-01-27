package ast

import "github.com/ChrysanthemumT/gointerpreter/token"

type Node interface {
    TokenLiteral() string
}

type Statement interface {
    Node // Statement interface embeds the Node interface, any type that implements Statement must implement Node
    statementNode()
}

type Expression interface {
    Node
    expressionNode()
}

/*
Program node is the root of every ast the parser produces. 
*/
type Program struct {
    Statements []Statement
}

func (p *Program) TokenLiteral() string {
    if len(p.Statements) > 0 {
        return p.Statements[0].TokenLiteral()
    } else {
        return ""
    }
}

type LetStatement struct {
    Token token.Token
    Name *Identifier
    Value Expression
}

func (ls *LetStatement) statement() {}
func (ls *LetStatement) TokenLiteral() string { return ls.Token.Literal }

type Identifier struct {
    Token token.Token
    Value string
}

func (i *Identifier) expressionNode() {}
func (i *Identifier) TokenLiteral() string { return i.Token.Literal }

