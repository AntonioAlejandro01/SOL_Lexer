package lexer

import (
	"testing"
)

func TestIllegal(t *testing.T) {
	source := "¡@¿"
	lexer := NewLexer(source)

	tokens := []*Token{}

	for i := 0; i < len(lexer.source); i++ {
		tokens = append(tokens, lexer.NextToken())
	}

	expectedTokens := []Token{
		{TokenType: ILLEGAL, Literal: "¡"},
		{TokenType: ILLEGAL, Literal: "@"},
		{TokenType: ILLEGAL, Literal: "¿"},
		{TokenType: EOF, Literal: ""},
	}

	for i := 0; i < len(tokens); i++ {
		if tokens[i].Literal != expectedTokens[i].Literal {
			t.Errorf("Expected these tokes have the same Literal value, but got %s for first and %s for second token", tokens[i].Literal, expectedTokens[i].Literal)
		} else if tokens[i].TokenType != expectedTokens[i].TokenType {
			t.Errorf("Expected these tokens was be the same TokenType , but got this [%s,%s]", tokens[i].TokenType, expectedTokens[i].TokenType)
		}
	}

}

func TestEOF(t *testing.T) {
	source := "("
	lexer := NewLexer(source)

	tokens := []*Token{}

	for i := 0; i < len(source)+1; i++ {
		tokens = append(tokens, lexer.NextToken())
	}

	expectedTokens := []Token{

		{
			TokenType: LPAREN,
			Literal:   "(",
		},
		{
			TokenType: EOF,
			Literal:   "",
		},
	}

	for i := 0; i < len(tokens); i++ {
		if tokens[i].Literal != expectedTokens[i].Literal {
			t.Errorf("Expected these tokes have the same Literal value, but got %s for first and %s for second token", tokens[i].Literal, expectedTokens[i].Literal)
		} else if tokens[i].TokenType != expectedTokens[i].TokenType {
			t.Errorf("Expected these tokens was be the same TokenType , but got this [%s,%s]", tokens[i].TokenType, expectedTokens[i].TokenType)
		}
	}
}

func TestDelimiters(t *testing.T) {
	source := "(){},;"

	lexer := NewLexer(source)

	tokens := []*Token{}

	for i := 0; i < len(source); i++ {
		tokens = append(tokens, lexer.NextToken())
	}

	expectedTokens := []Token{

		{
			TokenType: LPAREN,
			Literal:   "(",
		},
		{
			TokenType: RPAREN,
			Literal:   ")",
		},
		{
			TokenType: LBRACE,
			Literal:   "{",
		},
		{
			TokenType: RBRACE,
			Literal:   "}",
		},
		{
			TokenType: COMMA,
			Literal:   ",",
		},
		{
			TokenType: SEMICOLON,
			Literal:   ";",
		},
		{
			TokenType: EOF,
			Literal:   "",
		},
	}
	for i := 0; i < len(tokens); i++ {
		if tokens[i].Literal != expectedTokens[i].Literal {
			t.Errorf("Expected these tokes have the same Literal value, but got %s for first and %s for second token", tokens[i].Literal, expectedTokens[i].Literal)
		} else if tokens[i].TokenType != expectedTokens[i].TokenType {
			t.Errorf("Expected these tokens was be the same TokenType , but got this [%s,%s]", tokens[i].TokenType, expectedTokens[i].TokenType)
		}
	}

}

func TestImport(t *testing.T) {
	source := "import handlerText from \"handlerText.go\""
	lexer := NewLexer(source)

	tokens := []*Token{}

	for i := 0; i < 8; i++ {
		tokens = append(tokens, lexer.NextToken())
	}

	expectedTokens := []Token{
		{
			TokenType: IMPORT,
			Literal:   "import",
		},
		{
			TokenType: IDENT,
			Literal:   "handlerText",
		},
		{
			TokenType: FROM,
			Literal:   "from",
		},
		{
			TokenType: QUOTES,
			Literal:   "\"",
		},
		{
			TokenType: IDENT,
			Literal:   "handlerText",
		},
		{
			TokenType: DOT,
			Literal:   ".",
		},
		{
			TokenType: IDENT,
			Literal:   "go",
		},
		{
			TokenType: QUOTES,
			Literal:   "\"",
		},
	}
	for i := 0; i < len(tokens); i++ {
		if tokens[i].Literal != expectedTokens[i].Literal {
			t.Errorf("Expected these tokes have the same Literal value, but got %s for first and %s for second token", tokens[i].Literal, expectedTokens[i].Literal)
		} else if tokens[i].TokenType != expectedTokens[i].TokenType {
			t.Errorf("Expected these tokens was be the same TokenType , but got this [%s,%s]", tokens[i].TokenType, expectedTokens[i].TokenType)
		}
	}
}

