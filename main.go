package main

import (

	"log"
	"net/http"

	"github.com/gabrielluizsf/dll_call/handlers"
)

const port = ":4000"

func main() {
	http.HandleFunc("/v1/hello",handlers.HelloWorld)
	log.Println("Server listening on port"+port)
	log.Fatal(http.ListenAndServe(port, nil))
}