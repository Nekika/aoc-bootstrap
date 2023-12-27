package main

import (
	"fmt"
	"github.com/nekika/aoc-bootstrap/internal/templates"
	"io"
	"net/http"
	"os"
	"path"
	"regexp"
	"strconv"

	md "github.com/JohannesKaufmann/html-to-markdown"
)

func main() {
	var token string
	fmt.Print("Token: ")
	_, err := fmt.Scanln(&token)

	year := 2023

	day := 10

	dayurl := fmt.Sprintf("https://adventofcode.com/%d/day/%d", year, day)
	dayreq, err := http.NewRequest(http.MethodGet, dayurl, nil)
	if err != nil {
		panic(err)
	}
	dayreq.AddCookie(&http.Cookie{Name: "session", Value: token})
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
	inputreq, err := http.NewRequest("GET", inputurl, nil)
	if err != nil {
		panic(err)
	}
	inputreq.AddCookie(&http.Cookie{Name: "session", Value: token})
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
	if err := os.WriteFile(path.Join(dirname, dayfilename), templates.DayFileElixir, 0o755); err != nil {
		panic(err)
	}

	inputfilename := "input.txt"
	if err := os.WriteFile(path.Join(dirname, inputfilename), input, 0o755); err != nil {
		panic(err)
	}
}
