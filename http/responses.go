package http

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

func writeJSON(w http.ResponseWriter, statusCode int, data any) {
	body, err := json.MarshalIndent(data, "", "	")
	if err != nil {
		fmt.Println("failed to marshal HTTP response:", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "applicatoin/json")
	w.WriteHeader(statusCode)
	if _, err := w.Write(body); err != nil {
		fmt.Println("failed to write HTTP response:", err)
	}
}

func writeError(w http.ResponseWriter, statusCode int, err error) {
	writeJSON(w, statusCode, NewErrDTO(err.Error(), time.Now()))
}
