package tokenizer

import (
	"errors"
)

type TokenReader struct {
	PossibleTokens TokenList
	ReadingIndex   int
	String         string
}

func (reader *TokenReader) ReadToken() (Token, error) {
	prevTokenValue := ""
	newTokenValue := prevTokenValue

	for !reader.PossibleTokens.Empty() {
		newPossibleTokens := TokenList{}

		if !reader.HasReachedLastCharacter() {
			newTokenValue += string(reader.String[reader.ReadingIndex])
			newPossibleTokens = reader.PossibleTokens.FindPossibleTokens(newTokenValue)
		}

		if newPossibleTokens.Empty() {
			return reader.PossibleTokens.FindPerfectMatchToken(prevTokenValue, newTokenValue)
		} else {
			reader.ReadingIndex++
			reader.PossibleTokens = newPossibleTokens
			prevTokenValue = newTokenValue
		}
	}

	return Token{}, errors.New("Unable to identify token")
}

func (reader TokenReader) HasReachedLastCharacter() bool {
	return len(reader.String) == reader.ReadingIndex
}
