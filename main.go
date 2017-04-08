package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

func extractChildlen(key string, raw []byte) (interface{}, error) {
	var root interface{}
	err := json.Unmarshal(raw, &root)
	if err != nil {
		panic(err)
	}
	n, ok := root.(map[string]interface{})
	if !ok {
		return nil, errors.New("invalid structure")
	}

	e, ok := n[key]
	if !ok {
		return nil, errors.New("invalid structure")
	}

	return e, nil
}

func main() {
	jsonStr := []byte(`{
		"hoge": [
			"a", "b", "c"
		]
	}`)

	e, err := extractChildlen("hoge", jsonStr)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	fmt.Println(e)
}
