package api

import (
	"encoding/json"
	"log/slog"
	"net/http"
)

type Response struct {
	Error string `json:"error,omitempty"`
	Data  any    `json:"data,omitempty"`
}

func sendJson(w http.ResponseWriter, resp Response, status int) {
	data, err := json.Marshal(resp)
	if err != nil {
		slog.Error("Failed to marchal response", "err", err)
		sendJson(
			w,
			Response{
				Error: "Something went wrong",
			},
			http.StatusInternalServerError,
		)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if _, err := w.Write(data); err != nil {
		slog.Error("failed to write a response to client", "error", err)
		return
	}
}
