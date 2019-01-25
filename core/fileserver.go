package main

import (
	"log"
	"net/http"
)

func main() {

	log.Fatal(http.ListenAndServe(":9900", http.FileServer(http.Dir("./"))))
}
