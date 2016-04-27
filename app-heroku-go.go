package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"os"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "\n\tHello World From Golang on Heroku")
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", homeHandler)

	fmt.Println("Server Listening...")
	err := http.ListenAndServe(":"+os.Getenv("PORT"), r)
	if err != nil {
		panic("ListenAndServe")
	}
}
