package server

import (
	"fmt"
	"log"
	"net/http"
)

type People struct {
	name string
	age  int
}

type Vehicle struct {
	model string
	old   int
}

func (p People) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, p)
}

func (v *Vehicle) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, v)
}

func StartExampleServer() {
	fmt.Println("Starting the Example Web Server - try http://localhost:8090 !!")
	fs := http.FileServer(http.Dir("./files"))
	http.Handle("/people/kamal", People{"kamal", 45})
	http.Handle("/people/anish", People{"anish", 12})
	http.Handle("/vehicle/car", &Vehicle{"car", 5})
	http.Handle("/files/", http.StripPrefix("/files/", fs))
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		log.Printf("received call for path %s\n", r.URL.Path)
		fmt.Fprintf(rw, "%s , I love %s!", PrintMessage(), r.URL.Path[1:])
	})
	log.Fatal(http.ListenAndServe(":8090", nil))
}

// PrintMessage Function to Test Export Functionality using Modules
func PrintMessage() string {
	return "Hello World !!!"
}
