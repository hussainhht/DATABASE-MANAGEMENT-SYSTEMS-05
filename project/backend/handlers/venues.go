package handlers

import (
	"database/sql"
	"net/http"
)

func venuesHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		getVenues(w, r)
	case http.MethodPost:
		createVenue(w, r)
	default:
		methodNotAllowed(w)
	}
}

func venueByIDHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		getVenueByID(w, r)
	case http.MethodPut:
		updateVenue(w, r)
	case http.MethodDelete:
		deleteVenue(w, r)
	default:
		methodNotAllowed(w)
	}
}

func getVenues(w http.ResponseWriter, r *http.Request) {
	rows, err := DB.Query(`
		SELECT VenueID, VenueName, VenueType, COALESCE(Publisher, '') AS Publisher
		FROM Venue
		ORDER BY VenueID
	`)
	if err != nil {
		sendError(w, http.StatusInternalServerError, err.Error())
		return
	}
	defer rows.Close()

	venues := []Venue{}
	for rows.Next() {
		var venue Venue
		err := rows.Scan(&venue.VenueID, &venue.VenueName, &venue.VenueType, &venue.Publisher)
		if err != nil {
			sendError(w, http.StatusInternalServerError, err.Error())
			return
		}

		venues = append(venues, venue)
	}

	if err := rows.Err(); err != nil {
		sendError(w, http.StatusInternalServerError, err.Error())
		return
	}

	sendJSON(w, http.StatusOK, venues)
}

func getVenueByID(w http.ResponseWriter, r *http.Request) {
	id, err := parseIDFromPath(r.URL.Path, "/api/venues/")
	if err != nil {
		sendError(w, http.StatusBadRequest, "invalid venue id")
		return
	}

	var venue Venue
	err = DB.QueryRow(`
		SELECT VenueID, VenueName, VenueType, COALESCE(Publisher, '') AS Publisher
		FROM Venue
		WHERE VenueID = ?
	`, id).Scan(&venue.VenueID, &venue.VenueName, &venue.VenueType, &venue.Publisher)
	if err == sql.ErrNoRows {
		sendError(w, http.StatusNotFound, "venue not found")
		return
	}
	if err != nil {
		sendError(w, http.StatusInternalServerError, err.Error())
		return
	}

	sendJSON(w, http.StatusOK, venue)
}

func createVenue(w http.ResponseWriter, r *http.Request) {
	var venue Venue
	if err := decodeJSONBody(r, &venue); err != nil {
		sendError(w, http.StatusBadRequest, "invalid JSON body")
		return
	}

	result, err := DB.Exec(`
		INSERT INTO Venue (VenueName, VenueType, Publisher)
		VALUES (?, ?, ?)
	`, venue.VenueName, venue.VenueType, nullIfEmpty(venue.Publisher))
	if err != nil {
		sendError(w, http.StatusInternalServerError, err.Error())
		return
	}

	id, err := result.LastInsertId()
	if err != nil {
		sendError(w, http.StatusInternalServerError, err.Error())
		return
	}

	venue.VenueID = int(id)
	sendJSON(w, http.StatusCreated, venue)
}

func updateVenue(w http.ResponseWriter, r *http.Request) {
	id, err := parseIDFromPath(r.URL.Path, "/api/venues/")
	if err != nil {
		sendError(w, http.StatusBadRequest, "invalid venue id")
		return
	}

	var venue Venue
	if err := decodeJSONBody(r, &venue); err != nil {
		sendError(w, http.StatusBadRequest, "invalid JSON body")
		return
	}

	result, err := DB.Exec(`
		UPDATE Venue
		SET VenueName = ?, VenueType = ?, Publisher = ?
		WHERE VenueID = ?
	`, venue.VenueName, venue.VenueType, nullIfEmpty(venue.Publisher), id)
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
		sendError(w, http.StatusNotFound, "venue not found")
		return
	}

	venue.VenueID = id
	sendJSON(w, http.StatusOK, venue)
}

func deleteVenue(w http.ResponseWriter, r *http.Request) {
	id, err := parseIDFromPath(r.URL.Path, "/api/venues/")
	if err != nil {
		sendError(w, http.StatusBadRequest, "invalid venue id")
		return
	}

	result, err := DB.Exec(`
		DELETE FROM Venue
		WHERE VenueID = ?
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
		sendError(w, http.StatusNotFound, "venue not found")
		return
	}

	sendJSON(w, http.StatusOK, map[string]string{"message": "venue deleted successfully"})
}
