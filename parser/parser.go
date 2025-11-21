//page 46 parsing return statements
package parser

import (
	"fmt"

	"github.com/hababisha/mon/ast"
	"github.com/hababisha/mon/lexer"
	"github.com/hababisha/mon/token"
)


type Parser struct {
	l *lexer.Lexer //pointer to an instance of the lexer on which we repeatedly call nextToken()
	curToken token.Token 
	peekToken token.Token

	errors []string
}

func New(l *lexer.Lexer) *Parser{
	p := &Parser{
		l : l,
		errors: []string{},
		}
	// read 2 tokens, so curToken and peekToken are both set
	p.nextToken()
	p.nextToken()

	return p
}

func (p *Parser) Errors() []string{
	return p.errors
}

func (p *Parser) peekError(t token.TokenType){
	msg := fmt.Sprintf("Expected next token to be %s, got %s instead", t, p.peekToken.Type)
	p.errors = append(p.errors, msg)
}

func (p *Parser) nextToken(){
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

func (p *Parser) ParseProgram() *ast.Program{
	program := &ast.Program{}
	program.Statements = []ast.Statement{}
	for p.curToken.Type != token.EOF{
		stmt := p.parseStatement()
		if stmt != nil{
			program.Statements = append(program.Statements, stmt)
		}
		p.nextToken()
	}
	return program
}

func (p *Parser) parseStatement() ast.Statement{
	switch p.curToken.Type{
	case token.LET:
		return p.parseLetStatement()
	case token.RETURN:
		return p.parseReturnStatement()
	default:
		return nil
	}
}


func (p *Parser) parseLetStatement() *ast.LetStatement{
	stmt := &ast.LetStatement{Token: p.curToken}
	if !p.expectPeek(token.IDENT){
		return nil
	}

	stmt.Name = &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}
	
	if !p.expectPeek(token.ASSIGN){
		return nil
	}
	
	//todo -> skipping expressions here
	// until we encounter a semi colon
	for !p.curTokenIs(token.SEMICOLON){
		p.nextToken()
	}

	return stmt
}

func (p *Parser) parseReturnStatement() *ast.ReturnStatement{
	stmt :=  &ast.ReturnStatement{Token: p.curToken}
	p.nextToken()

	// todo -> skipping the expressions here until
	//we encounter a semi colon

	for !p.curTokenIs(token.SEMICOLON){
		p.nextToken()
	}

	return stmt
} 	

func (p *Parser) curTokenIs(t token.TokenType) bool{
	return t == p.curToken.Type
}

func (p *Parser) peekTokenIs(t token.TokenType) bool{
	return t == p.peekToken.Type
}

func (p *Parser) expectPeek(t token.TokenType) bool {
	if p.peekTokenIs(t){
		p.nextToken()
		return true
	}else{
		p.peekError(t)
		return false
	}
}

//parsing statements is relatively straight forward. process the tokens from left to right
// expect or reject the next tokens if everything fits we return an AST node
