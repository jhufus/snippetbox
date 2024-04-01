package main

import (
	"fmt"
	"net/http"
	"strconv"
)

// обработчик главной страницы
func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	w.Write([]byte("Hello from Snippetbox"))
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
		w.Header()["Data"] = nil
		http.Error(w, "Method is restricted", 405)
		return
	}
	w.Write([]byte("Form for creating a new snippet..."))
}
