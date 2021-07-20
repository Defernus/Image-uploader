package web

import (
	"encoding/json"
	"net/http"
)

type Error struct {
	Message    string `json:"message"`
	StatusCode int    `json:"status_code"`
}

type Response struct {
	Success bool        `json:"status"`
	Error   *Error      `json:"error"`
	Body    interface{} `json:"body"`
}

func Success(w http.ResponseWriter, body interface{}) {
	result, err := json.Marshal(Response{
		Success: true,
		Body:    body,
	})
	if err != nil {
		panic(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}

func Failed(w http.ResponseWriter, err Error) {
	result, marshalErr := json.Marshal(Response{
		Success: true,
		Error:   &err,
	})
	if marshalErr != nil {
		panic(marshalErr)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(err.StatusCode)
	w.Write(result)
}

func NotFound(w http.ResponseWriter) {
	Failed(w, Error{
		"not_found",
		http.StatusNotFound,
	})
}

func InternalServerError(w http.ResponseWriter) {
	Failed(w, Error{
		StatusCode: http.StatusInternalServerError,
		Message:    "Internal server error",
	})
}
