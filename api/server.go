package api

import (
	"net/http"

	"github.com/hacdan/receipt-processor-challenge/storage"
)

type Server struct {
	listenAddr string
	storage    storage.Storage
}

func NewServer(listenAddr string) Server {
	storage := storage.NewStorage()

	return Server{
		listenAddr: listenAddr,
		storage:    storage,
	}
}

func (s *Server) Start() error {
	http.HandleFunc("GET /receipts/{id}/points", s.HandleGetPoints)
	http.HandleFunc("POST /receipts/process", s.HandlePostReceipt)

	return http.ListenAndServe(s.listenAddr, nil)
}
