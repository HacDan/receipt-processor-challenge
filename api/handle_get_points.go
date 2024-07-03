package api

import (
	"fmt"
	"net/http"

	"github.com/hacdan/receipt-processor-challenge/types"
)

func (s *Server) HandleGetPoints(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	points, err := s.storage.GetPoints(id)
	if err != nil {
		respError(w, err.Error(), http.StatusNotFound)
		return
	}
	fmt.Printf("Total points for receipt %s: %d\n", id, points)

	responsePoints := types.Points{
		Points: points,
	}

	respondJSON(w, responsePoints, http.StatusOK)
}
