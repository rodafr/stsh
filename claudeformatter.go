package main

import (
	"errors"
	"fmt"
	"io"
	"strings"
)

// Format writes the stsh struct as a markdown file to the provided writer
func Format(w io.Writer, s *stsh) error {
	// Write header
	if s.header != "" {
		if _, err := fmt.Fprintf(w, "# %s\n\n", s.header); err != nil {
			return err
		}
	} else {
		return errors.New("empty header line")
	}

	// Write comment if exists
	if s.comment != "" {
		if _, err := fmt.Fprintf(w, "> %s\n\n", s.comment); err != nil {
			return err
		}
	} else {
		return errors.New("empty comment field")
	}

	if len(s.sols) == 0 {
		return errors.New("no h2 found")
	}

	// Write each solution
	for _, sol := range s.sols {
		if sol.header != "" {
			if _, err := fmt.Fprintf(w, "## %s\n\n", sol.header); err != nil {
				return err
			}
		} else {
			return errors.New("empty solution header line")
		}

		if len(sol.feats) == 0 {
			return errors.New("no h3 found")
		}
		// Write features
		for _, feat := range sol.feats {
			if feat.header != "" {
				if _, err := fmt.Fprintf(w, "### %s\n\n", feat.header); err != nil {
					return err
				}
			} else {
				return errors.New("empty feature header line")
			}

			if len(feat.cmds) == 0 {
				return errors.New("no h4 found")
			}
			// Write commands
			for _, cmd := range feat.cmds {
				// Command header
				if cmd.header != "" {
					if _, err := fmt.Fprintf(w, "#### %s\n\n", cmd.header); err != nil {
						return err
					}
				} else {
					return errors.New("empty command header line")
				}

				// Command description
				if cmd.desc != "" {
					if _, err := fmt.Fprintf(w, "%s\n\n", cmd.desc); err != nil {
						return err
					}
				} else {
					return errors.New("empty command description")
				}

				// Command code block
				if cmd.cmd != "" {
					if _, err := fmt.Fprintf(w, "```sh\n%s\n```\n\n", cmd.cmd); err != nil {
						return err
					}
				} else {
					return errors.New("empty command code block")
				}

				// Command tags
				if len(cmd.tags) > 0 {
					if _, err := fmt.Fprintf(w, "Tags: %s\n\n", strings.Join(cmd.tags, ", ")); err != nil {
						return err
					}
				}
			}
		}
	}

	return nil
}

// FormatToString returns the markdown as a string
func FormatToString(s *stsh) (string, error) {
	var b strings.Builder
	err := Format(&b, s)
	if err != nil {
		return "", err
	}
	return b.String(), nil
}
