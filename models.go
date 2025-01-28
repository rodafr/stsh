package main

type stsh struct {
	header  string
	comment string
	sols    []solution
}

type solution struct {
	header string
	feats  []feature
}

type feature struct {
	header string
	cmds   []command
}

type command struct {
	header string
	desc   string
	cmd    string
	tags   []string
}
