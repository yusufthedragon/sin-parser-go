package sinparser

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestParseSIN(t *testing.T) {
	test, err := ParseSIN("3204110609970001")

	if err != nil {
		panic(err)
	}

	s, _ := json.MarshalIndent(test, "", "\t")

	fmt.Println(string(s))
}
