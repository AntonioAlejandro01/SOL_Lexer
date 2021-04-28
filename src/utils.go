package lexer

import "strings"

func filterSlice(source string) []string {
	sourceSlice := strings.Split(source, "")
	newSource := []string{}
	for _, character := range sourceSlice {
		if character != "" && character != "\n" && character != "\t" {
			newSource = append(newSource, character)
		}
	}
	return newSource
}
