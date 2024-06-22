package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

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
	h := cyoa.NewHandler(story)
	fmt.Printf("Starting the server on port %d\n", *port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), h))
	// fmt.Printf("%+v\n", story)
}
