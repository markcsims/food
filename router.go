package main

import (
	"fmt"
	"net/http"

	"github.com/mergermarket/gotools"
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

func newRouter(config *appConfig, log logger) http.Handler {
	router := http.NewServeMux()
	logConfigHandler := tools.NewInternalLogConfig(config, log)
	router.HandleFunc("/internal/healthcheck", tools.InternalHealthCheck)
	router.HandleFunc("/internal/log-config", logConfigHandler)
	router.Handle("/hello", &HelloHandler{log})
	return router
}
