package web

import (
	"net/http"
	"time"

	"github.com/sirupsen/logrus"
)

func (s *Server) LogRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		s.log.WithFields(logrus.Fields{"method": r.Method, "url": r.URL.Path}).Debug("start request")
		respWriter := newResponseWriter(w)
		now := time.Now()
		next.ServeHTTP(respWriter, r)
		since := time.Since(now)

		s.log.WithFields(logrus.Fields{
			"time":           since,
			"status_code":    respWriter.statusCode,
			"content_length": respWriter.contentLength,
		}).Debug("finish request")
	})
}

type responseWriter struct {
	http.ResponseWriter
	statusCode    int
	contentLength int
}

func newResponseWriter(w http.ResponseWriter) *responseWriter {
	return &responseWriter{
		ResponseWriter: w,
		statusCode:     http.StatusOK,
		contentLength:  0,
	}
}

func (w *responseWriter) WriteHeader(statusCode int) {
	w.statusCode = statusCode
	w.ResponseWriter.WriteHeader(statusCode)
}

func (w *responseWriter) Write(data []byte) (int, error) {
	w.contentLength += len(data)
	return w.ResponseWriter.Write(data)
}
