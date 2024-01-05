package main

import (
	"html/template"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	err := openDB()
	if err != nil {
		log.Panic(err)
	}
	defer closeDB()
	err = setupDB()
	if err != nil {
		log.Panic(err)
	}

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Handle("/static/*", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl, err := template.New("").ParseFiles("templates/index.html")
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		err = tmpl.ExecuteTemplate(w, "Base", nil)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	})
	http.ListenAndServe(":3000", r)
}
