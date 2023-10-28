package main

import (
	ascii "ascii/pkg"
	"fmt"
	"log"
	"net/http"
	"text/template"
)

type Error struct {
	Message string
	Code    int
}
type ResponseBody struct {
	Output string
	Error  Error
}

var templates *template.Template

func formHandler(res http.ResponseWriter, req *http.Request) {
	// code for form related stuff
	// Convert input to art and print result

	data := ResponseBody{}

	if req.Method == "POST" {
		input := req.FormValue("input")
		font := req.FormValue("font")

		art, isError, errorCode := ascii.AsciiArtFS(input, font)

		if isError {
			data.Error.Code = errorCode
			data.Error.Message = art
			errorHandler(res, req, &data)
			return
		}

		data.Output = art
	}

	templates.ExecuteTemplate(res, "index.html", data)
}

func serverHandler(res http.ResponseWriter, req *http.Request) {

	data := ResponseBody{}

	if req.URL.Path != "/" {
		data.Error.Code = 404
		data.Error.Message = "Page not found."
		errorHandler(res, req, &data)
		return
	}

	if req.Method == "GET" {

		templates.ExecuteTemplate(res, "index.html", data)

	}

}

func errorHandler(res http.ResponseWriter, req *http.Request, data *ResponseBody) {
	res.WriteHeader(data.Error.Code)
	templates.ExecuteTemplate(res, "error.html", data)
}

func init() {
	templates = template.Must(template.ParseGlob("templates/*.html"))

}

func main() {
	http.HandleFunc("/", serverHandler)
	http.HandleFunc("/ascii-art", formHandler)
	http.Handle("/static/",
		http.StripPrefix("/static/",
			http.FileServer(http.Dir("static"))))
	fmt.Println("Server is listening at Port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
