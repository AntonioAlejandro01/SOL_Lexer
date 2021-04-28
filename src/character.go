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

func (c character) isAssing() bool {
	return c.matchWith("^=$")
}

func (c character) isPlus() bool {
	return c.matchWith(fmt.Sprintf("^%s$", regexp.QuoteMeta("+")))
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
func (c character) isLT() bool {
	return c.matchWith(fmt.Sprintf("^%s$", regexp.QuoteMeta("<")))
}
func (c character) isGT() bool {
	return c.matchWith(fmt.Sprintf("^%s$", regexp.QuoteMeta(">")))
}
func (c character) isMinus() bool {
	return c.matchWith(fmt.Sprintf("^%s$", regexp.QuoteMeta("-")))
}
func (c character) isDiv() bool {
	return c.matchWith(fmt.Sprintf("^%s$", regexp.QuoteMeta("/")))
}
func (c character) isAdd() bool {
	return c.matchWith(fmt.Sprintf("^%s$", regexp.QuoteMeta("*")))
}
func (c character) isNegation() bool {
	return c.matchWith(fmt.Sprintf("^%s$", regexp.QuoteMeta("!")))
}
