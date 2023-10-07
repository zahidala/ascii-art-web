package main

import (
	ascii "ascii/pkg"
	"fmt"
	"html/template"
	"log"
	"net/http"
)

// type Welcome struct {
// 	Name string
// 	Time string
// }

type ResponseBody struct {
	Output string
}

func formHandler() {
	// code for form related stuff
}

func serverHandler(res http.ResponseWriter, req *http.Request) {
	// Convert input to art and print result

	data := ResponseBody{}

	art, err := ascii.AsciiArtFS("d")

	fmt.Println(art)

	data.Output = art

	if err {
		data.Output = art
	}

	// welcome := Welcome{"Anonymous", time.Now().Format(time.Stamp)}

	templates := template.Must(template.ParseFiles("templates/template.html"))

	http.Handle("/static/", // final URL can be anything
		http.StripPrefix("/static/",
			http.FileServer(http.Dir("static"))))

	http.HandleFunc("/ascii-art", func(w http.ResponseWriter, r *http.Request) {
		// if have name from http request r -> /?name=your-name-here -> name
		// if name := r.FormValue("name"); name != "" {
		// welcome.Name = name // set struct object to hold name from URL
		// }
		// if have error
		if err := templates.ExecuteTemplate(w, "template.html", data); err != nil {
			// show error message
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

	})

}

func main() {
	http.HandleFunc("/", serverHandler)
	fmt.Println("Server is listening at Port 8080...")
	http.ListenAndServe(":8080", nil)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
