package main

import (
    "encoding/json"
    "log"
    "net/http"
    "github.com/gorilla/mux"
)

// our main function
func main() {
	router := mux.NewRouter()
	router.HandleFunc("")
	router.HandleFunc("")
	router.HandleFunc("")
	router.HandleFunc("")
	router.HandleFunc("")
    log.Fatal(http.ListenAndServe(":8000", router))
}