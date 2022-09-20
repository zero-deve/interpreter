package tokenizer

import (
	"testing"
)

func Test_TokenList_Equals(t *testing.T) {
	tokenList := TokenList{
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
				Character: "/>",
				Type:      "operator",
			},
		},
	}

	if !tokenList.Equals(tokenList) {
		t.Fatal("Token list must be equal")
	}
}
