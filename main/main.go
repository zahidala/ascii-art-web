package main

import (
	ascii "ascii/pkg"
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type ResponseBody struct {
	Output string
}

func formHandler() {
	// code for form related stuff
}

func serverHandler(res http.ResponseWriter, req *http.Request) {
	// Convert input to art and print result

	data := ResponseBody{}

	art, err := ascii.AsciiArtFS("d", "standard")

	fmt.Println(art)

	data.Output = art

	if err {
		data.Output = art
	}

	templates := template.Must(template.ParseGlob("templates/*.html"))

	templates.ExecuteTemplate(res, "index.html", data)

}

func main() {
	http.HandleFunc("/", serverHandler)
	http.Handle("/static/",
		http.StripPrefix("/static/",
			http.FileServer(http.Dir("static"))))
	fmt.Println("Server is listening at Port 8080...")
	http.ListenAndServe(":8080", nil)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
