package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var file = "./stsh.md"

func parse(f string) (stsh, error) {
	var s stsh

	input, err := os.Open(file)
	if err != nil {
		return s, err
	}
	defer input.Close()

	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		line := scanner.Text()
		// fmt.Printf("line: %s\n", line)
		if strings.HasPrefix(line, "# ") {
			l := strings.TrimPrefix(line, "# ")
			s.header = l
		}
		if strings.HasPrefix(line, "> ") {
			l := strings.TrimPrefix(line, "> ")
			s.comment = l
		}
		if strings.HasPrefix(line, "## ") {
			l := strings.TrimPrefix(line, "## ")
			sol := solution{header: l}
			s.sols = append(s.sols, sol)
		}

	}
	fmt.Printf("%+v\n", s)

	return s, nil
}
