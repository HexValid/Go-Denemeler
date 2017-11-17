package main

import (
	"html/template"
	"net/http"
)

type Kedi struct {
	Yasi, Kilo int
	Cinsi      string
}

func Roothandler(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("ders3.html")
	pisi := Kedi{
		Yasi:  5,
		Kilo:  2,
		Cinsi: "Van Kedisi",
	}
	t.Execute(w, pisi)
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/", Roothandler)
	server.ListenAndServe()
}
