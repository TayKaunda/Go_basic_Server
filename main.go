package main

import (
	"fmt"
	"log"
	"net/http"
)

func main () {
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)

	fmt.Printf("starting server at port 4040\n")
	if err := http.ListenAndServe(":4040", nil); err != nil {
		log.Fatal(err)
	}
	
}

func helloHandler(write http.ResponseWriter, read *http.Request ) {
	if read.URL.Path != "/hello" {
	http.Error(write, "404 not found", http.StatusNotFound)
	return
	}

if read.Method != "GET" {
	http.Error(write, "method is not supported", http.StatusNotFound)
	return
}
fmt.Fprint(write, "Hello!")

}

func formHandler(write http.ResponseWriter, read *http.Request) {
	if err := read.ParseForm(); err != nil {
		fmt.Fprintf(write, "ParseForm() err: %v", err)
		return
	} 
	fmt.Fprintf(write, "POST request successful")
	name := read.FormValue("name")
	address := read.FormValue("address")
	fmt.Fprintf(write, "Name %s\n", name)
	fmt.Fprintf(write, "Address = %s\n", address )

	}

	