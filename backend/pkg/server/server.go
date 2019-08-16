package server

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/CESARBR/knot-gateway-webui/backend/pkg/interactors"
	"github.com/CESARBR/knot-gateway-webui/backend/pkg/logging"
	"github.com/CESARBR/knot-gateway-webui/backend/pkg/services"

	"github.com/gorilla/mux"
)

// Health represents the service's health status
type Health struct {
	Status string `json:"status"`
}

// Server represents the HTTP server
type Server struct {
	port         int
	stateService *services.StateService
}

// NewServer creates a new server instance
func NewServer(port int) Server {
	updateStateInteractor := interactors.NewUpdateStateInteractor(nil)
	stateService := services.NewStateService(updateStateInteractor)
	return Server{port: port, stateService: stateService}
}

// Start starts the http server
func (s *Server) Start() {
	routers := s.createRouters()
	logger := logging.Get("Server")

	logger.Info(fmt.Sprintf("Listening on %d", s.port))
	err := http.ListenAndServe(fmt.Sprintf(":%d", s.port), logRequest(routers))
	if err != nil {
		logger.Error(err)
	}
}

func (s *Server) createRouters() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/healthcheck", s.healthcheckHandler)
	r.HandleFunc("/state", s.updateStateHandler).Methods("PUT")
	return r
}

func logRequest(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger := logging.Get("Server")
		logger.Info(fmt.Sprintf("%s %s %s\n", r.RemoteAddr, r.Method, r.URL))
		handler.ServeHTTP(w, r)
	})
}

func (s *Server) healthcheckHandler(w http.ResponseWriter, r *http.Request) {
	logger := logging.Get("Server")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	response, _ := json.Marshal(&Health{Status: "online"})
	_, err := w.Write(response)
	if err != nil {
		logger.Error(fmt.Sprintf("Error sending response, %s\n", err))
	}
}

type StateData struct {
	Type string `json:type`
}

func (s *Server) updateStateHandler(w http.ResponseWriter, r *http.Request) {
	// logger := logging.Get("Server")
	// _, err := w.Write(response)
	// if err != nil {
	// 	logger.Error(fmt.Sprintf("Error sending response, %s\n", err))
	// }

	decoder := json.NewDecoder(r.Body)
	var stateData StateData
	err := decoder.Decode(&stateData)
	if err != nil {
		panic(err)
	}

	s.stateService.UpdateStateInteractor.Execute(stateData.Type)
}
