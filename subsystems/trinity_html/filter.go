package thtml

import "C"

import (
	"strings"

	"github.com/clownpriest/trinity/core/token"
)

const (
	// MaxTokenLength is the maximum length of a token
	MaxTokenLength = 20
)

/*
FilterPipeline is the main function which runs the raw block
of text through all the different filters, resuling in a list
of tokens
*/
func FilterPipeline(str string) token.List {
	str = FilterStringBlock(str)

	tokens := token.Tokenize(str)

	// get rid of all pure numbers, e.g. "33", "209853", etc...
	tokens.RemoveNumbers()
	// get rid of all numbers separated by hyphens, e.g. "123-098", "0-1992", "4-67-6666"
	tokens.RemoveHyphenedNums()
	// get rid of tokens shorter than 2 characters
	tokens.RemoveSmallerThan(2)
	// get rid of tokens longer than MaxTokenLength
	tokens.RemoveLargerThan(MaxTokenLength)
	// stem all the tokens
	// tokens.Stem()

	return tokens
}

/*
FilterStringBlock takes a block of string, removes punctuation,
and casts everything to uppercase
*/
func FilterStringBlock(str string) string {
	str = FilterPunct(str)     // get rid of puncuation
	str = strings.ToLower(str) // make all uppercase
	return str
}

/*
FilterPunct removes all the unnecessary punctuation
from a block of text
*/
func FilterPunct(str string) string {
	str = strings.Replace(str, "\"", " ", -1)
	str = strings.Replace(str, "/", " ", -1)
	str = strings.Replace(str, "\\", " ", -1)
	str = strings.Replace(str, "“", " ", -1)
	str = strings.Replace(str, ".", " ", -1)
	str = strings.Replace(str, ",", " ", -1)
	str = strings.Replace(str, ";", " ", -1)
	str = strings.Replace(str, ":", " ", -1)
	str = strings.Replace(str, "'s", " ", -1)
	str = strings.Replace(str, "'", "", -1)
	str = strings.Replace(str, "^", " ", -1)
	str = strings.Replace(str, "[", " ", -1)
	str = strings.Replace(str, "]", " ", -1)
	str = strings.Replace(str, "{", " ", -1)
	str = strings.Replace(str, "}", " ", -1)
	str = strings.Replace(str, "(", " ", -1)
	str = strings.Replace(str, ")", " ", -1)
	str = strings.Replace(str, "'s", "", -1)

	return str
}
