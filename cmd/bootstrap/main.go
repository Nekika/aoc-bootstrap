package main

import (
	"flag"
	"fmt"
	"github.com/nekika/aoc-bootstrap/internal"
	"github.com/nekika/aoc-bootstrap/internal/lang"
	"github.com/nekika/aoc-bootstrap/internal/templates"
	"io"
	"net/http"
	"os"
	"path"
	"regexp"
	"strconv"
	"time"

	md "github.com/JohannesKaufmann/html-to-markdown"
)

func main() {
	var (
		day  int
		year int

		langalias string

		token string
	)

	flag.IntVar(&day, "d", time.Now().Day(), "day number")
	flag.IntVar(&year, "y", time.Now().Year(), "year number")
	flag.StringVar(&langalias, "l", "", "programming language alias")
	flag.StringVar(&token, "t", "", "session token")
	flag.Parse()

	l := lang.FromAlias(langalias)
	if l == -1 {
		msg := fmt.Sprintf("alias %s is not a known alias to a supported programming language", langalias)
		panic(msg)
	}

	if token == "" {
		panic("token flag missing")
	}

	dayurl := fmt.Sprintf("https://adventofcode.com/%d/day/%d", year, day)
	dayreq, err := internal.HttpGetWithSessionCookie(dayurl, token)
	dayres, err := http.DefaultClient.Do(dayreq)
	if err != nil {
		panic(err)
	}

	doc, err := io.ReadAll(dayres.Body)
	if err != nil && err != io.EOF {
		panic(err)
	}

	mainre := regexp.MustCompile("(?s)<main>.*</main>")
	mainelem := mainre.Find(doc)
	if len(mainelem) == 0 {
		panic("error locating main node in problem document")
	}

	examplere := regexp.MustCompile("(?sU)<pre><code>(.*)</code></pre>")
	example := examplere.FindSubmatch(mainelem)[1]

	inputurl := fmt.Sprintf("%s/input", dayurl)
	inputreq, err := internal.HttpGetWithSessionCookie(inputurl, token)
	inputres, err := http.DefaultClient.Do(inputreq)
	if err != nil {
		panic(err)
	}
	input, err := io.ReadAll(inputres.Body)
	if err != nil {
		panic(err)
	}

	dirname := strconv.Itoa(day)
	if err := os.Mkdir(dirname, 0o755); err != nil {
		panic(err)
	}

	converter := md.NewConverter("", true, nil)

	readme, err := converter.ConvertBytes(mainelem)
	if err != nil {
		panic(err)
	}

	if err := os.WriteFile(path.Join(dirname, "README.md"), readme, 0o755); err != nil {
		panic(err)
	}

	if err := os.WriteFile(path.Join(dirname, "example.txt"), example, 0o755); err != nil {
		panic(err)
	}

	dayfilename := fmt.Sprintf("%d.exs", day)
	tmpl := templates.ForLanguage(l)
	if err := os.WriteFile(path.Join(dirname, dayfilename), tmpl, 0o755); err != nil {
		panic(err)
	}

	inputfilename := "input.txt"
	if err := os.WriteFile(path.Join(dirname, inputfilename), input, 0o755); err != nil {
		panic(err)
	}
}
