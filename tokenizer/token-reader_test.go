package tokenizer

import "testing"

func prepareTokenReader(str string) TokenReader {
	tokenizer := LoadTokenizer("../tests/tokenizer.json")

	return TokenReader{
		PossibleTokens: tokenizer.TokenList,
		ReadingIndex:   0,
		String:         str,
	}
}

func verifyTokenReaderRead(
	expectedToken Token,
	tokenReader TokenReader,
	t *testing.T,
) {
	token, err := tokenReader.ReadToken()

	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	if !token.HasSameValue(expectedToken) {
		t.Fatalf(
			"Tokens must have the same value\nRead: %s %s\nExpected: %s %s\n",
			token.Type,
			token.Value,
			expectedToken.Type,
			expectedToken.Value,
		)
	}

	expectedReadingIndex := len(expectedToken.Value) - 1
	isReadingIndexWrong := tokenReader.ReadingIndex == expectedReadingIndex
	if isReadingIndexWrong {
		t.Fatalf(
			"Reading index is wrong\nRead: %d\nExpected: %d\n",
			tokenReader.ReadingIndex,
			expectedReadingIndex,
		)
	}
}

type TestTokenReaderParams struct {
	String        string
	ExpectedToken Token
	T             *testing.T
}

func testTokenReader(params TestTokenReaderParams) {
	tokenReader := prepareTokenReader(params.String)

	verifyTokenReaderRead(
		params.ExpectedToken,
		tokenReader,
		params.T,
	)
}

func Test_TokenReader_ReadToken_Operator(t *testing.T) {
	testTokenReader(TestTokenReaderParams{
		T: t,
		ExpectedToken: Token{
			Type:  "operator",
			Value: "<",
		},
		String: "<Game>",
	})
}

func Test_TokenReader_ReadToken_Identifier(t *testing.T) {
	testTokenReader(TestTokenReaderParams{
		T: t,
		ExpectedToken: Token{
			Type:  "identifier",
			Value: "Game",
		},
		String: "Game>",
	})
}

func Test_TokenReader_ReadToken_Number(t *testing.T) {
	testTokenReader(TestTokenReaderParams{
		T: t,
		ExpectedToken: Token{
			Type:  "number",
			Value: "26",
		},
		String: "26",
	})
}

func Test_TokenReader_ReadToken_String(t *testing.T) {
	testTokenReader(TestTokenReaderParams{
		T: t,
		ExpectedToken: Token{
			Type:  "string",
			Value: "cham",
		},
		String: `"cham"}`,
	})
}
