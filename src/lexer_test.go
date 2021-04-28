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
			t.Errorf("Expected these tokes have the same Literal value, but got %s for first and %s for sencond token", tokens[i].Literal, expectedTokens[i].Literal)
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
			t.Errorf("Expected these tokes have the same Literal value, but got %s for first and %s for sencond token", tokens[i].Literal, expectedTokens[i].Literal)
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
			t.Errorf("Expected these tokes have the same Literal value, but got %s for first and %s for sencond token", tokens[i].Literal, expectedTokens[i].Literal)
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
			t.Errorf("Expected these tokes have the same Literal value, but got %s for first and %s for sencond token", tokens[i].Literal, expectedTokens[i].Literal)
		} else if tokens[i].TokenType != expectedTokens[i].TokenType {
			t.Errorf("Expected these tokens was be the same TokenType , but got this [%s,%s]", tokens[i].TokenType, expectedTokens[i].TokenType)
		}
	}

}

func TestAssigment(t *testing.T) {
	source := "variable cinco = 5;"
	lexer := NewLexer(source)

	tokens := []Token{}

	for i := 0; i < 5; i++ {
		tokens = append(tokens, lexer.NextToken())
	}

	expectedTokens := []Token{
		{
			TokenType: LET,
			Literal:   "variable",
		},
		{
			TokenType: IDENT,
			Literal:   "cinco",
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
			t.Errorf("Expected these tokes have the same Literal value, but got %s for first and %s for sencond token", tokens[i].Literal, expectedTokens[i].Literal)
		} else if tokens[i].TokenType != expectedTokens[i].TokenType {
			t.Errorf("Expected these tokens was be the same TokenType , but got this [%s,%s]", tokens[i].TokenType, expectedTokens[i].TokenType)
		}
	}
}

func TestFunctionDeclaration(t *testing.T) {
	source := "variable suma = procedimiento(x, y) {\n\tx + y;\n};"

	lexer := NewLexer(source)

	tokens := []Token{}
	for i := 0; i < 16; i++ {
		tokens = append(tokens, lexer.NextToken())
	}

	expectedTokens := []Token{
		{TokenType: LET, Literal: "variable"},
		{TokenType: IDENT, Literal: "suma"},
		{TokenType: ASSING, Literal: "="},
		{TokenType: FUNCTION, Literal: "procedimiento"},
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
			t.Errorf("Expected these tokes have the same Literal value, but got %s for first and %s for sencond token", tokens[i].Literal, expectedTokens[i].Literal)
		} else if tokens[i].TokenType != expectedTokens[i].TokenType {
			t.Errorf("Expected these tokens was be the same TokenType , but got this [%s,%s]", tokens[i].TokenType, expectedTokens[i].TokenType)
		}
	}

}

func TestFunctionCall(t *testing.T) {
	//t.Error("Not Implemented yet")
	source := "variable resultado = suma(dos, tres);"

	lexer := NewLexer(source)

	tokens := []Token{}

	for i := 0; i < 9; i++ {
		tokens = append(tokens, lexer.NextToken())
	}

	expectedTokens := []Token{
		{TokenType: LET, Literal: "variable"},
		{TokenType: IDENT, Literal: "resultado"},
		{TokenType: ASSING, Literal: "="},
		{TokenType: IDENT, Literal: "suma"},
		{TokenType: LPAREN, Literal: "("},
		{TokenType: IDENT, Literal: "dos"},
		{TokenType: COMMA, Literal: ","},
		{TokenType: IDENT, Literal: "tres"},
		{TokenType: RPAREN, Literal: ")"},
		{TokenType: SEMICOLON, Literal: ";"},
		{TokenType: EOF, Literal: ""},
	}

	for i := 0; i < len(tokens); i++ {
		if tokens[i].Literal != expectedTokens[i].Literal {
			t.Errorf("Expected these tokes have the same Literal value, but got %s for first and %s for sencond token", tokens[i].Literal, expectedTokens[i].Literal)
		} else if tokens[i].TokenType != expectedTokens[i].TokenType {
			t.Errorf("Expected these tokens was be the same TokenType , but got this [%s,%s]", tokens[i].TokenType, expectedTokens[i].TokenType)
		}
	}

}

func TestControlStatement(t *testing.T) {
	source := "si (5 < 10) {\n\tregresa verdadero;\n} si_no {\n\tregresa falso;\n}"
	lexer := NewLexer(source)
	tokens := []Token{}
	for i := 0; i < 17; i++ {
		tokens = append(tokens, lexer.NextToken())
	}
	expectedTokens := []Token{
		{TokenType: IF, Literal: "si"},
		{TokenType: LPAREN, Literal: "("},
		{TokenType: INT, Literal: "5"},
		{TokenType: LT, Literal: "<"},
		{TokenType: INT, Literal: "10"},
		{TokenType: RPAREN, Literal: ")"},
		{TokenType: LBRACE, Literal: "{"},
		{TokenType: RETURN, Literal: "regresa"},
		{TokenType: TRUE, Literal: "verdadero"},
		{TokenType: SEMICOLON, Literal: ";"},
		{TokenType: RBRACE, Literal: "}"},
		{TokenType: ELSE, Literal: "si_no"},
		{TokenType: LBRACE, Literal: "{"},
		{TokenType: RETURN, Literal: "regresa"},
		{TokenType: FALSE, Literal: "falso"},
		{TokenType: SEMICOLON, Literal: ";"},
		{TokenType: RBRACE, Literal: "}"},
	}

	for i := 0; i < len(tokens); i++ {
		if tokens[i].Literal != expectedTokens[i].Literal {
			t.Errorf("Expected these tokes have the same Literal value, but got %s for first and %s for sencond token", tokens[i].Literal, expectedTokens[i].Literal)
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
			t.Errorf("Expected these tokes have the same Literal value, but got %s for first and %s for sencond token", tokens[i].Literal, expectedTokens[i].Literal)
		} else if tokens[i].TokenType != expectedTokens[i].TokenType {
			t.Errorf("Expected these tokens was be the same TokenType , but got this [%s,%s]", tokens[i].TokenType, expectedTokens[i].TokenType)
		}
	}
}

func TestVariableLettersAndNumbers(t *testing.T) {
	source := "variable suma_1 = suma_casa2 + 5;"

	lexer := NewLexer(source)

	tokens := []Token{}

	for i := 0; i < 7; i++ {
		tokens = append(tokens, lexer.NextToken())
	}

	expectedTokens := []Token{
		{TokenType: LET, Literal: "variable"},
		{TokenType: IDENT, Literal: "suma_1"},
		{TokenType: ASSING, Literal: "="},
		{TokenType: IDENT, Literal: "suma_casa2"},
		{TokenType: PLUS, Literal: "+"},
		{TokenType: INT, Literal: "5"},
		{TokenType: SEMICOLON, Literal: ";"},
	}
	t.Log(tokens)
	for i := 0; i < len(tokens); i++ {
		if tokens[i].Literal != expectedTokens[i].Literal {
			t.Errorf("Expected these tokes have the same Literal value, but got %s for first and %s for sencond token", tokens[i].Literal, expectedTokens[i].Literal)
		} else if tokens[i].TokenType != expectedTokens[i].TokenType {
			t.Errorf("Expected these tokens was be the same TokenType , but got this [%s,%s]", tokens[i].TokenType, expectedTokens[i].TokenType)
		}
	}

}
