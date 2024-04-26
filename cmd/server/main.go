package main

import (
	"github.com/michael-martinez-dev/globalwarfront-server/internal/api"

	"net/http"

	log "github.com/sirupsen/logrus"
	prefixed "github.com/x-cray/logrus-prefixed-formatter"
)

func main() {
	log.SetFormatter(new(prefixed.TextFormatter))
	router := api.NewRouter()
	log.Infoln("Server starting on :8000...")
	err := http.ListenAndServe(":8000", router)
	if err != nil {
		log.Warn(err)
	}
}
