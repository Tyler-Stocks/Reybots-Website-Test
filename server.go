package main

import (
	// STD imports
	"html/template"
	"log"
	"net/http"
	"fmt"
	"io"
	"time"

	// Third party imports
	"github.com/Tyler-Stocks/Reybots-Website-Test/context/Context"
)

const STATIC_URL string = "/static/"
const STATIC_ROOT string = "static/"

func Home(writer http.ResponseWriter, request *http.Request) {
	context := Context{Title: "Home"}
	render(writer, "home", context)	
}

func Sponsors(writer http.ResponseWriter , request *http.Request) {
	context := Context{Title: "Sponsors"}
	render(writer, "sponsors", context)
}

func About(writer http.ResponseWriter, request *http.Request) {
	context := Context{Title: "About"}
	render(writer, "about", context)
}

func Competitions(writer http.ResponseWriter, request * http.Request) {
	context := Context{Title: "Competitions"}
	render(writer, "about", context)
}

func render(writer http.ResponseWriter, templateName string, context Context) {
	context.Static = STATIC_URL
(
	templatePathList := []string{"templates/base.html", fmt.Sprintf("templates/%s.html", templateName)}

	template, error := template.ParseFiles(templatePathList...)
	
	if error != nil {
		log.Print("Render: Template Parsing Error " + error.Error())
	}

	log.Print("Context Title: " + context.Title)
	error = template.Execute(writer, context)

	if error != nil {
		log.Print("Renderer: Template Executing Error " + error.Error()) 
	}
}

func StaticHandler(writer http.ResponseWriter, request *http.Request) {
	static_file := request.URL.Path[len(STATIC_URL):]

	if len(static_file) == 0 {
		http.NotFound(writer, request)
		return
	}

	file, error := http.Dir(STATIC_ROOT).Open(static_file)

 	if error != nil {
		http.NotFound(writer, request)
		return
	}

	content := io.ReadSeeker(file)
	http.ServeContent(writer, request, static_file, time.Now(), content)
}

func main() {
	http.HandleFunc("/sponsors/", Sponsors)
	// This HandleFunc has to happen after all the other paths or it just *doesn't work* for some reason
	http.HandleFunc("/", Home) 
	http.HandleFunc(STATIC_URL, StaticHandler)
	http.NotFoundHandler()
	error := http.ListenAndServe(":8080", nil)

	if error != nil {
		log.Fatal("ListenAndServe: " + error.Error())
	}
}


