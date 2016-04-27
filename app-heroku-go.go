package main

import (
	"database/sql"
	"fmt"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"os"
)

var (
	repeat int
	db     *sql.DB = nil
)

func homeHandler(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "\n\tHello World From Golang on Heroku")
}

func apiTest(w http.ResponseWriter, r *http.Request) {

	v := mux.Vars(r)
	fmt.Fprintf(w, "\n\t Hello%s", v["id"])
}

func dbFunc(w http.ResponseWriter, r *http.Request) {

	if _, err := db.Exec("CREATE TABLE IF NOT EXISTS ticks (tick timestamp)"); err != nil {
		fmt.Sprintf("Error creating database table: %q", err)
		return
	}

	http.redirect(w, r, "/", nil)
}
func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", homeHandler)
	r.HandleFunc("/db", dbFunc)
	r.HandleFunc("/a/{id}", apiTest)
	fmt.Println("Server Listening...")
	err := http.ListenAndServe(":"+os.Getenv("PORT"), r)
	if err != nil {
		panic("ListenAndServe")
	}
	db, err = sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatalf("Error opening database: %q", err)
	}

}
