package lexer

import "fmt"

//TokenType  type of token
type TokenType string

// Enum for type of tokens
const (
	ALL            TokenType = "ALL"
	AS             TokenType = "AS"
	BASE           TokenType = "BASE"
	BEFORE         TokenType = "BEFORE"
	BODY           TokenType = "BODY"
	COMMA          TokenType = "COMMA"
	COLON          TokenType = "COLON"
	DELETE         TokenType = "DELETE"
	DOT            TokenType = "DOT"
	EOF            TokenType = "EOF"
	ERRORSHANDLERS TokenType = "ERRORSHANDLERS"
	FROM           TokenType = "FROM"
	GET            TokenType = "GET"
	HANDLER        TokenType = "HANDLER"
	HEAD           TokenType = "HEAD"
	HEADERS        TokenType = "HEADERS"
	IDENT          TokenType = "IDENT"
	IMPORT         TokenType = "IMPORT"
	ILLEGAL        TokenType = "ILLEGAL"
	LBRACE         TokenType = "LBRACE"
	LPAREN         TokenType = "LPAREN"
	OPTIONS        TokenType = "OPTIONS"
	PARAMS         TokenType = "PARAMS"
	PATCH          TokenType = "PATCH"
	POST           TokenType = "POST"
	PUT            TokenType = "PUT"
	QUOTES         TokenType = "QUOTES"
	RBRACE         TokenType = "RBRACE"
	RPAREN         TokenType = "RPAREN"
	SERVICE        TokenType = "SERVICE"
	SERVICEOPTIONS TokenType = "SERVICEOPTIONS"
	SEMICOLON      TokenType = "SEMICOLON"
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

	keywords["as"] = AS
	keywords["before"] = BEFORE
	keywords["body"] = BODY
	keywords["DELETE"] = DELETE
	keywords["errorsHandlers"] = ERRORSHANDLERS
	keywords["from"] = FROM
	keywords["GET"] = GET
	keywords["handler"] = HANDLER
	keywords["HEAD"] = HEAD
	keywords["headers"] = HEADERS
	keywords["import"] = IMPORT
	keywords["OPTIONS"] = OPTIONS
	keywords["options"] = SERVICEOPTIONS
	keywords["params"] = PARAMS
	keywords["PATCH"] = PATCH
	keywords["POST"] = POST
	keywords["PUT"] = PUT
	keywords["service"] = SERVICE
	keywords["base"] = BASE

	tt := keywords[literal]

	if tt != "" {
		return tt
	}

	return IDENT
}
