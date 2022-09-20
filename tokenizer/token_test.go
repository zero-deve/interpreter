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
