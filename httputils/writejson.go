package httputils

import "net/http"

// WriteJSON ...
func WriteJSON(w http.ResponseWriter, v []byte) {
	w.Header().Set("Content-Type", "application/json")
	w.Write(v)
}
