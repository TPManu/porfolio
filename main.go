package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("web/templates/index.html"))

	data := map[string]string{
		"Title": "Home Page",
	}

	if err := tmpl.Execute(w, data); err != nil {
		fmt.Println("Huh")
	}
}

func main() {
	router := http.NewServeMux()

	router.HandleFunc("GET /{$}", indexHandler)
	router.Handle("GET /static/",
		http.StripPrefix("/static/",
			http.FileServer(http.Dir("web/static")),
		),
	)
	fmt.Println("server starting on localhost:8080")
	http.ListenAndServe(":8080", router)
}
