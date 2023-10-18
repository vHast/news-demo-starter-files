package main

import (
	"html/template"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

// 8ef3d8f6e3a44ac5858e1db1c19039c5 API KEY

var tpl = template.Must(template.ParseFiles("index.html")) // This line parses an HTML template and stores it in the tp1 variable using the 'template.ParseFiles' function, template.Must is used to ensure that the template parsing succeds, if there's an error, it will panic.

func indexHandler(w http.ResponseWriter, r *http.Request) {
	tpl.Execute(w, nil)
} // * HTTP Request handler function
//	When a user accesses the root URL ("/") the indexHandler() function is called, it takes an HTTP response writer ('w') and a HTTP ('r')

// Inside the function it executes the parsed HTML template (tpl)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}

	// Attempt to load environment variables from an .env file using the godotenv package, if an error is encountered, it will print an error message.

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	// Fetching the value of PORT using os.Getenv() if the PORT environment variable is not set, it defaults to the value of 3000

	fs := http.FileServer(http.Dir("assets"))

	// This line creates a file server handler ('fs') that serves static files from the assets directory, the http.Dir function is used to specify the directory from which files should be served

	mux := http.NewServeMux()

	// A new HTTP request multiplexer ('mux') is created usingg http.NewServeMux() this multiplexer will be used to route incoming requests to the appopiate handlers.

	mux.Handle("/assets/", http.StripPrefix("/assets/", fs))

	// This line associates the file server handler ('fs') with the URL path "/assets/", the http.StripPrefix() function is used to strip the "/assets/" prefix from the URL path, so requests for files like "assets/style.css" can be mapped to thje corresponding file in the "assets" directory

	mux.HandleFunc("/", indexHandler)

	// The root URL "/" is associated with the indexHandler function, which serves the HTML template

	http.ListenAndServe(":"+port, mux)

	// Starts the HTTP server and listens on the port specified in the PORT environment variable( or defaults to port 3000), it uses the ('mux') multiplexer to handle incoming requests
}
