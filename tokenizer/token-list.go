package tokenizer

import (
	"errors"
	"fmt"

	"github.com/zero-deve/interpreter/sliceutils"
)

type TokenList struct {
	Tokens []Token
}

func (list TokenList) Equals(another TokenList) bool {
	return sliceutils.CompareSlice(
		list.Tokens,
		another.Tokens,
		func(a Token, b Token) bool {
			return a.Equals(b)
		},
	)
}

func (list TokenList) HasSameValue(another TokenList) bool {
	return sliceutils.CompareSlice(
		list.Tokens,
		another.Tokens,
		func(a Token, b Token) bool {
			return a.HasSameValue(b)
		},
	)
}

func (list TokenList) Empty() bool {
	return len(list.Tokens) == 0
}

func (list TokenList) FindPossibleTokens(str string) TokenList {
	return TokenList{
		Tokens: sliceutils.FilterSlice(list.Tokens, func(token Token) bool {
			return token.IsPossibleMatch(str)
		}),
	}
}

func (list TokenList) FindPerfectMatchToken(tokenValue string, nextTokenValue string) (Token, error) {
	perfectMatchs := TokenList{
		Tokens: sliceutils.FilterSlice(list.Tokens, func(token Token) bool {
			return token.IsPerfectMatch(tokenValue)
		}),
	}

	if perfectMatchs.Empty() {
		return Token{}, errors.New(fmt.Sprintf("Unknown token %s\n", nextTokenValue))
	}

	if len(perfectMatchs.Tokens) > 1 {
		return Token{}, errors.New(fmt.Sprintf("Unable to identify token %s\n", nextTokenValue))
	}

	perfectMatchToken := perfectMatchs.Tokens[0]
	perfectMatchToken.Value = tokenValue

	if perfectMatchToken.IsStringLike {
		perfectMatchToken.ValueAsString()
	}

	return perfectMatchToken, nil
}
