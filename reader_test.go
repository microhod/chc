package main

import (
	"regexp"
	"slices"
	"strings"
	"testing"
)

func TestReader(t *testing.T) {
	input := strings.NewReader("some text")
	expected := []CharCount{
		{Char: 't', Count: 2},
		{Char: 'e', Count: 2},
		{Char: 'x', Count: 1},
		{Char: 's', Count: 1},
		{Char: 'o', Count: 1},
		{Char: 'm', Count: 1},
		{Char: ' ', Count: 1},
	}

	reader := NewReader(regexp.MustCompile(".*"))
	err := reader.Read(input)
	if err != nil {
		t.Errorf("reading: %s", err)
	}

	actual := reader.Counts()
	if !slices.Equal(expected, actual) {
		t.Errorf("[%v] != [%v]", expected, actual)
	}
}
