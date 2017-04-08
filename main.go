package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	jsonStr := []byte(`{
		"hoge": [
			"a", "b", "c"
		]
	}`)

	var root interface{}
	err := json.Unmarshal(jsonStr, &root)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%#v", root)
}
