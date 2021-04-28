package lexer

import "testing"

func TestFilterSlice(t *testing.T) {
	source := "Testing\n\t Test"
	slice := filterSlice(source)
	expectedSlice := []string{
		"T",
		"e",
		"s",
		"t",
		"i",
		"n",
		"g",
		" ",
		"T",
		"e",
		"s",
		"t",
	}

	for i := 0; i < len(slice); i++ {
		if slice[i] != expectedSlice[i] {
			t.Errorf("Expected %s, but got %s", expectedSlice[i], slice[i])
		}
	}

}

func TestEmptyString(t *testing.T) {
	source := ""
	slice := filterSlice(source)
	lenSpected := 0

	if lenSpected != len(slice) {
		t.Errorf("Expected empty slice, but got %d", len(slice))
	}
}
