package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

// обработчик главной страницы
func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	files := []string{
		"C:\\Users\\XIaomi\\go\\src\\github.com\\jhufus\\snippetbox\\ui\\html\\home.page.tmpl.html",
		"C:\\Users\\XIaomi\\go\\src\\github.com\\jhufus\\snippetbox\\ui\\html\\base.layout.tmpl.html",
		"C:\\Users\\XIaomi\\go\\src\\github.com\\jhufus\\snippetbox\\ui\\html\\footer.partial.tmpl.html",
	}
	ts, err := template.ParseFiles(files...)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
	}

	err = ts.Execute(w, nil)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
	}
}

// обработчик для отображения содержимого заметки
func showSnippet(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}
	fmt.Fprintf(w, "Displaying the selected mark ID %d...", id)
}

// обработчик для создания новой заметки
func createSnippet(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		http.Error(w, "Method is restricted", 405)
		return
	}
	w.Write([]byte("Form for creating a new snippet..."))
}
