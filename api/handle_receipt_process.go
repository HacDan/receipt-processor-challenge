package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/hacdan/receipt-processor-challenge/types"
)

func (s *Server) HandlePostReceipt(w http.ResponseWriter, r *http.Request) {
	receipt := types.Receipt{}

	err := json.NewDecoder(r.Body).Decode(&receipt)
	if err != nil {
		fmt.Printf("Error parsing JSON with error:  %s\n", err.Error())
		respError(w, "Error parsing json", http.StatusUnprocessableEntity)
		return
	}

	id := s.storage.AddReceipt(receipt)
	fmt.Printf("Added receipt with id: %s\n", id.Id)

	respondJSON(w, id, http.StatusOK)
}
