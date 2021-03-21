package gosh

import (
	"strings"
)

// A helper from collecting string flags into an array
type StringFlags []string

func (sf *StringFlags) String() string {
	return strings.Join(*sf, " ")
}

func (sf *StringFlags) Set(value string) error {
	*sf = append(*sf, value)
	return nil
}
