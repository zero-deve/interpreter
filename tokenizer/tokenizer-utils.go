package tokenizer

import (
	"encoding/json"

	"github.com/zero-deve/interpreter/files"
)

func LoadTokenizer(jsonFile string) Tokenizer {
	fileContent := files.ReadFile(jsonFile)

	tokenizer := Tokenizer{}

	err := json.Unmarshal(fileContent, &tokenizer)

	if err != nil {
		panic(err)
	}

	return tokenizer
}
