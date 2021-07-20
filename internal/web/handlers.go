package web

import (
	"net/http"

	"imageapi.lavrentev.dev/rest/internal/database"
)

func (s *Server) handleGetImages(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	images := []database.Image{}
	err := s.db.DB.Find(&images).Error
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(internalError))
	}

	if err := Success(w, images); err != nil {
		panic(err)
	}
}
