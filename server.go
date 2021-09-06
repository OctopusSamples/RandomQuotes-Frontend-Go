package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

func main() {
	http.HandleFunc("/", serveTemplate)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}

	fmt.Printf("Starting server at port " + port + "\n")
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Printf("Error: Unable to bind to the port " + port)
	}
}

func getApiEndpoint() string {
	endpoint := os.Getenv("APIENDPOINT")
	if endpoint == "" {
		log.Print("Error: The APIENDPOINT environment variable was not set.")
		endpoint = "/api/quote"
	}
	return endpoint
}

func serveTemplate(w http.ResponseWriter, req *http.Request) {
	cleanPath := filepath.Clean(req.URL.Path)

	if strings.HasSuffix(cleanPath, ".js") {
		w.Header().Set("Content-Type", "text/javascript")
	} else if strings.HasSuffix(cleanPath, ".css") {
		w.Header().Set("Content-Type", "text/css")
	} else if strings.HasSuffix(cleanPath, ".html") {
		w.Header().Set("Content-Type", "html")
	} else if strings.HasSuffix(cleanPath, ".ico") {
		w.Header().Set("Content-Type", "image/png")
	}

	fp := filepath.Join("web", cleanPath)

	tmpl, err := template.ParseFiles(fp)
	if err == nil {
		tmpl.Execute(w, map[string]string{
			"api": getApiEndpoint(),
		})
	} else {
		w.WriteHeader(404)
	}
}
