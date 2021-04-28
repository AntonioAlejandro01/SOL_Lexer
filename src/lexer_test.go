package lexer

import (
	"testing"
)

func TestIllegal(t *testing.T) {
	source := "¡@¿"
	lexer := NewLexer(source)

	tokens := []Token{}

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

func TestOneCharacterOperator(t *testing.T) {
	source := "=+-/*<>!"
	lexer := NewLexer(source)

	tokens := []Token{}

	for i := 0; i < 8; i++ {
		tokens = append(tokens, lexer.NextToken())
	}
	t.Log(tokens)
	expectedTokens := []Token{
		{TokenType: ASSING, Literal: "="},
		{TokenType: PLUS, Literal: "+"},
		{TokenType: MINUS, Literal: "-"},
		{TokenType: DIV, Literal: "/"},
		{TokenType: MULTIPLICATION, Literal: "*"},
		{TokenType: LT, Literal: "<"},
		{TokenType: GT, Literal: ">"},
		{TokenType: NEGATION, Literal: "!"},
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
	source := "+"
	lexer := NewLexer(source)

	tokens := []Token{}

	for i := 0; i < len(source)+1; i++ {
		tokens = append(tokens, lexer.NextToken())
	}

	expectedTokens := []Token{

		{
			TokenType: PLUS,
			Literal:   "+",
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

	tokens := []Token{}

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

func TestAssignment(t *testing.T) {
	source := "let five = 5;"
	lexer := NewLexer(source)

	tokens := []Token{}

	for i := 0; i < 5; i++ {
		tokens = append(tokens, lexer.NextToken())
	}

	expectedTokens := []Token{
		{
			TokenType: LET,
			Literal:   "let",
		},
		{
			TokenType: IDENT,
			Literal:   "five",
		},
		{
			TokenType: ASSING,
			Literal:   "=",
		},
		{
			TokenType: INT,
			Literal:   "5",
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

func TestFunctionDeclaration(t *testing.T) {
	source := "let sum = fun(x, y) {\n\tx + y;\n};"

	lexer := NewLexer(source)

	tokens := []Token{}
	for i := 0; i < 16; i++ {
		tokens = append(tokens, lexer.NextToken())
	}

	expectedTokens := []Token{
		{TokenType: LET, Literal: "let"},
		{TokenType: IDENT, Literal: "sum"},
		{TokenType: ASSING, Literal: "="},
		{TokenType: FUNCTION, Literal: "fun"},
		{TokenType: LPAREN, Literal: "("},
		{TokenType: IDENT, Literal: "x"},
		{TokenType: COMMA, Literal: ","},
		{TokenType: IDENT, Literal: "y"},
		{TokenType: RPAREN, Literal: ")"},
		{TokenType: LBRACE, Literal: "{"},
		{TokenType: IDENT, Literal: "x"},
		{TokenType: PLUS, Literal: "+"},
		{TokenType: IDENT, Literal: "y"},
		{TokenType: SEMICOLON, Literal: ";"},
		{TokenType: RBRACE, Literal: "}"},
		{TokenType: SEMICOLON, Literal: ";"},
		{TokenType: EOF, Literal: ""},
	}
	t.Log(tokens)
	t.Log(expectedTokens)
	for i := 0; i < len(tokens); i++ {
		if tokens[i].Literal != expectedTokens[i].Literal {
			t.Errorf("Expected these tokes have the same Literal value, but got %s for first and %s for second token", tokens[i].Literal, expectedTokens[i].Literal)
		} else if tokens[i].TokenType != expectedTokens[i].TokenType {
			t.Errorf("Expected these tokens was be the same TokenType , but got this [%s,%s]", tokens[i].TokenType, expectedTokens[i].TokenType)
		}
	}

}

func TestFunctionCall(t *testing.T) {
	source := "let result = sum(two, three);"

	lexer := NewLexer(source)

	tokens := []Token{}

	for i := 0; i < 9; i++ {
		tokens = append(tokens, lexer.NextToken())
	}

	expectedTokens := []Token{
		{TokenType: LET, Literal: "let"},
		{TokenType: IDENT, Literal: "result"},
		{TokenType: ASSING, Literal: "="},
		{TokenType: IDENT, Literal: "sum"},
		{TokenType: LPAREN, Literal: "("},
		{TokenType: IDENT, Literal: "two"},
		{TokenType: COMMA, Literal: ","},
		{TokenType: IDENT, Literal: "three"},
		{TokenType: RPAREN, Literal: ")"},
		{TokenType: SEMICOLON, Literal: ";"},
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

func TestControlStatement(t *testing.T) {
	source := "if (5 < 10) {\n\treturn  true;\n} else {\n\treturn false;\n}"
	lexer := NewLexer(source)
	tokens := []Token{}
	for i := 0; i < 17; i++ {
		tokens = append(tokens, lexer.NextToken())
	}
	expectedTokens := []Token{
		{TokenType: IF, Literal: "if"},
		{TokenType: LPAREN, Literal: "("},
		{TokenType: INT, Literal: "5"},
		{TokenType: LT, Literal: "<"},
		{TokenType: INT, Literal: "10"},
		{TokenType: RPAREN, Literal: ")"},
		{TokenType: LBRACE, Literal: "{"},
		{TokenType: RETURN, Literal: "return"},
		{TokenType: TRUE, Literal: "true"},
		{TokenType: SEMICOLON, Literal: ";"},
		{TokenType: RBRACE, Literal: "}"},
		{TokenType: ELSE, Literal: "else"},
		{TokenType: LBRACE, Literal: "{"},
		{TokenType: RETURN, Literal: "return"},
		{TokenType: FALSE, Literal: "false"},
		{TokenType: SEMICOLON, Literal: ";"},
		{TokenType: RBRACE, Literal: "}"},
	}

	for i := 0; i < len(tokens); i++ {
		if tokens[i].Literal != expectedTokens[i].Literal {
			t.Errorf("Expected these tokes have the same Literal value, but got %s for first and %s for second token", tokens[i].Literal, expectedTokens[i].Literal)
		} else if tokens[i].TokenType != expectedTokens[i].TokenType {
			t.Errorf("Expected these tokens was be the same TokenType , but got this [%s,%s]", tokens[i].TokenType, expectedTokens[i].TokenType)
		}
	}
}

func TestTwoCharacterOperator(t *testing.T) {
	source := "10 == 10;\n 10 != 9;"

	lexer := NewLexer(source)

	tokens := []Token{}

	for i := 0; i < 8; i++ {
		tokens = append(tokens, lexer.NextToken())
	}

	expectedTokens := []Token{
		{TokenType: INT, Literal: "10"},
		{TokenType: EQ, Literal: "=="},
		{TokenType: INT, Literal: "10"},
		{TokenType: SEMICOLON, Literal: ";"},
		{TokenType: INT, Literal: "10"},
		{TokenType: NOTEQ, Literal: "!="},
		{TokenType: INT, Literal: "9"},
		{TokenType: SEMICOLON, Literal: ";"},
	}
	for i := 0; i < len(tokens); i++ {
		if tokens[i].Literal != expectedTokens[i].Literal {
			t.Errorf("Expected these tokes have the same Literal value, but got %s for first and %s for second token", tokens[i].Literal, expectedTokens[i].Literal)
		} else if tokens[i].TokenType != expectedTokens[i].TokenType {
			t.Errorf("Expected these tokens was be the same TokenType , but got this [%s,%s]", tokens[i].TokenType, expectedTokens[i].TokenType)
		}
	}
}

func TestVariableLettersAndNumbers(t *testing.T) {
	source := "let sum_1 = sum_house2 + 5;"

	lexer := NewLexer(source)

	tokens := []Token{}

	for i := 0; i < 7; i++ {
		tokens = append(tokens, lexer.NextToken())
	}

	expectedTokens := []Token{
		{TokenType: LET, Literal: "let"},
		{TokenType: IDENT, Literal: "sum_1"},
		{TokenType: ASSING, Literal: "="},
		{TokenType: IDENT, Literal: "sum_house2"},
		{TokenType: PLUS, Literal: "+"},
		{TokenType: INT, Literal: "5"},
		{TokenType: SEMICOLON, Literal: ";"},
	}
	t.Log(tokens)
	for i := 0; i < len(tokens); i++ {
		if tokens[i].Literal != expectedTokens[i].Literal {
			t.Errorf("Expected these tokes have the same Literal value, but got %s for first and %s for second token", tokens[i].Literal, expectedTokens[i].Literal)
		} else if tokens[i].TokenType != expectedTokens[i].TokenType {
			t.Errorf("Expected these tokens was be the same TokenType , but got this [%s,%s]", tokens[i].TokenType, expectedTokens[i].TokenType)
		}
	}

}
