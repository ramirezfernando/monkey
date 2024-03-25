package token

const (
	ILLEGAL = "ILLEGAL" // signifies a token/character we don't recognize
	EOF     = "EOF" // "end of file" which tells our parser that it can stop

	IDENT = "IDENT" // variable/function names
	INT   = "INT"

	// Operators
	ASSIGN = "="
	PLUS   = "+"

	// Delimiters
	COMMA     = ","
	SEMICOLON = ";"
	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	// Keywords
	FUNCTION = "FUNCTION"
	LET      = "LET"
)

type TokenType string

type Token struct {
	Type    TokenType
	Literal string // literal value of the token
}
