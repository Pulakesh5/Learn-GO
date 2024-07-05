package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"text/template"

	cyoa "example.com"
)

func main() {
	port := flag.Int("port", 3000, "denotes the port no of the running server")
	fileName := flag.String("file", "gopher.json", "the JSON file with the CYOA soty")
	flag.Parse()
	// fmt.Printf("Using the story in %s.\n", *fileName)

	file, err := os.Open(*fileName)
	if err != nil {
		panic(err)
	}
	story, err := cyoa.JsonStory(file)
	if err != nil {
		panic(err)
	}

	tpl := template.Must(template.New("").Parse(storyTmplt))
	h := cyoa.NewHandler(story,
		cyoa.WithTemplate(tpl),
	)
	// mux := http.NewServeMux()
	// mux.Handle("/story/", h)
	fmt.Printf("Starting the server on port %d\n", *port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), h))
}

func pathFn(r *http.Request) string {
	path := strings.TrimSpace(r.URL.Path)
	if path == "/story" || path == "/story/" {
		path = "story/intro"
	}
	path = path[len("/story/"):]
	return path
}

var storyTmplt = `<!DOCTYPE html>
<html>
    <head>
        <title>Choose Your Own Adventure</title>
        <meta charset="utf-8">
    </head>
    <body>
        <h1>{{.Title}}</h1>
        {{range .Paragraphs}}
        <p>{{.}}</p>
        {{end}}
        <ul>
        {{range .Options}}
            <li><a href="/{{.Chapter}}">{{.Text}}</a></li>
        {{end}}
        </ul>
    </body>
</html>`
