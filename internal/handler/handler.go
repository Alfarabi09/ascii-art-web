package handler

import (
	"asciiart/pkg/ascii"
	"net/http"
	"text/template"
)

func MainPage(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}
	if r.URL.Path != "/" {
		http.Error(w, http.StatusText(404), 404)
		return
	}

	tmpl, err := template.ParseFiles("template/index.html")
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
	}

	if err := tmpl.Execute(w, ""); err != nil {
		http.Error(w, http.StatusText(500), 500)
	}
}

func PostArt(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}
	tmpl, err := template.ParseFiles("template/index.html")
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}
	text := r.FormValue("text")
	if text == "" {
		http.Error(w, http.StatusText(400), 400)
		return
	}
	file := r.FormValue("file")
	ascii, err1 := ascii.Ascii(text, file)
	if err1 != http.StatusOK {
		http.Error(w, ascii, err1)
		return
	}
	if err := tmpl.Execute(w, ascii); err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}
}
