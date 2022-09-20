package tokenizer

type Tokenizer struct {
	TokenList TokenList
}

func (tokenizer Tokenizer) Tokenize(str string) (TokenList, error) {
	return TokenList{}, nil
}
