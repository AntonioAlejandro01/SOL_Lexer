package lexer

import "testing"

func TestSpecificCharacters(t *testing.T) {
	chars := []character{"A", "3", "=", "+", "(", ")", "{", "}", ",", ";", "<", ">", "-", "*", "/", "!"}
	if !chars[0].isLetter() {
		t.Errorf("Expected that %v was a letter but didn't", chars[0])
	}
	if !chars[1].isNumber() {
		t.Errorf("Expected that %v was a number but didn't", chars[1])
	}
	if !chars[2].isAssing() {
		t.Errorf("Expected that %v was a = but didn't", chars[2])
	}
	if !chars[3].isPlus() {
		t.Errorf("Expected that %v was a + but didn't", chars[0])
	}
	if !chars[4].isLParen() {
		t.Errorf("Expected that %v was a ( but didn't", chars[0])
	}
	if !chars[5].isRParen() {
		t.Errorf("Expected that %v was a ) but didn't", chars[0])
	}
	if !chars[6].isLBrace() {
		t.Errorf("Expected that %v was a { but didn't", chars[0])
	}
	if !chars[7].isRBrace() {
		t.Errorf("Expected that %v was a } but didn't", chars[0])
	}
	if !chars[8].isComma() {
		t.Errorf("Expected that %v was a , but didn't", chars[0])
	}
	if !chars[9].isSemicolon() {
		t.Errorf("Expected that %v was a ; but didn't", chars[0])
	}
	if !chars[10].isLT() {
		t.Errorf("Expected that %v was a < but didn't", chars[0])
	}
	if !chars[11].isGT() {
		t.Errorf("Expected that %v was a > but didn't", chars[0])
	}
	if !chars[12].isMinus() {
		t.Errorf("Expected that %v was a - but didn't", chars[0])
	}
	if !chars[13].isAdd() {
		t.Errorf("Expected that %v was a * but didn't", chars[0])
	}
	if !chars[14].isDiv() {
		t.Errorf("Expected that %v was a / but didn't", chars[0])
	}
	if !chars[15].isNegation() {
		t.Errorf("Expected that %v was a ! but didn't", chars[0])
	}
}
