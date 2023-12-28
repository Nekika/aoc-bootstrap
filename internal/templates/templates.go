package templates

import (
	_ "embed"
	"github.com/nekika/aoc-bootstrap/internal/lang"
)

var (
	//go:embed day.exs.tmpl
	dayFileElixir []byte

	//go:embed day.go.tmpl
	dayFileGo []byte
)

func ForLanguage(l lang.Lang) []byte {
	switch l {
	case lang.Elixir:
		return dayFileElixir
	case lang.Go:
		return dayFileGo
	default:
		return make([]byte, 0)
	}
}
