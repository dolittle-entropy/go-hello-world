package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gorilla/mux"
)

func main() {
	bind := os.Getenv("BIND")
	if bind == "" {
		bind = "localhost"
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	urlPrefix := os.Getenv("URL_PREFIX")
	urlPrefix = strings.TrimSuffix(urlPrefix, "/")
	if urlPrefix == "" {
		urlPrefix = "/"
	}

	if !strings.HasPrefix(urlPrefix, "/") {
		log.Fatal("URL_PREFIX needs to start with /")
	}

	fmt.Printf("Running URL_PREFIX:%s on %s:%s\n", urlPrefix, bind, port)

	router := mux.NewRouter()
	// Handle redirect if just urlPrefix
	router.Handle(urlPrefix, http.RedirectHandler(urlPrefix+"/", http.StatusPermanentRedirect))

	s := router.PathPrefix(urlPrefix).Subrouter()
	s.HandleFunc("/", helloRoot)
	s.HandleFunc("/hi", helloHi)
	s.HandleFunc("/hello-world", helloWorld)

	fmt.Println("Starting up on " + port)
	log.Fatal(http.ListenAndServe(bind+":"+port, router))
}

func helloRoot(w http.ResponseWriter, req *http.Request) {
	fmt.Println("RequestURI", req.RequestURI)
	fmt.Fprintln(w, "Hi Root!")
}

func helloHi(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(w, "Hi!")
}

func helloWorld(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(w, "Hello world!")
}
