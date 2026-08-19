package http

import (
	"encoding/json"
	"net/http"
)

func SetJsonResponse(rw http.ResponseWriter, data any) {
	rw.Header().Set("Content-Type", "application/json")
	json.NewEncoder(rw).Encode(data)
}
