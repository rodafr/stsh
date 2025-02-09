package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

const (
	spellbook string = "splbk.md"
)

func main() {
	// define subcommands and flags
	storeCmd := flag.NewFlagSet("store", flag.ExitOnError)
	listCmd := flag.NewFlagSet("list", flag.ExitOnError)

	if len(os.Args) < 2 {
		fmt.Println("expected 'store' subcommand")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "store":
		storeCmd.Parse(os.Args[2:])
		cmd := os.Args[2:]
		cmds := strings.Join(cmd, " ")
		cmds += "\n"
		writeFile(cmds)
	case "list":
		listCmd.Parse(os.Args[2:])
		file, err := readFile()
		fmt.Printf("file:\n%s\n", file)
		if err != nil {
			panic(err)
		}
	default:
		fmt.Println("eh?")
	}
}

func readCmd() (string, error) {
	r := bufio.NewReader(os.Stdin)
	cmd, err := r.ReadString('\n')
	if err != nil {
		return "", fmt.Errorf("couldn't read from stdin: %w", err)
	}

	return cmd, err
}

func writeFile(cmd string) error {
	// create new file or open stash file to append to it
	f, err := os.OpenFile(stashfile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		return fmt.Errorf("couldn't open file: %w", err)
	}
	_, err = f.WriteString(cmd)
	if err != nil {
		return fmt.Errorf("couldn't write to file: %w", err)
	}

	// close the opened file and handle error
	err = f.Close()
	if err != nil {
		panic(err)
	}
	return err
}

func readFile() (string, error) {
	cmd, err := os.ReadFile(stashfile)
	return string(cmd), err
}
