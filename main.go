package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"

	"kylehodgetts.com/url-shortener/handler"
)

const yamlString = `
- path: /urlshort
  url: https://github.com/gophercises/urlshort
- path: /urlshort-final
  url: https://github.com/gophercises/urlshort/tree/solution
`

const jsonString = `
[
  {
    "path": "/urlshort",
    "url": "https://github.com/gophercises/urlshort"
  },
  {
    "path": "/urlshort-final",
    "url": "https://github.com/gophercises/urlshort/tree/solution"
  }
]
`

var pathsToUrls = map[string]string{
	"/urlshort-godoc": "https://godoc.org/github.com/gophercises/urlshort",
	"/yaml-godoc":     "https://godoc.org/gopkg.in/yaml.v2",
}

func main() {
	var yamlFile = flag.String("urls-yaml", "./data/urls.yaml", "yaml file containing url pairs")
	var jsonFile = flag.String("urls-json", "./data/urls.json", "json file containing url pairs")
	flag.Parse()
	mux := defaultMux()

	// Build the MapHandler using the mux as the fallback
	mapHandler := handler.MapHandler(pathsToUrls, mux)

	// Build the JSONHandler using the mapHandler as the
	// fallback
	json := parseFile(*jsonFile, jsonString)
	jsonHandler, err := handler.JSONHandler(json, mapHandler)
	if err != nil {
		panic(err)
	}

	// Build the YAMLHandler using the jsonHandler as the
	// fallback
	yaml := parseFile(*yamlFile, yamlString)
	yamlHandler, err := handler.YAMLHandler(yaml, jsonHandler)
	if err != nil {
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
	fmt.Fprintln(w, "Hello, world!")
}

func parseFile(filePath string, defaultContent string) []byte {
	var b []byte
	b, err := ioutil.ReadFile(filePath)
	if err != nil {
		b = []byte(defaultContent)
	}

	return b
}