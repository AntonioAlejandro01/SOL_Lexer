package lexer

import (
	"fmt"
	"regexp"
	"strings"
)

// Lexer - It's create tokens for source that it's passed as string . IMPORTANT!! no create and lexer , use the function NewLexer
type Lexer struct {
	source       []string
	character    character
	readPosition int
	position     int
}

// NextToken -  Next token
func (l *Lexer) NextToken() *Token {
	skipWhiteSpace(l)
	var tt TokenType
	var literal string
	if l.character.isLetter() {
		literal = l.readIdentifier()
		tt = lookupTokenType(literal)
	} else if l.character.isColon() {
		literal = string(l.character)
		tt = COLON
		l.readCharacter()
	} else if l.character.isLParen() {
		literal = string(l.character)
		tt = LPAREN
		l.readCharacter()

	} else if l.character.isRParen() {
		literal = string(l.character)
		tt = RPAREN
		l.readCharacter()

	} else if l.character.isLBrace() {
		literal = string(l.character)
		tt = LBRACE
		l.readCharacter()

	} else if l.character.isRBrace() {
		literal = string(l.character)
		tt = RBRACE
		l.readCharacter()

	} else if l.character.isComma() {
		literal = string(l.character)
		tt = COMMA
		l.readCharacter()

	} else if l.character.isSemicolon() {
		literal = string(l.character)
		tt = SEMICOLON
		l.readCharacter()

	} else if l.character.isDot() {
		literal = string(l.character)
		tt = DOT
		l.readCharacter()

	} else if l.character.isQuote() {
		literal = string(l.character)
		tt = QUOTES
		l.readCharacter()

	} else if l.character.isAll() {
		literal = string(l.character)
		tt = ALL
		l.readCharacter()

	} else if l.character.isEOF() {
		literal = string(l.character)
		tt = EOF
	} else {
		literal = string(l.character)
		tt = ILLEGAL
		l.readCharacter()
	}

	return &Token{
		TokenType: tt,
		Literal:   literal,
	}
}

func (l *Lexer) readCharacter() {
	if l.readPosition >= len(l.source) {
		l.character = character("")
	} else {
		l.character = character(l.source[l.readPosition])
	}
	l.position = l.readPosition
	l.readPosition++
}

func (l *Lexer) readIdentifier() string {
	initialPosition := l.position
	for l.character.isLetter() || l.character.isNumber() {
		l.readCharacter()
	}

	return strings.Join(l.source[initialPosition:l.position], "")

}

func skipWhiteSpace(l *Lexer) {
	regSpace, err := regexp.Compile(fmt.Sprintf("^%s$", regexp.QuoteMeta(" ")))
	if err != nil {
		panic(err)
	}
	isSpace := regSpace.MatchString(string(l.character))
	for isSpace {
		l.readCharacter()
		isSpace = regSpace.MatchString(string(l.character))
	}
}

// NewLexer create a lexer with string passed as parameter
func NewLexer(source string) Lexer {
	l := Lexer{
		source:       filterSlice(source),
		readPosition: 0,
		position:     0,
	}
	l.readCharacter()
	return l
}
