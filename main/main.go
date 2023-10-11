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

		art, err := ascii.AsciiArtFS(input, font)

		if err {
			data.Error.Code = 500
			data.Error.Message = "Internal Server Error"
		}

		data.Output = art
	}

	templates.ExecuteTemplate(res, "index.html", data)
}

func serverHandler(res http.ResponseWriter, req *http.Request) {

	// art, err := ascii.AsciiArtFS("d", "standard")

	// fmt.Println(art)

	// data.Output = art

	// if err {
	// 	data.Output = art
	// 	data.Error.Code = 500
	// 	data.Error.Message = "Internal Server Error"
	// }

	// templates := template.Must(template.ParseGlob("templates/*.html"))

	data := ResponseBody{}
	templates.ExecuteTemplate(res, "index.html", data)

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
	http.ListenAndServe(":8080", nil)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
