package server

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/CESARBR/knot-gateway-webui/backend/pkg/logging"
	"github.com/gorilla/mux"
)

// Start starts the http server
func Start() {
	r := createRouters()
	logger := logging.Get("Server")
	logger.Info("Listening on " + strconv.Itoa(8080))
	http.ListenAndServe(":8080", r)
}

func createRouters() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/healthcheck", healthcheckHandler)
	return r
}

func healthcheckHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "online")
}
