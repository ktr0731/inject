package main

import (
	"bytes"
	"encoding/json"
	"testing"
)

func Test_extractChildlen(t *testing.T) {
	testsuite := []struct {
		key    string
		raw    []byte
		expect []byte
	}{
		{
			"hoge",
			[]byte(`{"hoge": ["a", "b", "c"]}`),
			[]byte(`["a","b","c"]`),
		},
	}

	for _, test := range testsuite {
		childlen, err := extractChildlen(test.key, test.raw)
		if err != nil {
			t.Errorf("%s in %s: %s", test.key, string(test.raw), err)
		}

		b, err := json.Marshal(childlen)
		if err != nil {
			t.Error(err)
		}

		if !bytes.Equal(test.expect, b) {
			t.Errorf("expect: %s, actual: %s", string(test.expect), string(b))
		}
	}
}
