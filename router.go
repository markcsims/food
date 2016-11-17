package main

import (
	"fmt"
	"net/http"

	"github.com/mergermarket/gotools"
	"strings"
)

// Hello greets a given name
func Hello(name string) string {
	return fmt.Sprintf("Hello, %s", name)
}

type logger interface {
	Info(...interface{})
	Debug(...interface{})
}

// HelloHandler is a http handler which says hello to you
type HelloHandler struct {
	log logger
}

func (hr *HelloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	hr.log.Info("oi!")
	fmt.Fprint(w, Hello(r.URL.Query().Get("name")))
}

type SearchHandler struct {
	log logger
}

func (hr *SearchHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	hr.log.Info("search SearchHandler!")
	searchTerm := r.URL.Query().Get("keywords")
	searchTerm = strings.Join(strings.Split(searchTerm, " "), "+")
	if searchTerm == "" {
		fmt.Fprint(w, "Please enter a search term with keywords query string")
	} else {
		fmt.Fprint(w, recipeSearchScraper(searchTerm))
	}
}

func newRouter(config *appConfig, log logger) http.Handler {
	router := http.NewServeMux()
	logConfigHandler := tools.NewInternalLogConfig(config, log)
	router.HandleFunc("/internal/healthcheck", tools.InternalHealthCheck)
	router.HandleFunc("/internal/log-config", logConfigHandler)
	router.Handle("/hello", &HelloHandler{log})
	router.Handle("/search", &SearchHandler{log})
	return router
}
