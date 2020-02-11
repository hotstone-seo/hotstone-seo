package urlstore

import (
	"strings"
)

// Param is parameter in string as represented by <PARAM>
type Param struct {
	Start int
	End   int

	// string before the parameter
	StringBefore string

	// string after the parameter
	StringAfter string

	// parameter in raw string
	Raw string

	// name of parameter
	Name string

	// pattern of parameter
	Pattern string

	AtLastPos bool
}

// CreateParam to create first occurence param in string
func CreateParam(s string) *Param {
	var (
		name    string
		pattern string
		start   = -1
		end     = -1
	)

	for i := 0; i < len(s); i++ {
		if start < 0 && s[i] == '<' {
			start = i
		}
		if start >= 0 && s[i] == '>' {
			end = i
			break
		}
	}

	if start < 0 || end < 0 {
		return nil
	}

	inside := s[start+1 : end]
	if i := strings.Index(inside, ":"); i >= 0 {
		name = inside[:i]
		pattern = inside[i+1:]
	} else {
		name = inside
	}

	return &Param{
		Start:        start,
		End:          end,
		Raw:          s[start : end+1],
		StringBefore: s[:start],
		StringAfter:  s[end+1:],
		Name:         name,
		Pattern:      pattern,
		AtLastPos:    end >= len(s)-1,
	}
}
