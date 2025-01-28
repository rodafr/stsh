package main

import (
	"bufio"
	"io"
	"strings"
)

type Parser struct {
	currentStsh     *stsh
	currentSolution *solution
	currentFeature  *feature
	currentCommand  *command
	inCodeBlock     bool
}

func (p *Parser) ClaudeParse(r io.Reader) (*stsh, error) {
	scanner := bufio.NewScanner(r)
	p.currentStsh = &stsh{}

	for scanner.Scan() {
		line := scanner.Text()

		// Skip empty lines
		if strings.TrimSpace(line) == "" {
			continue
		}

		// Handle top level description
		if strings.HasPrefix(line, ">") {
			l := strings.TrimPrefix(line, ">")
			p.currentStsh.comment = l
		}

		// Handle code blocks
		if strings.HasPrefix(line, "```") {
			p.inCodeBlock = !p.inCodeBlock
			continue
		}

		// Handle command content when inside code block
		if p.inCodeBlock && p.currentCommand != nil {
			p.currentCommand.cmd = line
			continue
		}

		// Determine the heading level and content
		if strings.HasPrefix(line, "#") {
			level := 0
			for i := 0; i < len(line) && line[i] == '#'; i++ {
				level++
			}
			content := strings.TrimSpace(line[level:])

			switch level {
			case 1: // Top level - stsh header
				p.currentStsh.header = content
			case 2: // Solution level
				p.handleSolution(content)
			case 3: // Feature level
				p.handleFeature(content)
			case 4: // Command level
				p.handleCommand(content)
			}
			continue
		}

		// Handle tags
		if strings.HasPrefix(line, "Tags: ") {
			if p.currentCommand != nil {
				tags := strings.TrimPrefix(line, "Tags: ")
				p.currentCommand.tags = strings.Split(tags, ", ")
			}
			continue
		}

		// Any other text is considered description for the current command
		if p.currentCommand != nil && !p.inCodeBlock {
			if p.currentCommand.desc == "" {
				p.currentCommand.desc = line
			} else {
				p.currentCommand.desc += "\n" + line
			}
		}
	}

	return p.currentStsh, scanner.Err()
}

func (p *Parser) handleSolution(header string) {
	newSolution := &solution{header: header}
	p.currentStsh.sols = append(p.currentStsh.sols, *newSolution)
	p.currentSolution = &p.currentStsh.sols[len(p.currentStsh.sols)-1]
	p.currentFeature = nil
	p.currentCommand = nil
}

func (p *Parser) handleFeature(header string) {
	if p.currentSolution == nil {
		return
	}
	newFeature := &feature{header: header}
	p.currentSolution.feats = append(p.currentSolution.feats, *newFeature)
	p.currentFeature = &p.currentSolution.feats[len(p.currentSolution.feats)-1]
	p.currentCommand = nil
}

func (p *Parser) handleCommand(header string) {
	if p.currentFeature == nil {
		return
	}
	newCommand := &command{header: header}
	p.currentFeature.cmds = append(p.currentFeature.cmds, *newCommand)
	p.currentCommand = &p.currentFeature.cmds[len(p.currentFeature.cmds)-1]
}
