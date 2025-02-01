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

func TestClaudeFormatter(t *testing.T) {
	s := &stsh{
		header:  "test stsh header",
		comment: "test stsh comment",
		sols: []solution{
			{
				header: "test sol header",
				feats: []feature{
					{
						header: "test feat header",
						cmds: []command{
							{
								header: "test cmd header",
								desc:   "test cmd desc",
								cmd:    "ping 127.0.0.1",
								tags: []string{
									"ping", "test",
								},
							},
						},
					},
				},
			},
		},
	}
	fmt.Print(FormatToString(s))
}
