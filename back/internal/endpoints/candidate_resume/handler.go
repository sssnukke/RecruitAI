package candidate_resume

import (
	"encoding/json"
	"net/http"
)

type Handler struct {
	service *Service
}

func NewHandler(service *Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) CheckCandidate(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	if err := r.ParseMultipartForm(10 << 20); err != nil {
		http.Error(w, "invalid multipart form", http.StatusBadRequest)
		return
	}

	vacancy := r.FormValue("vacancy")
	if vacancy == "" {
		http.Error(w, "vacancy text is required", http.StatusBadRequest)
		return
	}

	file, header, err := r.FormFile("resume")
	if err != nil {
		http.Error(w, "resume file is required", http.StatusBadRequest)
		return
	}
	defer file.Close()

	if header.Size == 0 {
		http.Error(w, "empty resume file", http.StatusBadRequest)
		return
	}

	match, err := h.service.CheckCandidate(file, header.Filename, vacancy)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]bool{
		"match": match,
	})
}
