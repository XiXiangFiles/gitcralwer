package cmd

import (
	"io/ioutil"
	"testing"
)

func TestParseMessage(t *testing.T) {
	expected, _ := ioutil.ReadFile("test_data/memo.txt")

	actual := ParseMessage("test_data/comments.json", "https://example.com/")
	if len(string(expected)) != len(actual) {
		t.Error("parse message error")
	}
}
