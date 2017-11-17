package main

import (
	"html/template"
	"net/http"
)

func anaSayfaDinleyici(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("ders4.html")
	t.Execute(w, nil)
}

func formDinleyici(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("ders4.html")
	adsoyad := r.FormValue("firstname") + " " + r.FormValue("lastname")
	t.Execute(w, adsoyad)
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/", anaSayfaDinleyici)
	http.HandleFunc("/postadresi", formDinleyici)
	server.ListenAndServe()
}
