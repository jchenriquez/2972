package validparenthesis

import (
	"bufio"
	"encoding/json"
	"io"
	"os"
	"testing"
)

type Test struct {
	Input  string `json:"input"`
	Output bool   `json:"output"`
}

func TestIsValid(t *testing.T) {
	f, err := os.Open("tests.json")

	if err != nil {
		t.Error(err)
		return
	}

	defer f.Close()

	var tests map[string]Test
	reader := bufio.NewReader(f)
	decoder := json.NewDecoder(reader)

	for {
		err = decoder.Decode(&tests)

		if err == nil {
			for name, test := range tests {
				t.Run(name, func(k *testing.T) {
					result := IsValid(test.Input)

					if result != test.Output {
						k.Errorf("result returned is %v", result)
					}
				})
			}
		} else if err == io.EOF {
			break
		} else {
			t.Error(err)
			break
		}
	}

}
