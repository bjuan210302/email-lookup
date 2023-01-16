package main

import (
	"elookup/wrapper"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/render"
)

func main() {
	r := chi.NewRouter()

	r.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "HEAD", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"*"},
		ExposedHeaders: []string{"Link"},
	}))

	r.Use(middleware.Logger)
	r.Use(middleware.RequestID)
	r.Use(middleware.Recoverer)
	r.Use(middleware.URLFormat)
	r.Use(render.SetContentType(render.ContentTypeJSON))

	r.Route("/api/v1", func(r chi.Router) {
		r.Get("/lookup", searchWord)
	})

	log.Println("Serving on port 3000")
	http.ListenAndServe(":3000", r)
}

func searchWord(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	word := query.Get("word")
	page, _ := strconv.Atoi(query.Get("page"))
	maxResults, _ := strconv.Atoi(query.Get("max_per_page"))

	queryResult := wrapper.SearchByWord(word, page, maxResults)
	render.JSON(w, r, queryResult)
}