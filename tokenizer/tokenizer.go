package tokenizer

import "errors"

type Tokenizer struct {
	TokenList TokenList
}

func (tokenizer Tokenizer) Tokenize(str string) (TokenList, error) {
	readTokenParams := TokenizerReadTokenParams{
		String:       str,
		ReadingIndex: 0,
	}

	tokenList := TokenList{}
	var err error = nil

	for err == nil {
		token, readError := tokenizer.ReadToken(&readTokenParams)

		if readError != nil {
			if readError.Error() != "No more tokens" {
				err = readError
			}

			break
		}

		tokenList.Tokens = append(tokenList.Tokens, token)
	}

	return tokenList, err
}

type TokenizerReadTokenParams struct {
	String       string
	ReadingIndex int
}

func (tokenizer Tokenizer) ReadToken(params *TokenizerReadTokenParams) (Token, error) {
	if params.ReadingIndex == len(params.String) {
		return Token{}, errors.New("No more tokens")
	}

	isReadTokenUseless := true
	token := Token{}
	var err error = nil

	for isReadTokenUseless {
		tokenReader := TokenReader{
			PossibleTokens: tokenizer.TokenList,
			ReadingIndex:   params.ReadingIndex,
			String:         params.String,
		}

		token, err = tokenReader.ReadToken()
		params.ReadingIndex = tokenReader.ReadingIndex
		isReadTokenUseless = token.IsUseless
	}

	return token, err
}
