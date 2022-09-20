package files

import (
	"fmt"
	"io/ioutil"
)

func ReadFile(fileName string) []byte {
	buffer, err := ioutil.ReadFile(fileName)

	if err != nil {
		panic(fmt.Sprintf("Unable to read file %s\n%v\n", fileName, err))
	}

	return buffer
}
