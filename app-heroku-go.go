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

func apiTest(w http.ResponseWriter, r *http.Request) {

	v := mux.Vars(r)
	fmt.Fprintf(w, "\n\t Hello%s", v["id"])
}
func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", homeHandler)
	r.HandleFunc("/a/{id}", apiTest)
	fmt.Println("Server Listening...")
	err := http.ListenAndServe(":"+os.Getenv("PORT"), r)
	if err != nil {
		panic("ListenAndServe")
	}
}
