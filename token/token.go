package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	// 識別子 + リテラル
	IDENTIFIER = "IDENTIFIER" // add, foobar, x, y, ...
	INT        = "INT"

	// 演算子
	ASSIGN   = "="
	PLUS     = "+"
	MINUS    = "-"
	BANG     = "!"
	ASTERISK = "*"
	SLASH    = "/"

	LT       = "<"
	GT       = ">"

	EQ     = "=="
	NOT_EQ = "!="

	// デリミタ
	COMMA     = ","
	SEMICOLON = ";"

	LEFT_PAREN  = "("
	RIGHT_PAREN = ")"
	LEFT_BRACE  = "{"
	RIGHT_BRACE = "}"

	// キーワード
	FUNCTION = "FUNCTION"
	LET      = "LET"
	TRUE     = "TRUE"
	FALSE    = "FALSE"
	IF       = "IF"
	ELSE     = "ELSE"
	RETURN   = "RETURN"
)

var keywords = map[string]TokenType{
	"fn":     FUNCTION,
	"let":    LET,
	"true":   TRUE,
	"false":  FALSE,
	"if":     IF,
	"else":   ELSE,
	"return": RETURN,
}

// 渡された識別子がキーワードか確認する。
// キーワードならそのTokenType定数を返し、そうでなければtoken.IDENTIFIERを返す
func LookupIdentifier(identifier string) TokenType {
	if tok, ok := keywords[identifier]; ok {
		return tok
	}
	return IDENTIFIER
}