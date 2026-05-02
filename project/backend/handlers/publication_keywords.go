package handlers

import "net/http"

func publicationKeywordsHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		getPublicationKeywords(w, r)
	case http.MethodPost:
		createPublicationKeyword(w, r)
	default:
		methodNotAllowed(w)
	}
}

func publicationKeywordByIDsHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodDelete:
		deletePublicationKeyword(w, r)
	default:
		methodNotAllowed(w)
	}
}

func getPublicationKeywords(w http.ResponseWriter, r *http.Request) {
	rows, err := DB.Query(`
		SELECT
			pk.PublicationID,
			pk.KeywordID,
			p.Title,
			k.KeywordText
		FROM PublicationKeyword pk, Publication p, Keyword k
		WHERE pk.PublicationID = p.PublicationID
		  AND pk.KeywordID = k.KeywordID
		ORDER BY pk.PublicationID, pk.KeywordID
	`)
	if err != nil {
		sendError(w, http.StatusInternalServerError, err.Error())
		return
	}
	defer rows.Close()

	publicationKeywords := []PublicationKeyword{}
	for rows.Next() {
		var publicationKeyword PublicationKeyword
		err := rows.Scan(
			&publicationKeyword.PublicationID,
			&publicationKeyword.KeywordID,
			&publicationKeyword.Title,
			&publicationKeyword.KeywordText,
		)
		if err != nil {
			sendError(w, http.StatusInternalServerError, err.Error())
			return
		}

		publicationKeywords = append(publicationKeywords, publicationKeyword)
	}

	if err := rows.Err(); err != nil {
		sendError(w, http.StatusInternalServerError, err.Error())
		return
	}

	sendJSON(w, http.StatusOK, publicationKeywords)
}

func createPublicationKeyword(w http.ResponseWriter, r *http.Request) {
	var publicationKeyword PublicationKeyword
	if err := decodeJSONBody(r, &publicationKeyword); err != nil {
		sendError(w, http.StatusBadRequest, "invalid JSON body")
		return
	}

	_, err := DB.Exec(`
		INSERT INTO PublicationKeyword (PublicationID, KeywordID)
		VALUES (?, ?)
	`, publicationKeyword.PublicationID, publicationKeyword.KeywordID)
	if err != nil {
		sendError(w, http.StatusInternalServerError, err.Error())
		return
	}

	sendJSON(w, http.StatusCreated, publicationKeyword)
}

func deletePublicationKeyword(w http.ResponseWriter, r *http.Request) {
	publicationID, keywordID, err := parseTwoIDsFromPath(r.URL.Path, "/api/publication-keywords/")
	if err != nil {
		sendError(w, http.StatusBadRequest, "invalid publication keyword ids")
		return
	}

	result, err := DB.Exec(`
		DELETE FROM PublicationKeyword
		WHERE PublicationID = ?
		  AND KeywordID = ?
	`, publicationID, keywordID)
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
		sendError(w, http.StatusNotFound, "publication keyword record not found")
		return
	}

	sendJSON(w, http.StatusOK, map[string]string{"message": "publication keyword record deleted successfully"})
}
