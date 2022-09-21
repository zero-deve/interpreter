package tokenizer

import (
	"testing"
)

func Test_Token_Equals(t *testing.T) {
	token := Token{
		Character: "<",
		Type:      "operator",
	}

	if !token.Equals(token) {
		t.Fatal("Token must be equal")
	}
}

func Test_Token_IsPossibleMatch_Normal(t *testing.T) {
	token := Token{
		Character: "</",
	}

	if !token.IsPossibleMatch("<") {
		t.Fatal("< is a possible </ token")
	}
}

func Test_Token_Not_IsPossibleMatch_Normal(t *testing.T) {
	token := Token{
		Character: "<",
	}

	if token.IsPossibleMatch("</") {
		t.Fatal("</ is not a possible < token")
	}
}

func Test_Token_IsPossibleMatch_RegEx(t *testing.T) {
	tokenValue := "^([aA-zZ]+)$"
	token := Token{
		Character: tokenValue,
		IsRegEx:   true,
	}

	value := "variableName"

	if !token.IsPossibleMatch(value) {
		t.Fatalf("%s must be a possible match of %s\n", value, tokenValue)
	}
}

func Test_Token_Not_IsPossibleMatch_RegEx(t *testing.T) {
	tokenValue := "^([aA-zZ]+)$"
	token := Token{
		Character: tokenValue,
		IsRegEx:   true,
	}

	value := "#variableName"

	if token.IsPossibleMatch(value) {
		t.Fatalf("%s must not be a possible match of %s\n", value, tokenValue)
	}
}

func Test_Token_IsPossibleMatch_StringLike(t *testing.T) {
	token := Token{
		IsStringLike: true,
		Openning:     "'",
	}

	value := "'This is"

	if !token.IsPossibleMatch(value) {
		t.Fatalf("%s must be a possible match of %s(.*)%s\n", value, token.Openning, token.Openning)
	}
}

func Test_Token_Not_IsPossibleMatch_StringLike(t *testing.T) {
	token := Token{
		IsStringLike: true,
		Openning:     "'",
	}

	value := "This '"

	if token.IsPossibleMatch(value) {
		t.Fatalf("%s must not be a possible match of %s(.*)%s\n", value, token.Openning, token.Openning)
	}
}

func Test_Token_IsPerfectMatch_Normal(t *testing.T) {
	token := Token{
		Character: "</",
	}

	if !token.IsPerfectMatch("</") {
		t.Fatal("</ must be perfect match of </")
	}
}

func Test_Token_Not_IsPerfectMatch_Normal(t *testing.T) {
	token := Token{
		Character: "</",
	}

	if token.IsPerfectMatch("<") {
		t.Fatal("< must not be perfect match of </")
	}
}

func Test_Token_IsPerfectMatch_RegEx(t *testing.T) {
	possibleValue := "^([aA-zZ]+)$"
	token := Token{
		Character: possibleValue,
		IsRegEx:   true,
	}

	value := "variableName"

	if !token.IsPerfectMatch(value) {
		t.Fatalf("%s must be a perfect match of %s\n", value, possibleValue)
	}
}

func Test_Token_Not_IsPerfectMatch_RegEx(t *testing.T) {
	possibleValue := "^([aA-zZ]+)$"
	token := Token{
		Character: possibleValue,
		IsRegEx:   true,
	}

	value := "variableName#"

	if token.IsPerfectMatch(value) {
		t.Fatalf("%s must not be a perfect match of %s\n", value, possibleValue)
	}
}
func Test_Token_IsPerfectMatch_StringLike(t *testing.T) {
	token := Token{
		IsStringLike: true,
		Openning:     "'",
	}

	value := "'This is a string'"

	if !token.IsPerfectMatch(value) {
		t.Fatalf("%s must be a perfect match of %s(.*)%s\n", value, token.Openning, token.Openning)
	}
}

func Test_Token_Not_IsPerfectMatch_StringLike(t *testing.T) {
	token := Token{
		IsStringLike: true,
		Openning:     "'",
	}

	value := "'This is a string"

	if token.IsPerfectMatch(value) {
		t.Fatalf("%s must not be a perfect match of %s(.*)%s\n", value, token.Openning, token.Openning)
	}
}
