package lexer

import "monkey-interpreter/token"

type Lexer struct {
	input        string
	position     int  // 入力における現在の位置(現在の文字)
	readPosition int  // これから読み込む位置(現在の次の文字)
	char         byte // 現在検査中の文字
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	switch l.char {
	case '=':
		tok = newToken(token.ASSIGN, l.char)
	case ';':
		tok = newToken(token.SEMICOLON, l.char)
	case '(':
		tok = newToken(token.LEFT_PAREN, l.char)
	case ')':
		tok = newToken(token.RIGHT_PAREN, l.char)
	case ',':
		tok = newToken(token.COMMA, l.char)
	case '+':
		tok = newToken(token.PLUS, l.char)
	case '{':
		tok = newToken(token.LEFT_BRACE, l.char)
	case '}':
		tok = newToken(token.RIGHT_BRACE, l.char)
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	}

	l.readChar()
	return tok
}

func newToken(tokenType token.TokenType, char byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(char)}
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.char = 0 // NUL文字相当。未読み込み or EOFを表す
	} else {
		l.char = l.input[l.readPosition] // ascii文字のみサポート
	}
	l.position = l.readPosition
	l.readPosition += 1
}
