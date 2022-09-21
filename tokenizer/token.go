package tokenizer

import "regexp"

type Token struct {
	Character string
	Type      string
	Value     string

	IsRegEx      bool
	IsStringLike bool
	IsUseless    bool

	Openning string
	Ending   string
}

func (token Token) Equals(another Token) bool {
	return token.Character == another.Character &&
		token.Type == another.Type
}

func (token Token) HasSameValue(another Token) bool {
	return token.Type == another.Type &&
		token.Value == another.Value
}

func (token Token) IsPossibleMatch(str string) bool {
	if token.IsRegEx {
		return token.IsPossibleRegExMatch(str)
	}

	if token.IsStringLike {
		return token.IsPossibleStringLikeMatch(str)
	}

	if len(str) > len(token.Character) {
		return false
	}

	for index := 0; index < len(str); index++ {
		if str[index] != token.Character[index] {
			return false
		}
	}

	return true
}

func (token Token) IsPossibleRegExMatch(str string) bool {
	match, _ := regexp.MatchString(token.Character, str)

	return match
}

func (token Token) IsPossibleStringLikeMatch(str string) bool {
	if len(str) == 0 {
		return false
	}

	isItAlreadyClosed := len(str) > 2 && string(str[len(str)-2]) == token.GetEnding()
	if isItAlreadyClosed {
		return false
	}

	return string(str[0]) == token.Openning
}

func (token Token) IsPerfectMatch(str string) bool {
	if token.IsRegEx {
		return token.IsPossibleRegExMatch(str)
	}

	if token.IsStringLike {
		return token.IsPerfectStringLikeMatch(str)
	}

	return str == token.Character
}

func (token Token) IsPerfectStringLikeMatch(str string) bool {
	if len(str) < 2 {
		return false
	}

	return string(str[0]) == token.Openning &&
		string(str[len(str)-1]) == token.GetEnding()
}

func (token Token) GetEnding() string {
	if token.Ending != "" {
		return token.Ending
	}

	return token.Openning
}

func (token *Token) ValueAsString() {
	if len(token.Value) < 2 {
		return
	}

	token.Value = token.Value[1 : len(token.Value)-1]
}
