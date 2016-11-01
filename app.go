package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/mergermarket/gotools"
)

func main() {

	config, err := loadAppConfig()
	if err != nil {
		log.Fatal("Error loading config", err.Error())
	}

	logger := tools.NewLogger(config.IsLocal())
	logger.Info(fmt.Sprintf("Application config - %+v", config))
	recipeScraper("http://www.bbc.co.uk/food/recipes/rackoflambwithsmoked_90893")

	logger.Info("Listening on", config.Port)
	err = http.ListenAndServe(fmt.Sprintf(":%d", config.Port), newRouter(config, logger))

	if err != nil {
		logger.Error("Problem starting server", err.Error())
		os.Exit(1)
	}
}
