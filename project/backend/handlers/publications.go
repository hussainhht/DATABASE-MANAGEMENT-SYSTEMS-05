package handlers

import (
	"database/sql"
	"net/http"
)

func publicationsHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		getPublications(w, r)
	case http.MethodPost:
		createPublication(w, r)
	default:
		methodNotAllowed(w)
	}
}

func publicationByIDHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		getPublicationByID(w, r)
	case http.MethodPut:
		updatePublication(w, r)
	case http.MethodDelete:
		deletePublication(w, r)
	default:
		methodNotAllowed(w)
	}
}

func getPublications(w http.ResponseWriter, r *http.Request) {
	rows, err := DB.Query(`
		SELECT
			p.PublicationID,
			p.Title,
			p.PublicationYear,
			p.PublicationType,
			COALESCE(p.DOI, '') AS DOI,
			COALESCE(p.Abstract, '') AS Abstract,
			p.VenueID,
			v.VenueName
		FROM Publication p, Venue v
		WHERE p.VenueID = v.VenueID
		ORDER BY p.PublicationID
	`)
	if err != nil {
		sendError(w, http.StatusInternalServerError, err.Error())
		return
	}
	defer rows.Close()

	publications := []Publication{}
	for rows.Next() {
		var publication Publication
		err := rows.Scan(
			&publication.PublicationID,
			&publication.Title,
			&publication.PublicationYear,
			&publication.PublicationType,
			&publication.DOI,
			&publication.Abstract,
			&publication.VenueID,
			&publication.VenueName,
		)
		if err != nil {
			sendError(w, http.StatusInternalServerError, err.Error())
			return
		}

		publications = append(publications, publication)
	}

	if err := rows.Err(); err != nil {
		sendError(w, http.StatusInternalServerError, err.Error())
		return
	}

	sendJSON(w, http.StatusOK, publications)
}

func getPublicationByID(w http.ResponseWriter, r *http.Request) {
	id, err := parseIDFromPath(r.URL.Path, "/api/publications/")
	if err != nil {
		sendError(w, http.StatusBadRequest, "invalid publication id")
		return
	}

	var publication Publication
	err = DB.QueryRow(`
		SELECT
			p.PublicationID,
			p.Title,
			p.PublicationYear,
			p.PublicationType,
			COALESCE(p.DOI, '') AS DOI,
			COALESCE(p.Abstract, '') AS Abstract,
			p.VenueID,
			v.VenueName
		FROM Publication p, Venue v
		WHERE p.VenueID = v.VenueID
		  AND p.PublicationID = ?
	`, id).Scan(
		&publication.PublicationID,
		&publication.Title,
		&publication.PublicationYear,
		&publication.PublicationType,
		&publication.DOI,
		&publication.Abstract,
		&publication.VenueID,
		&publication.VenueName,
	)
	if err == sql.ErrNoRows {
		sendError(w, http.StatusNotFound, "publication not found")
		return
	}
	if err != nil {
		sendError(w, http.StatusInternalServerError, err.Error())
		return
	}

	sendJSON(w, http.StatusOK, publication)
}

func createPublication(w http.ResponseWriter, r *http.Request) {
	var publication Publication
	if err := decodeJSONBody(r, &publication); err != nil {
		sendError(w, http.StatusBadRequest, "invalid JSON body")
		return
	}

	result, err := DB.Exec(`
		INSERT INTO Publication (Title, PublicationYear, PublicationType, DOI, Abstract, VenueID)
		VALUES (?, ?, ?, ?, ?, ?)
	`, publication.Title, publication.PublicationYear, publication.PublicationType, nullIfEmpty(publication.DOI), publication.Abstract, publication.VenueID)
	if err != nil {
		sendError(w, http.StatusInternalServerError, err.Error())
		return
	}

	id, err := result.LastInsertId()
	if err != nil {
		sendError(w, http.StatusInternalServerError, err.Error())
		return
	}

	publication.PublicationID = int(id)
	sendJSON(w, http.StatusCreated, publication)
}

func updatePublication(w http.ResponseWriter, r *http.Request) {
	id, err := parseIDFromPath(r.URL.Path, "/api/publications/")
	if err != nil {
		sendError(w, http.StatusBadRequest, "invalid publication id")
		return
	}

	var publication Publication
	if err := decodeJSONBody(r, &publication); err != nil {
		sendError(w, http.StatusBadRequest, "invalid JSON body")
		return
	}

	result, err := DB.Exec(`
		UPDATE Publication
		SET Title = ?, PublicationYear = ?, PublicationType = ?, DOI = ?, Abstract = ?, VenueID = ?
		WHERE PublicationID = ?
	`, publication.Title, publication.PublicationYear, publication.PublicationType, nullIfEmpty(publication.DOI), publication.Abstract, publication.VenueID, id)
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
		sendError(w, http.StatusNotFound, "publication not found")
		return
	}

	publication.PublicationID = id
	sendJSON(w, http.StatusOK, publication)
}

func deletePublication(w http.ResponseWriter, r *http.Request) {
	id, err := parseIDFromPath(r.URL.Path, "/api/publications/")
	if err != nil {
		sendError(w, http.StatusBadRequest, "invalid publication id")
		return
	}

	result, err := DB.Exec(`
		DELETE FROM Publication
		WHERE PublicationID = ?
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
		sendError(w, http.StatusNotFound, "publication not found")
		return
	}

	sendJSON(w, http.StatusOK, map[string]string{"message": "publication deleted successfully"})
}
