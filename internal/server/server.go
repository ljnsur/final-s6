package server

import (
	"log"
	"net/http"
	"s6-final/internal/handlers"
	"time"
)

type Server struct {
	logger     *log.Logger
	httpServer *http.Server
}

func New(logger *log.Logger) *Server {
	mux := http.NewServeMux()

	mux.HandleFunc("/", handlers.MainHandler)
	mux.HandleFunc("/upload", handlers.UploadHandler)

	config := &http.Server{
		Addr:         ":8080",
		Handler:      mux,
		ErrorLog:     logger,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}
	return &Server{
		logger:     logger,
		httpServer: config,
	}

}
func (s *Server) Start() error {
	s.logger.Printf("сервер запущен напорту %s", s.httpServer.Addr)
	return s.httpServer.ListenAndServe()
}
