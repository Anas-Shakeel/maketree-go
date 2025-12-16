package specparser

import (
	"fmt"
	"strings"
)

// TODO: Implement the tokenizer

type TokenType int

const (
	TokenDir TokenType = iota
	TokenFile
)

type Token struct {
	Type     TokenType // Type of the token
	Value    string    // Name of the file or dir
	RawValue string    // Raw line text (for debugging and errors)
	Line     int       // Line no. of the token in file
	Level    int       // Indent level
}

func (t *Token) String() string {
	var tokenType string
	if t.Type == TokenDir {
		tokenType = "Directory"
	} else {
		tokenType = "File"
	}
	return fmt.Sprintf("Type: %s\nValue: %q\nRawValue: %q\nLine: %d\nLevel: %d",
		tokenType, t.Value, t.RawValue, t.Line, t.Level)
}

// tokenize tokenizes the input string and returns a slice of
// tokens []Token or error.
func tokenize(input string) (*[]Token, error) {
	var tokens []Token
	var lineNumber int

	// Iterate over lines in input
	for line := range strings.Lines(input) {
		lineNumber++

		value := strings.TrimSpace(line)

		// Is it an Empty line?
		if value == "" {
			continue
		}

		// Is it a line comment?
		if strings.HasPrefix(value, "#") {
			continue
		}

		// Validate indentation
		// TODO: Reject tabs or mixed indentation with an error
		if strings.HasPrefix(line, " ") {
			if indentCount := countIndent(line); indentCount%4 != 0 {
				return nil, &ErrInvalidIndent{
					Line:   lineNumber,
					Found:  indentCount,
					Reason: "indentation is not a multiple of 4 spaces",
				}
			}
		}

		var token Token

		// Is it a directory?
		if cleaned, found := strings.CutSuffix(value, "/"); found {
			token.Type = TokenDir
			value = cleaned // Name without '/' slash
		} else {
			token.Type = TokenFile
		}

		token.Value = value
		token.RawValue = line
		token.Line = lineNumber
		token.Level = (len(line) - len(strings.TrimLeft(line, " "))) / 4

		tokens = append(tokens, token)
		fmt.Println(token.String())
	}

	return &tokens, nil
}

// countIndent returns the number of spaces used in indentation
func countIndent(line string) int {
	var count int

	// Iterate over runes in line
	for _, r := range line {
		if r != ' ' {
			break
		}
		count++
	}
	return count
}
