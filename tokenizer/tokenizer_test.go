package tokenizer

import (
	"testing"
)

func Test_Tokenizer_IdentifyTokens(t *testing.T) {
	tokenizer := Tokenizer{}
	tokenList, err := tokenizer.Tokenize("<Game><WindowInfo title='Window title' size={{x: 500, y: 800}} /></Game>")

	if err != nil {
		t.Fatal(err)
	}

	expectedTokenList := TokenList{
		Tokens: []Token{
			{
				Character: "<",
				Type:      "operator",
			},
			{
				Character: "Game",
				Type:      "identifier",
			},
			{
				Character: ">",
				Type:      "operator",
			},
			{
				Character: "<",
				Type:      "operator",
			},
			{
				Character: "WindowInfo",
				Type:      "identifier",
			},
			{
				Character: "title",
				Type:      "identifier",
			},
			{
				Character: "=",
				Type:      "operator",
			},
			{
				Character: "Window title",
				Type:      "string",
			},
			{
				Character: "size",
				Type:      "identifier",
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
				Character: "{",
				Type:      "operator",
			},
			{
				Character: "x",
				Type:      "identifier",
			},
			{
				Character: ":",
				Type:      "operator",
			},
			{
				Character: "500",
				Type:      "number",
			},
			{
				Character: ",",
				Type:      "operator",
			},
			{
				Character: "y",
				Type:      "identifier",
			},
			{
				Character: ":",
				Type:      "operator",
			},
			{
				Character: "800",
				Type:      "number",
			},
			{
				Character: "}",
				Type:      "operator",
			},
			{
				Character: "}",
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
				Character: "Game",
				Type:      "identifier",
			},
			{
				Character: ">",
				Type:      "operator",
			},
		},
	}

	if !expectedTokenList.Equals(tokenList) {
		t.Fatal("Generated token list differs")
	}
}
