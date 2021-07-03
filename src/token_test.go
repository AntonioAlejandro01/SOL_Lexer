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

	if lookupTokenType("as") != AS {
		t.Errorf("Expected as was AS, but got %s", lookupTokenType("as"))
	}

	if lookupTokenType("before") != BEFORE {
		t.Errorf("Expected before was BEFORE, but got %s", lookupTokenType("before"))
	}

	if lookupTokenType("body") != BODY {
		t.Errorf("Expected return was RETURN, but got %s", lookupTokenType("body"))
	}

	if lookupTokenType("DELETE") != DELETE {
		t.Errorf("Expected DELETE was DELETE, but got %s", lookupTokenType("DELETE"))
	}

	if lookupTokenType("errorsHandlers") != ERRORSHANDLERS {
		t.Errorf("Expected errorsHandlers was ERRORSHANDLERS, but got %s", lookupTokenType("errorsHandlers"))
	}

	if lookupTokenType("handler") != HANDLER {
		t.Errorf("Expected handler was HANDLER, but got %s", lookupTokenType("handler"))
	}

	if lookupTokenType("from") != FROM {
		t.Errorf("Expected from was FROM, but got %s", lookupTokenType("from"))
	}
	if lookupTokenType("GET") != GET {
		t.Errorf("Expected GET was GET, but got %s", lookupTokenType("GET"))
	}

	if lookupTokenType("HEAD") != HEAD {
		t.Errorf("Expected HEAD was HEAD, but got %s", lookupTokenType("HEAD"))
	}

	if lookupTokenType("OPTIONS") != OPTIONS {
		t.Errorf("Expected OPTIONS was OPTIONS, but got %s", lookupTokenType("OPTIONS"))
	}

	if lookupTokenType("PATCH") != PATCH {
		t.Errorf("Expected PATCH was PATCH, but got %s", lookupTokenType("PATCH"))
	}

	if lookupTokenType("POST") != POST {
		t.Errorf("Expected POST was POST, but got %s", lookupTokenType("POST"))
	}

	if lookupTokenType("PUT") != PUT {
		t.Errorf("Expected PUT was PUT, but got %s", lookupTokenType("PUT"))
	}

	if lookupTokenType("service") != SERVICE {
		t.Errorf("Expected service was SERVICE, but got %s", lookupTokenType("service"))
	}

	if lookupTokenType("headers") != HEADERS {
		t.Errorf("Expected headers was HEADERS, but got %s", lookupTokenType("headers"))
	}

	if lookupTokenType("import") != IMPORT {
		t.Errorf("Expected import was IMPORT, but got %s", lookupTokenType("import"))
	}

	if lookupTokenType("params") != PARAMS {
		t.Errorf("Expected params was PARAMS, but got %s", lookupTokenType("params"))
	}

	if lookupTokenType("X") != IDENT {
		t.Errorf("Expected X was IDENT, but got %s", lookupTokenType("X"))
	}

}
