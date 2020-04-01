package token

type Type string

type Token struct {
	Type    Type
	Literal string
}

var keywords = map[string]Type {
	"fn": FUNCTION,
	"let": LET,
}

const (
	ILLEGAL = "ILLEGAL"
	EOF = "EOF"

	// Identifier + Literal
	IDENT = "IDENT" // add, foobar, x, y, ...
	INT = "INT" // 1, 2, ...

	// Operators
	ASSIGN = "="
	BANG = "!"
	PLUS = "+"
	MINUS = "-"
	ASTERISK = "*"
	SLASH = "/"

	LT = "<"
	GT = ">"

	// Delimiters
	COMMA = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	// Keywords
	FUNCTION = "FUNCTION"
	LET = "LET"
)

func LookupIdent(ident string) Type {
	if tokenType, ok := keywords[ident]; ok {
		return tokenType
	}
	return IDENT
}