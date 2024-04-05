package main

import (
	"flag"
	"log"
	"net/http"
)

func main() {
	addr := flag.String("addr", ":4000", "Сетевой адрес HTTP")

	flag.Parse()

	mux := http.NewServeMux() // инициализация нового рутера

	//fileServer := http.FileServer(neuteredFileSystem{http.Dir("./static/")})
	//mux.Handle("/static", http.NotFoundHandler())
	//mux.Handle("/static/", http.StripPrefix("./static", fileServer))
	mux.HandleFunc("/", home) // "home" регистрируется как обработчик для URL-шаблона "/"
	mux.HandleFunc("/snippet", showSnippet)
	mux.HandleFunc("/snippet/create", createSnippet)

	fileServer := http.FileServer(http.Dir("C:\\Users\\XIaomi\\go\\src\\github.com\\jhufus\\snippetbox\\ui\\static"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	log.Printf("Starting web-server %s", *addr)
	err := http.ListenAndServe(*addr, mux) // запуск нового веб-сервера (TCP-адрес сети для прослушивания, роутер)
	log.Fatal(err)
}

//type neuteredFileSystem struct {
//	fs http.FileSystem
//}
//
//func (nfs neuteredFileSystem) Open(path string) (http.File, error) {
//	f, err := nfs.fs.Open(path)
//	if err != nil {
//		return nil, err
//	}
//
//	s, err := f.Stat()
//	if s.IsDir() {
//		index := filepath.Join(path, "index.html")
//		if _, err := nfs.fs.Open(index); err != nil {
//			closeErr := f.Close()
//			if closeErr != nil {
//				return nil, closeErr
//			}
//		}
//	}
//	return f, nil
//}