func TestServiceDeclaration(t *testing.T) {
	source := "service : {base random:{text: { GET as handlerText :body : {}params: {size as int}headers: {Authorization as token}}}options:{}before:{* as middlewareAuth, text as middlewareAuthText}errorsHandlers : {* as handlerError, CustomError as handleErrorCustom}}"

	expectedTokens := []Token{
		{TokenType: SERVICE, Literal: "service"},
		{TokenType: COLON, Literal: ":"},
		{TokenType: LBRACE, Literal: "{"},
		{TokenType: BASE, Literal: "base"},
		{TokenType: IDENT, Literal: "random"},
		{TokenType: COLON, Literal: ":"},
		{TokenType: LBRACE, Literal: "{"},
		{TokenType: IDENT, Literal: "text"},
		{TokenType: COLON, Literal: ":"},
		{TokenType: LBRACE, Literal: "{"},
		{TokenType: GET, Literal: "GET"},
		{TokenType: AS, Literal: "as"},
		{TokenType: IDENT, Literal: "handlerText"},
		{TokenType: COLON, Literal: ":"},
		{TokenType: BODY, Literal: "body"},
		{TokenType: COLON, Literal: ":"},
		{TokenType: LBRACE, Literal: "{"},
		{TokenType: RBRACE, Literal: "}"},
		{TokenType: PARAMS, Literal: "params"},
		{TokenType: COLON, Literal: ":"},
		{TokenType: LBRACE, Literal: "{"},
		{TokenType: IDENT, Literal: "size"},
		{TokenType: AS, Literal: "as"},
		{TokenType: IDENT, Literal: "int"},
		{TokenType: RBRACE, Literal: "}"},
		{TokenType: HEADERS, Literal: "headers"},
		{TokenType: COLON, Literal: ":"},
		{TokenType: LBRACE, Literal: "{"},
		{TokenType: IDENT, Literal: "Authorization"},
		{TokenType: AS, Literal: "as"},
		{TokenType: IDENT, Literal: "token"},
		{TokenType: RBRACE, Literal: "}"},
		{TokenType: RBRACE, Literal: "}"},
		{TokenType: RBRACE, Literal: "}"},
		{TokenType: SERVICEOPTIONS, Literal: "options"},
		{TokenType: COLON, Literal: ":"},
		{TokenType: LBRACE, Literal: "{"},
		{TokenType: RBRACE, Literal: "}"},
		{TokenType: BEFORE, Literal: "before"},
		{TokenType: COLON, Literal: ":"},
		{TokenType: LBRACE, Literal: "{"},
		{TokenType: ALL, Literal: "*"},
		{TokenType: AS, Literal: "as"},
		{TokenType: IDENT, Literal: "middlewareAuth"},
		{TokenType: COMMA, Literal: ","},
		{TokenType: IDENT, Literal: "text"},
		{TokenType: AS, Literal: "as"},
		{TokenType: IDENT, Literal: "middlewareAuthText"},
		{TokenType: RBRACE, Literal: "}"},
		{TokenType: ERRORSHANDLERS, Literal: "errorsHandlers"},
		{TokenType: COLON, Literal: ":"},
		{TokenType: LBRACE, Literal: "{"},
		{TokenType: ALL, Literal: "*"},
		{TokenType: AS, Literal: "as"},
		{TokenType: IDENT, Literal: "handlerError"},
		{TokenType: COMMA, Literal: ","},
		{TokenType: IDENT, Literal: "CustomError"},
		{TokenType: AS, Literal: "as"},
		{TokenType: IDENT, Literal: "handleErrorCustom"},
		{TokenType: RBRACE, Literal: "}"},
		{TokenType: RBRACE, Literal: "}"},
		{TokenType: EOF, Literal: ""},
	}

	lexer := NewLexer(source)

	tokens := []*Token{}

	for i := 0; i < len(expectedTokens); i++ {
		tokens = append(tokens, lexer.NextToken())
	}
	for i := 0; i < len(tokens); i++ {
		if tokens[i].Literal != expectedTokens[i].Literal {
			t.Errorf("Expected these tokes have the same Literal value, but got %s for first and %s for second token", tokens[i].Literal, expectedTokens[i].Literal)
		} else if tokens[i].TokenType != expectedTokens[i].TokenType {
			t.Errorf("Expected these tokens was be the same TokenType , but got this [%s,%s]", tokens[i].TokenType, expectedTokens[i].TokenType)
		}
	}

}
