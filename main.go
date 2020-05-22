package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

var port = flag.Int("port", 2525, "port to listen on")

func main() {
	flag.Parse()
	http.Handle("/", http.FileServer(http.Dir(".")))
	log.Printf("Serving current directory, listening on http://localhost:%d", *port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), nil))
}
