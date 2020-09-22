package token

type Token struct {
	Type    TokenType
	Literal string
}

type TokenType string

const (
	// Token definitions for lexical analysis
	ILLEGAL   = "ILLEGAL"
	EOF       = "EOF"
	IDENT     = "IDENT"
	INT       = "INT"
	STRING    = "STRING"
	ASSIGN    = "="
	PLUS      = "+"
	MINUS     = "-"
	ASTERISK  = "*"
	SLASH     = "/"
	COMMA     = ","
	LPAREN    = "("
	RPAREN    = ")"
	LBRACE    = "{"
	RBRACE    = "}"
	NOT       = "!"
	LT        = "<"
	GT        = ">"
	EQ        = "=="
	NOT_EQ    = "!="
	SEMICOLON = ";"
	FUNCTION  = "FUNCTION"
	RETURN    = "RETURN"
	LET       = "LET"
	TRUE      = "TRUE"
	FALSE     = "FALSE"
	IF        = "IF"
	ELSE      = "ELSE"
)

var keywords = map[string]TokenType{
	"fn":     FUNCTION,
	"return": RETURN,
	"let":    LET,
	"true":   TRUE,
	"false":  FALSE,
	"if":     IF,
	"else":   ELSE,
}

func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}
