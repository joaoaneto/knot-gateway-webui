package main

import (
	"github.com/CESARBR/knot-gateway-webui/backend/pkg/server"
	log "github.com/sirupsen/logrus"
)

func main() {
	log.Info("Starting KNoT Gateway WebUI Backend")
	server.Start()
}
