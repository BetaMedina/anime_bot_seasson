package utils

import (
	"encoding/json"
	"net/http"
)

func HttpResponse(w http.ResponseWriter, statusCode int, payload map[string]interface{}) {
	w.WriteHeader(statusCode)
	_ = json.NewEncoder(w).Encode(payload)
	return
}
