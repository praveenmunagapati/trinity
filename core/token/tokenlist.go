package token

import (
	"errors"
	"strconv"
	"strings"
)

var (
	// ErrTokenListEmpty means that the resulting TokenList is empty
	ErrTokenListEmpty = errors.New("token list is empty")
)

// List is a slice of tokens
type List struct {
	filename string  // name of the file that these tokens belong to
	Tokens   []Token // the tokenized content of the file
}

/*
Add inserts a token into the TokenList
*/
func (list *List) Add(token Token) {
	list.Tokens = append(list.Tokens, token)
}

/*
Length returns the number of tokens in the TokenList
*/
func (list *List) Length() int {
	return len(list.Tokens)
}

/*
Filename returns the name of the file that these tokens
came form
*/
func (list *List) Filename() string {
	return list.filename
}

/*
SetFilename sets the name of the file that these tokens
came form
*/
func (list *List) SetFilename(name string) {
	list.filename = name
}

/*
RemoveNumbers deletes all tokens that are pure numbers
*/
func (list *List) RemoveNumbers() {
	for i := len(list.Tokens) - 1; i > 0; i-- {
		if _, err := strconv.Atoi(string(list.Tokens[i].Value)); err == nil {
			list.Tokens = append(list.Tokens[:i], list.Tokens[i+1:]...)
		}
	}
}

/*
RemoveLargerThan deletes all tokens that are longer than
some given length
*/
func (list *List) RemoveLargerThan(length int) {
	for i := len(list.Tokens) - 1; i > 0; i-- {
		if len(list.Tokens[i].Value) > length {
			list.Tokens = append(list.Tokens[:i], list.Tokens[i+1:]...)
		}
	}
}

/*
RemoveSmallerThan deletes all tokens that are longer than
some given length
*/
func (list *List) RemoveSmallerThan(length int) {
	for i := len(list.Tokens) - 1; i > 0; i-- {
		if len(list.Tokens[i].Value) < length {
			list.Tokens = append(list.Tokens[:i], list.Tokens[i+1:]...)
		}
	}
}

/*
RemoveHyphenedNums removes all tokens of the form "1234-5678"
*/
func (list *List) RemoveHyphenedNums() {
	for i := len(list.Tokens) - 1; i > 0; i-- {
		stringRep := string(list.Tokens[i].Value)
		if strings.Contains(stringRep, "-") {
			splitToken := strings.Split(stringRep, "-")
			allNumbers := true
			for _, element := range splitToken {
				if _, err := strconv.Atoi(element); err != nil {
					allNumbers = false
				}
			}
			if allNumbers {
				list.Tokens = append(list.Tokens[:i], list.Tokens[i+1:]...)
			}
		} else if strings.Contains(stringRep, "–") {

			splitToken := strings.Split(stringRep, "–")
			allNumbers := true
			for _, element := range splitToken {
				if _, err := strconv.Atoi(element); err != nil {
					allNumbers = false
				}
			}
			if allNumbers {
				list.Tokens = append(list.Tokens[:i], list.Tokens[i+1:]...)
			}
		}
	}
}

/*
RemoveConjunctions deletes all tokens that are conjunctions,
i.e. "and", "or",
*/
func (list *List) RemoveConjunctions(length int) {
	for i := len(list.Tokens) - 1; i > 0; i-- {
		if len(list.Tokens[i].Value) < length {
			list.Tokens = append(list.Tokens[:i], list.Tokens[i+1:]...)
		}
	}
}

/*
GetWithLength returns a new List with all items less than
or equal to the specified length
*/
func (list *List) GetWithLength(length int) (List, error) {
	result := List{}

	for _, token := range list.Tokens {
		if len(token.Value) <= length {
			result.Add(token)
		}
	}

	if list.Length() == 0 {
		return result, ErrTokenListEmpty
	}
	return result, nil
}

// func (list *List) Stem() {
// 	for _, token := range list.Tokens {
// 		token.Value = porterstemmer.Stem(token.Value)
// 	}
// }
