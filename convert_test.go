package feed2json_test

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"testing"

	"github.com/carlmjohnson/feed2json"
)

func readFile(t *testing.T, name string) []byte {
	t.Helper()
	b, err := ioutil.ReadFile(name)
	if err != nil {
		t.Fatalf("unexpected error reading file %q in testing: %v ",
			name, err)
	}
	return b
}

func equalJSON(t *testing.T, expect, have []byte) {
	var err error
	var expectData interface{}
	if err = json.Unmarshal(expect, &expectData); err != nil {
		t.Fatalf("unexpected error loading JSON %q: %v ",
			expect, err)
	}
	if expect, err = json.MarshalIndent(expectData, "", "  "); err != nil {
		t.Fatalf("unexpected error tidying JSON %q: %v ",
			expect, err)
	}
	var haveData interface{}
	if err = json.Unmarshal(have, &haveData); err != nil {
		t.Fatalf("unexpected error loading JSON %q: %v ",
			have, err)
	}
	if have, err = json.MarshalIndent(haveData, "", "  "); err != nil {
		t.Fatalf("unexpected error tidying JSON %q: %v ",
			have, err)
	}
	if string(have) != string(expect) {
		t.Fatalf("expect %q; have %q", string(have), string(expect))
	}
}

func TestConvert(t *testing.T) {
	inputFile := "testdata/rss.xml"
	outputFile := "testdata/feed.json"

	rssBuf := bytes.NewBuffer(readFile(t, inputFile))
	var jsonBuf bytes.Buffer

	if err := feed2json.Convert(rssBuf, &jsonBuf); err != nil {
		t.Fatalf("unexpected error converting %q testing: %v ",
			inputFile, err)
	}

	output := readFile(t, outputFile)
	equalJSON(t, output, jsonBuf.Bytes())
}
