package main

import (
	log "github.com/sirupsen/logrus"
	"github.com/unravellingtechnologies/elastic-apm-operator/cli"
	"github.com/unravellingtechnologies/elastic-apm-operator/server"
)

var (
	port            int
	tlsCert, tlsKey string
	logLevel        log.Level
)

// init function is called before main and initialises the needed configurations
func init() {
	// Initialises CLI
	cli.Init()
	// Parse CLI options
	cli.Parse(&port, &tlsCert, &tlsKey, &logLevel)

	// Initialize logging
	log.SetFormatter(&log.JSONFormatter{})
	log.SetLevel(logLevel)
}

// main function is the entrypoint of our application
func main() {
	log.Debugf("Read config - Port: %d, Certificate: %s, Key: %s", port, tlsCert, tlsKey)

	server.Start(port, tlsCert, tlsKey)
}
