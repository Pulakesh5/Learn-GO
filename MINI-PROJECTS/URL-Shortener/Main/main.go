package main

import (
	"fmt"
	"net/http"

	urlshort "example.com"
)

func main() {
	mux := defaultMux()

	pathsToUrls := map[string]string{
		"/urlshort-godoc": "https://pkg.go.dev/github.com/georgeyord/go-url-shortener/pkg/urlshort",
		"/yaml-godoc":     "https://pkg.go.dev/gopkg.in/yaml.v2",
	}

	mapHandler := urlshort.MapHandler(pathsToUrls, mux)
	// Build the YAMLHandler using mapHandler as the fallback

	yaml := `
- path: /urlshort
  url: https://pkg.go.dev/github.com/georgeyord/go-url-shortener/pkg/urlshort
- path: /urlshort-final
  url: https://pkg.go.dev/gopkg.in/yaml.v2
`

	yamlHandler, err := urlshort.YAMLHandler([]byte(yaml), mapHandler)

	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	fmt.Println("Starting the server on :8080")
	http.ListenAndServe(":8080", yamlHandler)
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello World!")
}
