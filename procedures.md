## Creating a web server

```go
func main() {
	port := os.Getenv("PORT")
 	if port == "" {
  		port = "3000"
    }

    mux := http.NewServeMux()

    mux.HandleFunc("/", indexHandler)
    http.ListenAndServe(":"+port, mux)
}
```

In the main() function, an attempt is made to set the **port** variable based on the value of the **PORT** environment variable.

If the variable is not present, **getenv** returns an empty sdtring and portt is set to 3000 so that the server is made available at _http://localhost:3000_

```go
	mux := http.NewServeMux()

	mux.HandleFunc("/", indexHandler)
	http.ListenAndServe(":"+port, mux)
```

The ** http.NewServeMux() ** method is used to create an HTTP request multiplexer which is subsequently assigned to the mux variable.

An HTTP request multiplexer is a software component/service that is responsible for handling and routing **incoming HTTP requests to the appropiate handlers or endpoints within a web application**

They organize and manage the flow of incoming requests efficiently.

In this case, it will match the URL of incoming requests against a list of registered patterns, and calls the associated handler for the pattern whenever a match is found.

Registering HTTP request handlers is done via the HandleFunc method which takes the pattern string as its first argument, and a function with the following signature

```go
	func (w http.ResponseWriter, r *http.Request)
```

Regarding the indexHandler function...

```go
	func indexHandler(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("<h1>Hello World!</h1>")
	}
```

IndexHandler is defiend to handle requests made to the root URL ("/"), it writes an HTML "Hello World!" message inside an **h1** element to the HTTP response.

### Reading variables from the environment

A common pattern regarding environmental variables is to load them from a .env file into the environment, the first thing we need to create is a **.env** file in the root of our project directory and open it

```bash
touch .env
nano .env
```

Then we can set the PORT environmental variable within the file

```
PORT = 300
```

After that we need to update our main.go as shown here

```c
	err := godotenv.Load()
 	if err != nil {
  		log.Println("Error loading .env file")
    }
```

This Load method reads the .env file and loads the set variables into the environment so that they can be accessed through the os.Getenv() method, helpful to store secred credentials in the environment.

### Templating in Go

Templates provide an easy way to customize the output of your web application depending on the route whithout having to write the same code in many places, for example, we can create a template for the navigation bar and use it across al pages of the site without duplicating the code.

GO provides two template packages in its stasndard library, **text/template** and **html/template**, they provide the same interface but **html/template** package is used to generate output that is safe against code injection.

```go
package main

import (
	"html/template"
 	"log"
  	"net/http"
   	"os"

    "github.com/joho/godotenv"
)

var tpl = template.Must(tempalte.ParseFiles("index.html"))

func indexHandler(w http.ResponseWriter, r *http.Request) {
	tpl.Execute(w, nil)
}
```

**tpl** is a package level variable that points to a tempalte definition from the provided files.

The call to template.Parsefiles is wrapped with template.Must so that the coded panics if an error obtained while parsing the template file.

The reason we panic here instead of trying to handle the error, is because a web app with a broken template is not much of a web app, it's a problem that should be fixed before attempting to restart the server.

In the **indexHandler** function, the **tpl** template is executed by providing two arguments.

- Where we want to write the output to
- The data that we want to pass to the template

```go
func indexHandler(w http.ResponseWriter, r *http.Request) {
	tpl.Execute(w, nil)
}
```

In the above case, we're writing the output to the ResponseWriter interface and, since we don't have any data to pass to our template, nil is passed as the second argument.

**Build snippet**

```go
go build && ./news-demo-starter-files
```

### Automatically restarting the server

##### Using Air to restart the server

Can be a pain in the ass to build and restart the server every time we make a change in the code, we can avoid that using the **Air package**

```go
go install github.com/cosmtrek/air@latest
```

Then run the air command at the root of your project directory:

```go
air
```

### Serving static files

Now that the HTML file for the navigation file has been added, it remains unstyled since assets/style.css file is linked correctly in the **head** of our document, however it doesnt show up.

This is because **we haven't registered the /assets pattern in the HTTP multiplexer.**

We need to make sure that all requests that  match this pattern are served as static files.

The first thing to do is to **instantiate a file server object by passing the directory where all our static files are placed**

```go
fs := http.FileServer(http.Dir("assets"))
```

Next we need to tell our router to use this file server object for all paths beginning with the /assets/ prefix

```go
mux.Handle("/assets/", http.StripPrefix("/assets/", fs))
```

The **http.StripPrefix()** method modifies the request URL by stripping off the specified prefix before forwaring the handling of the request to the http.Handler in the second parameter.

