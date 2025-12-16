// parser implements the parser that parses source spec file
package specparser

import (
	"fmt"
	"os"

	"github.com/anas-shakeel/maketree-go/internal/models"
)

// Parse parses the file and returns a *Node tree or an error.
//
// filename: name of the file that defines a valid file structure
func Parse(filename string) (*models.Node, error) {
	// Read the file
	contents, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	// Parse the file
	nodetree, err := ParseString(string(contents))
	if err != nil {
		return nil, err
	}

	return nodetree, nil

}

// Parse parses the input string (spec tree) and returns a *Node tree or an error
//
// input: the input string that defines a valid file structure
func ParseString(input string) (*models.Node, error) {
	// Tokenize the input
	tokens, err := tokenize(input)
	if err != nil {
		return nil, err
	}

	// Parse the tokens
	nodetree, err := parseTokens(tokens)
	if err != nil {
		return nil, err
	}

	return nodetree, nil

}

// parseTokens parses the tokens and returns a pointer to Node tree
func parseTokens(tokens *[]Token) (*models.Node, error) {
	// TODO: Implement the parsing logic
	fmt.Printf("%#v\n", tokens)

	return &models.Node{}, nil
}
