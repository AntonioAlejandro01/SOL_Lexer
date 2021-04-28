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
func (l *Lexer) NextToken() Token {
	skipWhiteSpace(l)
	var tt TokenType
	var literal string
	if l.character.isAssing() {
		if l.nextCharacter().matchWith("^=$") {
			tt, literal = l.makeTwoCharacterToken(EQ)
		} else {
			tt = ASSING
			literal = string(l.character)
		}
		l.readCharacter()

	} else if l.character.isPlus() {
		tt = PLUS
		literal = string(l.character)
		l.readCharacter()

	} else if l.character.isEOF() {
		literal = string(l.character)
		tt = EOF
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

	} else if l.character.isLT() {
		literal = string(l.character)
		tt = LT
		l.readCharacter()

	} else if l.character.isGT() {
		literal = string(l.character)
		tt = GT
		l.readCharacter()

	} else if l.character.isMinus() {
		literal = string(l.character)
		tt = MINUS
		l.readCharacter()

	} else if l.character.isDiv() {
		literal = string(l.character)
		tt = DIV
		l.readCharacter()

	} else if l.character.isAdd() {
		literal = string(l.character)
		tt = MULTIPLICATION
		l.readCharacter()

	} else if l.character.isNegation() {
		if l.nextCharacter().matchWith("^=$") {
			tt, literal = l.makeTwoCharacterToken(NOTEQ)
		} else {
			literal = string(l.character)
			tt = NEGATION
		}
		l.readCharacter()

	} else if l.character.isLetter() {
		literal = l.readIdentifier()
		tt = lookupTokenType(literal)
	} else if l.character.isNumber() {
		literal = l.readNumber()
		tt = INT
	} else {
		literal = string(l.character)
		tt = ILLEGAL
		l.readCharacter()

	}
	newToken := Token{
		TokenType: tt,
		Literal:   literal,
	}
	return newToken
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

func (l *Lexer) readNumber() string {
	initialPosition := l.position

	for l.character.isNumber() {
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

func (l *Lexer) makeTwoCharacterToken(tokentype TokenType) (TokenType, string) {
	prefix := l.character
	l.readCharacter()
	suffix := l.character

	return tokentype, fmt.Sprintf("%s%s", prefix, suffix)
}

func (l Lexer) nextCharacter() character {
	if l.readPosition >= len(l.source) {
		return character("")
	}
	return character(l.source[l.readPosition])
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
