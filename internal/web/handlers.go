package web

import (
	"net/http"

	"github.com/gorilla/mux"
	"imageapi.lavrentev.dev/rest/internal/database"
)

func (s *Server) handleGetImages(w http.ResponseWriter, r *http.Request) {
	images := []database.Image{}
	err := s.db.DB.Find(&images).Error
	if err != nil {
		InternalServerError(w)
	}

	Success(w, images)
}

func (s *Server) handleDeleteImage(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	err := s.db.DB.Delete(&database.Image{}, vars["id"]).Error
	if err != nil {
		InternalServerError(w)
	}

	Success(w, nil)
}
