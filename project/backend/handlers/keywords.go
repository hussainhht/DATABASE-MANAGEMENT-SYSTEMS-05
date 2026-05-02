package handlers

import (
	"database/sql"
	"net/http"
)

func keywordsHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		getKeywords(w, r)
	case http.MethodPost:
		createKeyword(w, r)
	default:
		methodNotAllowed(w)
	}
}

func keywordByIDHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		getKeywordByID(w, r)
	case http.MethodPut:
		updateKeyword(w, r)
	case http.MethodDelete:
		deleteKeyword(w, r)
	default:
		methodNotAllowed(w)
	}
}

func getKeywords(w http.ResponseWriter, r *http.Request) {
	rows, err := DB.Query(`
		SELECT KeywordID, KeywordText
		FROM Keyword
		ORDER BY KeywordID
	`)
	if err != nil {
		sendError(w, http.StatusInternalServerError, err.Error())
		return
	}
	defer rows.Close()

	keywords := []Keyword{}
	for rows.Next() {
		var keyword Keyword
		err := rows.Scan(&keyword.KeywordID, &keyword.KeywordText)
		if err != nil {
			sendError(w, http.StatusInternalServerError, err.Error())
			return
		}

		keywords = append(keywords, keyword)
	}

	if err := rows.Err(); err != nil {
		sendError(w, http.StatusInternalServerError, err.Error())
		return
	}

	sendJSON(w, http.StatusOK, keywords)
}

func getKeywordByID(w http.ResponseWriter, r *http.Request) {
	id, err := parseIDFromPath(r.URL.Path, "/api/keywords/")
	if err != nil {
		sendError(w, http.StatusBadRequest, "invalid keyword id")
		return
	}

	var keyword Keyword
	err = DB.QueryRow(`
		SELECT KeywordID, KeywordText
		FROM Keyword
		WHERE KeywordID = ?
	`, id).Scan(&keyword.KeywordID, &keyword.KeywordText)
	if err == sql.ErrNoRows {
		sendError(w, http.StatusNotFound, "keyword not found")
		return
	}
	if err != nil {
		sendError(w, http.StatusInternalServerError, err.Error())
		return
	}

	sendJSON(w, http.StatusOK, keyword)
}

func createKeyword(w http.ResponseWriter, r *http.Request) {
	var keyword Keyword
	if err := decodeJSONBody(r, &keyword); err != nil {
		sendError(w, http.StatusBadRequest, "invalid JSON body")
		return
	}

	result, err := DB.Exec(`
		INSERT INTO Keyword (KeywordText)
		VALUES (?)
	`, keyword.KeywordText)
	if err != nil {
		sendError(w, http.StatusInternalServerError, err.Error())
		return
	}

	id, err := result.LastInsertId()
	if err != nil {
		sendError(w, http.StatusInternalServerError, err.Error())
		return
	}

	keyword.KeywordID = int(id)
	sendJSON(w, http.StatusCreated, keyword)
}

func updateKeyword(w http.ResponseWriter, r *http.Request) {
	id, err := parseIDFromPath(r.URL.Path, "/api/keywords/")
	if err != nil {
		sendError(w, http.StatusBadRequest, "invalid keyword id")
		return
	}

	var keyword Keyword
	if err := decodeJSONBody(r, &keyword); err != nil {
		sendError(w, http.StatusBadRequest, "invalid JSON body")
		return
	}

	result, err := DB.Exec(`
		UPDATE Keyword
		SET KeywordText = ?
		WHERE KeywordID = ?
	`, keyword.KeywordText, id)
	if err != nil {
		sendError(w, http.StatusInternalServerError, err.Error())
		return
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		sendError(w, http.StatusInternalServerError, err.Error())
		return
	}
	if rowsAffected == 0 {
		sendError(w, http.StatusNotFound, "keyword not found")
		return
	}

	keyword.KeywordID = id
	sendJSON(w, http.StatusOK, keyword)
}

func deleteKeyword(w http.ResponseWriter, r *http.Request) {
	id, err := parseIDFromPath(r.URL.Path, "/api/keywords/")
	if err != nil {
		sendError(w, http.StatusBadRequest, "invalid keyword id")
		return
	}

	result, err := DB.Exec(`
		DELETE FROM Keyword
		WHERE KeywordID = ?
	`, id)
	if err != nil {
		sendError(w, http.StatusInternalServerError, err.Error())
		return
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		sendError(w, http.StatusInternalServerError, err.Error())
		return
	}
	if rowsAffected == 0 {
		sendError(w, http.StatusNotFound, "keyword not found")
		return
	}

	sendJSON(w, http.StatusOK, map[string]string{"message": "keyword deleted successfully"})
}
