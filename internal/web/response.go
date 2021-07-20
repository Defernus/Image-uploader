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

func Success(w http.ResponseWriter, body interface{}) error {
	result, err := json.Marshal(Response{
		Success: true,
		Body:    body,
	})
	if err != nil {
		return err
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(result)
	return nil
}

func Failed(w http.ResponseWriter, err Error) error {
	result, marshalErr := json.Marshal(Response{
		Success: true,
		Error:   &err,
	})
	if marshalErr != nil {
		return marshalErr
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(err.StatusCode)
	w.Write(result)
	return nil
}

func NotFound(w http.ResponseWriter) error {
	return Failed(w, Error{
		"not_found",
		http.StatusNotFound,
	})
}
