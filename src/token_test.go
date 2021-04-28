package lexer

import "testing"

func TestToken(t *testing.T) {
	token := &Token{TokenType: ILLEGAL, Literal: "@"}

	expectedString := "Type: ILLEGAL, Literal: @"

	if token.ToString() != expectedString {
		t.Errorf("Expected \"%s\", but got \"%s\"", token.ToString(), expectedString)
	}
}

func TestLookupTokenType(t *testing.T) {
	if lookupTokenType("false") != FALSE {
		t.Errorf("Expected false was FALSE, but got %s", lookupTokenType("false"))
	}
	if lookupTokenType("fun") != FUNCTION {
		t.Errorf("Expected fun was FUNCTION, but got %s", lookupTokenType("fun"))
	}
	if lookupTokenType("return") != RETURN {
		t.Errorf("Expected return was RETURN, but got %s", lookupTokenType("return"))

	}
	if lookupTokenType("if") != IF {
		t.Errorf("Expected if was IF, but got %s", lookupTokenType("if"))

	}
	if lookupTokenType("else") != ELSE {
		t.Errorf("Expected else was ELSE, but got %s", lookupTokenType("else"))

	}
	if lookupTokenType("let") != LET {
		t.Errorf("Expected let was LET, but got %s", lookupTokenType("let"))

	}
	if lookupTokenType("true") != TRUE {
		t.Errorf("Expected true was TRUE, but got %s", lookupTokenType("true"))

	}
	if lookupTokenType("X") != IDENT {
		t.Errorf("Expected X was IDENT, but got %s", lookupTokenType("X"))

	}

}
