package server

import (
	"context"
	log "github.com/sirupsen/logrus"
	"github.com/unravellingtechnologies/elastic-apm-operator/handlers"
	zert "github.com/unravellingtechnologies/zert"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
)

func Start(port int, tlsCert string, tlsKey string) {
	serverTLSConf, _, err := zert.TLSSetup(tlsCert, tlsKey)
	if err != nil {
		panic(err)
	}

	mux := handlers.Setup()

	s := &http.Server{
		Addr:      ":" + strconv.Itoa(port),
		Handler:   mux,
		TLSConfig: serverTLSConf,
	}

	go func() {
		if err := s.ListenAndServeTLS("", ""); err != nil {
			log.Errorf("Failed to listen and serve: %v", err)
		}
	}()

	log.Info("Listening for requests on port ", port)

	// listen shutdown signal
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)
	<-signalChan

	log.Infof("Shutdown gracefully...")
	if err := s.Shutdown(context.Background()); err != nil {
		log.Error(err)
	}
}
