package handler

import (
	"context"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"sorting-service/internal/service"
)

type HTTPHandler struct {
	service service.NumberService
}

func NewHTTPHandler(svc service.NumberService) *HTTPHandler {
	return &HTTPHandler{service: svc}
}

func (h *HTTPHandler) AddNumber(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "failed read body", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	var payload struct {
		Value *int `json:"value"`
	}

	if err := json.Unmarshal(body, &payload); err == nil && payload.Value != nil {
		h.callServiceAndRespond(w, r, *payload.Value)
		return
	}

	s := strings.TrimSpace(string(body))
	if s == "" {
		http.Error(w, "empty body", http.StatusBadRequest)
		return
	}
	n, err := strconv.Atoi(s)
	if err != nil {
		http.Error(w, "invalid number", http.StatusBadRequest)
		return
	}
	h.callServiceAndRespond(w, r, n)
}

func (h *HTTPHandler) callServiceAndRespond(w http.ResponseWriter, r *http.Request, n int) {
	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	defer cancel()

	arr, err := h.service.AddNumber(ctx, n)
	if err != nil {
		log.Printf("service error: %v", err)
		http.Error(w, "internal error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(arr); err != nil {
		log.Printf("encode response: %v", err)
	}
}
