package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/markcsims/bbcfood/tools"
)

func main() {

	config, err := loadAppConfig()
	if err != nil {
		log.Fatal("Error loading config", err.Error())
	}

	logger := tools.NewLogger(config.IsLocal())
	logger.Info(fmt.Sprintf("Application config - %+v", config))

	statsdConfig := tools.StatsDConfig{IsProduction: !config.IsLocal(), Log: logger}
	statsd, err := tools.NewStatsD(statsdConfig)
	if err != nil {
		logger.Error("Error connecting to StatsD - defaulting to logging stats: ", err.Error())
	}

	logger.Info("Listening on", config.Port)
	err = http.ListenAndServe(fmt.Sprintf(":%d", config.Port), newRouter(config, logger, statsd))

	if err != nil {
		logger.Error("Problem starting server", err.Error())
		os.Exit(1)
	}
}
