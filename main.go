package main

import (
	"ascii-art-web/art"
	datas "ascii-art-web/data"
	"ascii-art-web/errors"
	"fmt"
	"html/template"
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

var tmpl *template.Template
var data datas.PageData

func HandleAscii(w http.ResponseWriter, r *http.Request) {

	data.Title = "Bienvenue sur ASCII ART WEB"
	data.Paragraph = "Écris du texte, choisis la style de police, push et regarde le resultat"
	data.Response = ""
	data.Input = ""
	data.Standard = "checked"
	data.Shadow = ""
	data.ThinkerToy = ""
	data.ButtonLoad = false

	switch r.Method {

	case "POST":
		if err := r.ParseForm(); err != nil {
			errors.Check_500_Error(w, r)
			return
		}
	}

	if r.URL.Path == "/ascii-art" || r.URL.Path == "/" {
		text := r.Form.Get("input")

		file := r.FormValue("bouton")
		if file != "shadow.txt" && file != "standard.txt" && file != "thinkertoy.txt" && file != "" {
			errors.Check_500_Error(w, r)
			return
		}

		for _, e := range text { // Text étant l'input du client
			// Si e n'est pas une valeur ascii && n'est pas espace / entrée return erreur bad request
			if (e < 32 || e > 127) && e != 13 && e != 10 {
				errors.Check_400_Error(w, r)
				return
			}
		}

		if file != "" {
			data.Response = art.Ascii_Art(text, "txtfiles/"+file)
			data.Input = text

			if file == "standard.txt" {
				data.Standard = "checked"
				data.Shadow = ""
				data.ThinkerToy = ""
			} else if file == "shadow.txt" {
				data.Standard = ""
				data.Shadow = "checked"
				data.ThinkerToy = ""
			} else {
				data.Standard = ""
				data.Shadow = ""
				data.ThinkerToy = "checked"
			}
			data.ButtonLoad = true
		}
		extension := r.FormValue("format")

		data.Download = extension

		tmpl.Execute(w, data)
	} else {
		errors.Check_404_Error(w, r)
	}
}

func Downloadfile(w http.ResponseWriter, r *http.Request) {

	filePath := "document/fichier." + data.Download

	err := ioutil.WriteFile(filePath, []byte(data.Response), 0644)
	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Disposition", "attachment; filename=fichier."+data.Download)
	w.Header().Set("Content-Type", "text/plain")

	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}

	_, err = io.Copy(w, file)
	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Length", fmt.Sprint(r.ContentLength))
}

func main() {
	fmt.Println("starting at : http://localhost:8070/")

	tmpl = template.Must(template.ParseFiles("templates/index.html"))

	http.HandleFunc("/", HandleAscii)
	http.HandleFunc("/document/", Downloadfile)

	http.ListenAndServe(":8070", nil)
}
