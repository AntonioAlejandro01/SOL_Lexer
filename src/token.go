package lexer

import "fmt"

//TokenType  type of token
type TokenType string

// Enum for type of tokens
const (
	ASSING         TokenType = "ASSING"
	COMMA          TokenType = "COMMA"
	DIV            TokenType = "DIV"
	ELSE           TokenType = "ELSE"
	EOF            TokenType = "EOF"
	EQ             TokenType = "EQ"
	FALSE          TokenType = "FALSE"
	FUNCTION       TokenType = "FUNCTION"
	GT             TokenType = "GT"
	IDENT          TokenType = "IDENT"
	IF             TokenType = "IF"
	ILLEGAL        TokenType = "ILLEGAL"
	INT            TokenType = "INT"
	LBRACE         TokenType = "LBRACE"
	LET            TokenType = "LET"
	LPAREN         TokenType = "LPAREN"
	LT             TokenType = "LT"
	MINUS          TokenType = "MINUS"
	MULTIPLICATION TokenType = "ADD"
	NEGATION       TokenType = "NEGATION"
	NOTEQ          TokenType = "NOTEQ"
	PLUS           TokenType = "PLUS"
	RBRACE         TokenType = "RBRACE"
	RETURN         TokenType = "RETURN"
	RPAREN         TokenType = "RPAREN"
	SEMICOLON      TokenType = "SEMICOLON"
	TRUE           TokenType = "TRUE"
)

// Token - Represent a minimun unit of sol programing language
type Token struct {
	TokenType TokenType
	Literal   string
}

// ToString  - to string
func (t Token) ToString() string {
	return fmt.Sprintf("Type: %s, Literal: %s", t.TokenType, t.Literal)
}

func lookupTokenType(literal string) TokenType {
	keywords := map[string]TokenType{}

	keywords["false"] = FALSE
	keywords["fun"] = FUNCTION
	keywords["return"] = RETURN
	keywords["if"] = IF
	keywords["else"] = ELSE
	keywords["let"] = LET
	keywords["true"] = TRUE

	tt := keywords[literal]

	if tt != "" {
		return tt
	}

	return IDENT
}
