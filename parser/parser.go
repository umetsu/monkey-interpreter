package parser

import (
	"monkey-interpreter/lexer"
	"monkey-interpreter/token"
	"monkey-interpreter/ast"
)

type Parser struct {
	lexer *lexer.Lexer

	currentToken token.Token
	peekToken    token.Token
}

func New(lexer *lexer.Lexer) *Parser {
	parser := &Parser{lexer: lexer}

	// 2つトークンを読み込む。currentToken, peekTokenの両方がセットされる
	parser.nextToken()
	parser.nextToken()

	return parser
}

func (parser *Parser) nextToken() {
	parser.currentToken = parser.peekToken
	parser.peekToken = parser.lexer.NextToken()
}

func (parser *Parser) ParseProgram() *ast.Program {
	return nil
}
