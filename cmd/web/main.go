package main

import (
	"log"
	"net/http"
)

func main() {

	mux := http.NewServeMux() // инициализация нового рутера
	mux.HandleFunc("/", home) // "home" регистрируется как обработчик для URL-шаблона "/"
	mux.HandleFunc("/snippet", showSnippet)
	mux.HandleFunc("/snippet/create", createSnippet)

	fileServer := http.FileServer(http.Dir("C:\\Users\\XIaomi\\go\\src\\github.com\\jhufus\\snippetbox\\ui\\static"))

	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	log.Println("Starting web-server http://127.0.0.1:4000")
	err := http.ListenAndServe(":4000", mux) // запуск нового веб-сервера (TCP-адрес сети для прослушивания, роутер)
	log.Fatal(err)
}
