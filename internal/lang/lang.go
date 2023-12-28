package lang

import "slices"

type Lang int

const (
	Elixir Lang = iota
	Go
)

var aliases = map[Lang][]string{
	Elixir: {"elixir", "el", "elx"},
	Go:     {"go", "golang"},
}

var extensions = map[Lang]string{
	Elixir: ".exs",
	Go:     ".go",
}

// FromAlias returns the lang associated with the alias a.
// It returns -1 if a is not associated  with a known lang.
func FromAlias(a string) Lang {
	for key, value := range aliases {
		if slices.Contains(value, a) {
			return key
		}
	}

	return -1
}

// Extension returns the appropriate file extension for the lang l.
// It returns an empty string if l is not a supported lang.
func Extension(l Lang) string {
	ext, ok := extensions[l]
	if !ok {
		return ""
	}
	return ext
}
