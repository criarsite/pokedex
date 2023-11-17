package main

import (
	"fmt"
	"net/http"
	"text/template"
)

var views = template.Must(template.ParseGlob("views/*"))

func main() {
	fmt.Println("Pokedex - Bootcamp DIO")
	fs := http.FileServer(http.Dir("assets"))
	http.Handle("/assets/", http.StripPrefix("/assets/", fs))

	http.HandleFunc("/", Inicio)
	http.HandleFunc("/detail", Detail)

	fmt.Println(`Server running on http://localhost:8080`)
	http.ListenAndServe(":8080", nil)
}

func Inicio(w http.ResponseWriter, r *http.Request) {

	views.ExecuteTemplate(w, "inicio", nil)
}

func Detail(w http.ResponseWriter, r *http.Request) {

	views.ExecuteTemplate(w, "detail", nil)
}

