package main

import (
	"fmt"
	"log"
	"net/http"
)

var fs = http.FileServer(http.Dir("assets/"))

func main() {
	fmt.Println("Starting the Programme again - try http://<HOSTNAME> !!")
	http.HandleFunc("/", pathHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func printMessage() string {
	return "Hello World !!!"
}

func pathHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("received call for path %s\n", r.URL.Path[1:])
	fmt.Fprintf(w, "%s , I love %s!", printMessage(), r.URL.Path[1:])
}
