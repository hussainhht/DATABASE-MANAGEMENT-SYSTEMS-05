package handlers

import "net/http"

func authorshipHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		getAuthorships(w, r)
	case http.MethodPost:
		createAuthorship(w, r)
	default:
		methodNotAllowed(w)
	}
}

func authorshipByIDsHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPut:
		updateAuthorship(w, r)
	case http.MethodDelete:
		deleteAuthorship(w, r)
	default:
		methodNotAllowed(w)
	}
}

func getAuthorships(w http.ResponseWriter, r *http.Request) {
	rows, err := DB.Query(`
		SELECT
			a.ResearcherID,
			a.PublicationID,
			a.AuthorOrder,
			r.FullName,
			p.Title
		FROM Authorship a, Researcher r, Publication p
		WHERE a.ResearcherID = r.ResearcherID
		  AND a.PublicationID = p.PublicationID
		ORDER BY a.PublicationID, a.AuthorOrder
	`)
	if err != nil {
		sendError(w, http.StatusInternalServerError, err.Error())
		return
	}
	defer rows.Close()

	authorships := []Authorship{}
	for rows.Next() {
		var authorship Authorship
		err := rows.Scan(
			&authorship.ResearcherID,
			&authorship.PublicationID,
			&authorship.AuthorOrder,
			&authorship.ResearcherName,
			&authorship.PublicationTitle,
		)
		if err != nil {
			sendError(w, http.StatusInternalServerError, err.Error())
			return
		}

		authorships = append(authorships, authorship)
	}

	if err := rows.Err(); err != nil {
		sendError(w, http.StatusInternalServerError, err.Error())
		return
	}

	sendJSON(w, http.StatusOK, authorships)
}

func createAuthorship(w http.ResponseWriter, r *http.Request) {
	var authorship Authorship
	if err := decodeJSONBody(r, &authorship); err != nil {
		sendError(w, http.StatusBadRequest, "invalid JSON body")
		return
	}

	_, err := DB.Exec(`
		INSERT INTO Authorship (ResearcherID, PublicationID, AuthorOrder)
		VALUES (?, ?, ?)
	`, authorship.ResearcherID, authorship.PublicationID, authorship.AuthorOrder)
	if err != nil {
		sendError(w, http.StatusInternalServerError, err.Error())
		return
	}

	sendJSON(w, http.StatusCreated, authorship)
}

func updateAuthorship(w http.ResponseWriter, r *http.Request) {
	researcherID, publicationID, err := parseTwoIDsFromPath(r.URL.Path, "/api/authorship/")
	if err != nil {
		sendError(w, http.StatusBadRequest, "invalid authorship ids")
		return
	}

	var authorship Authorship
	if err := decodeJSONBody(r, &authorship); err != nil {
		sendError(w, http.StatusBadRequest, "invalid JSON body")
		return
	}

	result, err := DB.Exec(`
		UPDATE Authorship
		SET AuthorOrder = ?
		WHERE ResearcherID = ?
		  AND PublicationID = ?
	`, authorship.AuthorOrder, researcherID, publicationID)
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
		sendError(w, http.StatusNotFound, "authorship record not found")
		return
	}

	authorship.ResearcherID = researcherID
	authorship.PublicationID = publicationID
	sendJSON(w, http.StatusOK, authorship)
}

func deleteAuthorship(w http.ResponseWriter, r *http.Request) {
	researcherID, publicationID, err := parseTwoIDsFromPath(r.URL.Path, "/api/authorship/")
	if err != nil {
		sendError(w, http.StatusBadRequest, "invalid authorship ids")
		return
	}

	result, err := DB.Exec(`
		DELETE FROM Authorship
		WHERE ResearcherID = ?
		  AND PublicationID = ?
	`, researcherID, publicationID)
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
		sendError(w, http.StatusNotFound, "authorship record not found")
		return
	}

	sendJSON(w, http.StatusOK, map[string]string{"message": "authorship record deleted successfully"})
}
