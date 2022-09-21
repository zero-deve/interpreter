package tokenizer

import (
	"testing"
)

func Test_Tokenizer_IdentifyTokens(t *testing.T) {
	tokenizer := LoadTokenizer("../tests/tokenizer.json")
	tokenList, err := tokenizer.Tokenize(`<Game><WindowInfo title="Window title" size={{x: 500, y: 800}} /></Game>`)

	if err != nil {
		t.Fatal(err)
	}

	expectedTokenList := TokenList{
		Tokens: []Token{
			{
				Value: "<",
				Type:  "operator",
			},
			{
				Value: "Game",
				Type:  "identifier",
			},
			{
				Value: ">",
				Type:  "operator",
			},
			{
				Value: "<",
				Type:  "operator",
			},
			{
				Value: "WindowInfo",
				Type:  "identifier",
			},
			{
				Value: "title",
				Type:  "identifier",
			},
			{
				Value: "=",
				Type:  "operator",
			},
			{
				Value: "Window title",
				Type:  "string",
			},
			{
				Value: "size",
				Type:  "identifier",
			},
			{
				Value: "=",
				Type:  "operator",
			},
			{
				Value: "{",
				Type:  "operator",
			},
			{
				Value: "{",
				Type:  "operator",
			},
			{
				Value: "x",
				Type:  "identifier",
			},
			{
				Value: ":",
				Type:  "operator",
			},
			{
				Value: "500",
				Type:  "number",
			},
			{
				Value: ",",
				Type:  "operator",
			},
			{
				Value: "y",
				Type:  "identifier",
			},
			{
				Value: ":",
				Type:  "operator",
			},
			{
				Value: "800",
				Type:  "number",
			},
			{
				Value: "}",
				Type:  "operator",
			},
			{
				Value: "}",
				Type:  "operator",
			},
			{
				Value: "/>",
				Type:  "operator",
			},
			{
				Value: "</",
				Type:  "operator",
			},
			{
				Value: "Game",
				Type:  "identifier",
			},
			{
				Value: ">",
				Type:  "operator",
			},
		},
	}

	if !expectedTokenList.HasSameValue(tokenList) {
		t.Log("Generated token list differs")
		t.Log("Generated========================")
		printTokenList(tokenList, t)
		t.Log("Expected========================")
		printTokenList(expectedTokenList, t)
		t.FailNow()
	}
}

func printTokenList(tokenList TokenList, t *testing.T) {
	for _, token := range tokenList.Tokens {
		t.Logf("%s %s\n", token.Type, token.Value)
	}
}
