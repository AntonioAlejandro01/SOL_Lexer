package lexer

import (
	"fmt"
	"regexp"
)

type character string

func (c character) matchWith(pattern string) bool {
	m, err := regexp.MatchString(pattern, string(c))
	if err != nil {
		panic("Error with patterns in lexer")
	}
	return m
}

func (c character) isLetter() bool {
	return c.matchWith("[a-zA-z_]$")
}

func (c character) isNumber() bool {
	return c.matchWith("^[0-9]$")
}

func (c character) isEOF() bool {
	return c.matchWith("^$")
}

func (c character) isLParen() bool {
	return c.matchWith(fmt.Sprintf("^%s$", regexp.QuoteMeta("(")))
}

func (c character) isRParen() bool {
	return c.matchWith(fmt.Sprintf("^%s$", regexp.QuoteMeta(")")))
}

func (c character) isRBrace() bool {
	return c.matchWith(fmt.Sprintf("^%s$", regexp.QuoteMeta("}")))
}

func (c character) isLBrace() bool {
	return c.matchWith(fmt.Sprintf("^%s$", regexp.QuoteMeta("{")))
}

func (c character) isComma() bool {
	return c.matchWith(fmt.Sprintf("^%s$", regexp.QuoteMeta(",")))
}
func (c character) isSemicolon() bool {
	return c.matchWith(fmt.Sprintf("^%s$", regexp.QuoteMeta(";")))
}

func (c character) isQuote() bool {
	return c.matchWith(fmt.Sprintf("^%s$", regexp.QuoteMeta("\"")))
}

func (c character) isColon() bool {
	return c.matchWith(fmt.Sprintf("^%s$", regexp.QuoteMeta(":")))
}

func (c character) isDot() bool {
	return c.matchWith(fmt.Sprintf("^%s$", regexp.QuoteMeta(".")))
}

func (c character) isAll() bool {
	return c.matchWith(fmt.Sprintf("^%s$", regexp.QuoteMeta("*")))
}
