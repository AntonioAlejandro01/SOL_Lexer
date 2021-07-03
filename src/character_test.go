package lexer

import "testing"

func TestSpecificCharacters(t *testing.T) {
	chars := []character{"A", "3", "(", ")", "{", "}", ",", ";", ".", ":", "\"", ""}
	if !chars[0].isLetter() {
		t.Errorf("Expected that %v was a letter but didn't", chars[0])
	}
	if !chars[1].isNumber() {
		t.Errorf("Expected that %v was a number but didn't", chars[1])
	}

	if !chars[2].isLParen() {
		t.Errorf("Expected that %v was a ( but didn't", chars[2])
	}

	if !chars[3].isRParen() {
		t.Errorf("Expected that %v was a ) but didn't", chars[3])
	}

	if !chars[4].isLBrace() {
		t.Errorf("Expected that %v was a { but didn't", chars[4])
	}

	if !chars[5].isRBrace() {
		t.Errorf("Expected that %v was a } but didn't", chars[5])
	}

	if !chars[6].isComma() {
		t.Errorf("Expected that %v was a , but didn't", chars[6])
	}

	if !chars[7].isSemicolon() {
		t.Errorf("Expected that %v was a ; but didn't", chars[7])
	}

	if !chars[8].isDot() {
		t.Errorf("Expected that %v was a . but didn't", chars[8])
	}

	if !chars[9].isColon() {
		t.Errorf("Expected that %v was a : but didn't", chars[9])
	}

	if !chars[10].isQuote() {
		t.Errorf("Expected that %v was a \" but didn't", chars[10])
	}

	if !chars[11].isEOF() {
		t.Errorf("Expected that %v was a EOF(\"\")but didn't", chars[11])
	}
}
