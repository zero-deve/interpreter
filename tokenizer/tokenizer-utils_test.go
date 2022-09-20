package tokenizer

import (
	"testing"
)

func Test_LoadTokenizer(t *testing.T) {
	tokenizer := LoadTokenizer("../tests/tokenizer.json")

	tokenList := TokenList{
		Tokens: []Token{
			{
				Character: "<",
				Type:      "operator",
			},
			{
				Character: ">",
				Type:      "operator",
			},
			{
				Character: "/>",
				Type:      "operator",
			},
			{
				Character: "</",
				Type:      "operator",
			},
			{
				Character: "=",
				Type:      "operator",
			},
			{
				Character: "{",
				Type:      "operator",
			},
			{
				Character: "}",
				Type:      "operator",
			},
			{
				Character: ":",
				Type:      "operator",
			},
			{
				Type:      "string",
				Delimiter: "'",
			},
			{
				Type: "identifier",
			},
			{
				Type: "number",
			},
		},
	}

	if !tokenList.Equals(tokenizer.TokenList) {
		t.Fatalf("Tokenizer didn't load correctly\nExpected tokens: %v\nLoaded tokens: %v\n", tokenList, tokenizer.TokenList)
	}
}
