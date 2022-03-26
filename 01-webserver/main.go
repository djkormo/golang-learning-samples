package main

import (
	"fmt"
	"net/http"
	"log"
)

func main() {
	
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/",fileServer)
	http.HandleFunc("/form",formHandler)
	http.HandleFunc("/hello",helloHandler)
	fmt.Println("Starting server at port 8000")
	if err:= http.ListenAndServe(":8000",nil);err!=nil {
      log.Fatal(err)
	}
}

func formHandler(w http.ResponseWriter, r *http.Request) {
  if err := r.ParseForm(); err!=nil {
	  fmt.Fprintf(w,"Error parsing form PerseForm(): %v",err)
  }
    fmt.Fprintf(w, "POST request successful")
	name := r.FormValue("name")
	address := r.FormValue("address")
	fmt.Fprintf(w, "Name = %s\n", name)
	fmt.Fprintf(w, "Address = %s\n", address)
}

func helloHandler(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path!="/hello" {
		http.Error(w,"404 not found",http.StatusNotFound)
		return
	}

	if r.Method !="GET" {
		http.Error(w,"Method not supported found",http.StatusMethodNotAllowed)
		return
	}
   fmt.Fprintf(w,"Hello!") 
}