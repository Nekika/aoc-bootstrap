package files

import (
	"fmt"
	"os"
	"path"

	"github.com/nekika/aoc-bootstrap/internal/lang"
	"github.com/nekika/aoc-bootstrap/internal/templates"
)

func write(dst, name string, data []byte) error {
	p := path.Join(dst, name)
	return os.WriteFile(p, data, 0o755)
}

func WriteInput(dst string, data []byte) error {
	return write(dst, "input.txt", data)
}

func WriteDay(dst string, n int, l lang.Lang) error {
	name := fmt.Sprintf("%d%s", n, lang.Extension(l))
	tmpl := templates.ForLanguage(l)
	return write(dst, name, tmpl)
}

func WriteExample(dst string, data []byte) error {
	return write(dst, "example.txt", data)
}

func WriteReadme(dst string, data []byte) error {
	return write(dst, "README.md", data)
}
