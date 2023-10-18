package main

import (
	"html/template"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

var tpl = template.Must(template.ParseFiles("index.html")) // This line parses an HTML template and stores it in the tp1 variable using the 'template.ParseFiles' function, template.Must is used to ensure that the template parsing succeds, if there's an error, it will panic.

func indexHandler(w http.ResponseWriter, r *http.Request) {
	tpl.Execute(w, nil)
} // * HTTP Request handler function
//	When a user accesses the root URL ("/") the indexHandler() function is called, it takes an HTTP response writer ('w') and a HTTP ('r')

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	fs := http.FileServer(http.Dir("assets"))

	mux := http.NewServeMux()

	mux.Handle("/assets/", http.StripPrefix("/assets/", fs))

	mux.HandleFunc("/", indexHandler)
	http.ListenAndServe(":"+port, mux)
}
