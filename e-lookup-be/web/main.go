package main

import (
	"elookup/wrapper"
	"flag"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/render"
)

var (
	port *string
)

func init() {
	port = flag.String("port", "3000", "port number")
	flag.Parse()
}

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
		r.Get("/indexes", allIndexNames)
		r.Get("/ping", ping)
	})

	log.Printf("Serving on port %s", *port)
	http.ListenAndServe(":"+*port, r)
}

func searchWord(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	word := query.Get("word")
	page, _ := strconv.Atoi(query.Get("page"))
	maxResults, _ := strconv.Atoi(query.Get("max_per_page"))
	indexName := query.Get("index_name")

	auth := r.Header.Get("Authorization")

	queryResult := wrapper.SearchByWord(word, page, maxResults, indexName, auth)
	render.JSON(w, r, queryResult)
}

func allIndexNames(w http.ResponseWriter, r *http.Request) {
	auth := r.Header.Get("Authorization")

	names := wrapper.GetIndexNamesList(auth)
	render.JSON(w, r, names)
}

func ping(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Pong!"))
}
