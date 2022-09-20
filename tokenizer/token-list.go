package tokenizer

type TokenList struct {
	Tokens []Token
}

func (list TokenList) Equals(another TokenList) bool {
	if len(list.Tokens) != len(another.Tokens) {
		return false
	}

	for tokenIndex, token := range list.Tokens {
		if !token.Equals(another.Tokens[tokenIndex]) {
			return false
		}
	}

	return true
}
