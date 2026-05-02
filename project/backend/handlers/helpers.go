package handlers

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"strings"
)

func sendJSON(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

func sendError(w http.ResponseWriter, status int, message string) {
	sendJSON(w, status, map[string]string{
		"error": message,
	})
}

func methodNotAllowed(w http.ResponseWriter) {
	sendError(w, http.StatusMethodNotAllowed, "method not allowed")
}

func decodeJSONBody(r *http.Request, data any) error {
	decoder := json.NewDecoder(r.Body)
	return decoder.Decode(data)
}

func parseIDFromPath(path string, prefix string) (int, error) {
	if !strings.HasPrefix(path, prefix) {
		return 0, errors.New("invalid path")
	}

	idText := strings.Trim(strings.TrimPrefix(path, prefix), "/")
	if idText == "" || strings.Contains(idText, "/") {
		return 0, errors.New("invalid id")
	}

	id, err := strconv.Atoi(idText)
	if err != nil || id <= 0 {
		return 0, errors.New("invalid id")
	}

	return id, nil
}

func parseTwoIDsFromPath(path string, prefix string) (int, int, error) {
	if !strings.HasPrefix(path, prefix) {
		return 0, 0, errors.New("invalid path")
	}

	idText := strings.Trim(strings.TrimPrefix(path, prefix), "/")
	parts := strings.Split(idText, "/")
	if len(parts) != 2 {
		return 0, 0, errors.New("invalid ids")
	}

	firstID, err := strconv.Atoi(parts[0])
	if err != nil || firstID <= 0 {
		return 0, 0, errors.New("invalid first id")
	}

	secondID, err := strconv.Atoi(parts[1])
	if err != nil || secondID <= 0 {
		return 0, 0, errors.New("invalid second id")
	}

	return firstID, secondID, nil
}

func nullIfEmpty(value string) any {
	if strings.TrimSpace(value) == "" {
		return nil
	}

	return value
}
