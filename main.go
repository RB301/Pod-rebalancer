package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Hello, Welcome to my first Go App!")

//when base route Get:/ is clicked I want it to respond with this function
http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
fmt.Fprintf(w, "Welcome to my first Go App")
})

log.Fatal(http.ListenAndServe(":8080", nil))

}

