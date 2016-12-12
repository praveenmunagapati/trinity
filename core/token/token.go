package token

import "strings"

/*
Token is the atomic unit of tokenization
*/
type Token struct {
	Value    []rune
	Location uint32
}

/*
NewToken returns a new Token initialized with the given
value and document location
*/
func NewToken(value string, location uint32) Token {
	return Token{Value: []rune(value), Location: location}
}

/*
Tokenize turns a string block into a token.List
*/
func Tokenize(str string) List {
	// split string block on whitespace
	var tokens []string
	tokens = strings.Fields(str)

	// turn []string into token.List
	var tokenList List
	for i, currToken := range tokens {
		tokenList.Add(NewToken(currToken, uint32(i)))
	}
	return tokenList
}
