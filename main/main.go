package main

import (
	"fmt"
	"html/template"
	"net/http"
	"time"
)

type Welcome struct {
	Name string
	Time string
}

func main() {
	// // Handle args error
	// err := ascii.ArgsErrors()
	// if err {
	// 	return
	// }

	// // Convert input to art and print result
	// art := ascii.AsciiArtFS(os.Args[1])
	// fmt.Print(art)

	welcome := Welcome{"Anonymous", time.Now().Format(time.Stamp)}

	templates := template.Must(template.ParseFiles("templates/template.html"))

	http.Handle("/static/", // final URL can be anything
		http.StripPrefix("/static/",
			http.FileServer(http.Dir("static"))))

	http.HandleFunc("/ascii", func(w http.ResponseWriter, r *http.Request) {
		// if have name from http request r -> /?name=your-name-here -> name
		if name := r.FormValue("name"); name != "" {
			welcome.Name = name // set struct object to hold name from URL
		}
		// if have error
		if err := templates.ExecuteTemplate(w, "template.html", welcome); err != nil {
			// show error message
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	fmt.Println(http.ListenAndServe(":8080", nil))
}
