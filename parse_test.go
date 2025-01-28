package main

import (
	"fmt"
	"os"
	"testing"
)

func TestParse(t *testing.T) {
	file := "./stsh.md"
	parse(file)
}

func TestClauseParserAndFormatter(t *testing.T) {
	file := "./stsh.md"
	input, err := os.Open(file)
	if err != nil {
		t.Fail()
	}
	defer input.Close()
	var p Parser
	s, err := p.ClaudeParse(input)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("claude says \n%+v\n", s)

	md, err := FormatToString(s)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("claude says \n%s\n", md)
}
