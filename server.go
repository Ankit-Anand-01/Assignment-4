package main

import (
	"fmt"
	"log"
	"net/http"
)

//formhandler accepts the /form request
func formHandler(w http.ResponseWriter, r *http.Request) {
	//error handling
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}
	fmt.Fprintf(w, "POST request successful\n")
	name := r.FormValue("name")
	city := r.FormValue("city")
	state := r.FormValue("state")
	fmt.Fprintf(w, "Name = %s\n", name)
	fmt.Fprintf(w, "City = %s\n", city)
	fmt.Fprintf(w, "State = %s\n", state)

}

//handle all the request related to hello
func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		//error handling
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}

	if r.Method != "GET" {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}

	fmt.Fprintf(w, "Hello! welcome here. Now you can fill out the form through form.html request.")
}

func main() {
	//to serve the static files
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)

	//to start the web server
	fmt.Printf("Starting server at port 8000\n")
	if err := http.ListenAndServe(":8000", nil); err != nil {
		log.Fatal(err)
	}
}
