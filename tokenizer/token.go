package tokenizer

type Token struct {
	Character string
	Type      string
	Delimiter string
}

func (token Token) Equals(another Token) bool {
	return token.Character == another.Character &&
		token.Type == another.Type
}
