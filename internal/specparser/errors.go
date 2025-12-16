package specparser

import "fmt"

// ErrInvalidIndent occurs when indentation is inconsistent/invalid
type ErrInvalidIndent struct {
	Line   int
	Found  int // Number of indent spaces found
	Reason string
}

func (e *ErrInvalidIndent) Error() string {
	return fmt.Sprintf(
		"invalid indentation at line %d: %s (found %d spaces)",
		e.Line, e.Reason, e.Found,
	)
}
