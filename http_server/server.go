package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/RaraykinValery/l0/database"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	data := make([]string, 1)
	t, _ := template.ParseFiles("index.html")
	t.Execute(w, data)
}

func orderHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodPost {
		uuid := r.FormValue("uuid")

		order, err := database.GetOrderFromDB(uuid)
		if err != nil {
			log.Printf("Couldn't get order from database: %s", err)
			return
		}

		fmt.Fprintf(w, "Requested uuid: %v", order)
	}
}

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/orders", orderHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
