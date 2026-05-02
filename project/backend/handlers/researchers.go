package handlers

import (
	"database/sql"
	"net/http"
)

func researchersHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		getResearchers(w, r)
	case http.MethodPost:
		createResearcher(w, r)
	default:
		methodNotAllowed(w)
	}
}

func researcherByIDHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		getResearcherByID(w, r)
	case http.MethodPut:
		updateResearcher(w, r)
	case http.MethodDelete:
		deleteResearcher(w, r)
	default:
		methodNotAllowed(w)
	}
}

func getResearchers(w http.ResponseWriter, r *http.Request) {
	rows, err := DB.Query(`
		SELECT r.ResearcherID, r.FullName, r.Email, r.Role, r.DepartmentID, d.DepartmentName
		FROM Researcher r, Department d
		WHERE r.DepartmentID = d.DepartmentID
		ORDER BY r.ResearcherID
	`)
	if err != nil {
		sendError(w, http.StatusInternalServerError, err.Error())
		return
	}
	defer rows.Close()

	researchers := []Researcher{}
	for rows.Next() {
		var researcher Researcher
		err := rows.Scan(
			&researcher.ResearcherID,
			&researcher.FullName,
			&researcher.Email,
			&researcher.Role,
			&researcher.DepartmentID,
			&researcher.DepartmentName,
		)
		if err != nil {
			sendError(w, http.StatusInternalServerError, err.Error())
			return
		}

		researchers = append(researchers, researcher)
	}

	if err := rows.Err(); err != nil {
		sendError(w, http.StatusInternalServerError, err.Error())
		return
	}

	sendJSON(w, http.StatusOK, researchers)
}

func getResearcherByID(w http.ResponseWriter, r *http.Request) {
	id, err := parseIDFromPath(r.URL.Path, "/api/researchers/")
	if err != nil {
		sendError(w, http.StatusBadRequest, "invalid researcher id")
		return
	}

	var researcher Researcher
	err = DB.QueryRow(`
		SELECT r.ResearcherID, r.FullName, r.Email, r.Role, r.DepartmentID, d.DepartmentName
		FROM Researcher r, Department d
		WHERE r.DepartmentID = d.DepartmentID
		  AND r.ResearcherID = ?
	`, id).Scan(
		&researcher.ResearcherID,
		&researcher.FullName,
		&researcher.Email,
		&researcher.Role,
		&researcher.DepartmentID,
		&researcher.DepartmentName,
	)
	if err == sql.ErrNoRows {
		sendError(w, http.StatusNotFound, "researcher not found")
		return
	}
	if err != nil {
		sendError(w, http.StatusInternalServerError, err.Error())
		return
	}

	sendJSON(w, http.StatusOK, researcher)
}

func createResearcher(w http.ResponseWriter, r *http.Request) {
	var researcher Researcher
	if err := decodeJSONBody(r, &researcher); err != nil {
		sendError(w, http.StatusBadRequest, "invalid JSON body")
		return
	}

	result, err := DB.Exec(`
		INSERT INTO Researcher (FullName, Email, Role, DepartmentID)
		VALUES (?, ?, ?, ?)
	`, researcher.FullName, researcher.Email, researcher.Role, researcher.DepartmentID)
	if err != nil {
		sendError(w, http.StatusInternalServerError, err.Error())
		return
	}

	id, err := result.LastInsertId()
	if err != nil {
		sendError(w, http.StatusInternalServerError, err.Error())
		return
	}

	researcher.ResearcherID = int(id)
	sendJSON(w, http.StatusCreated, researcher)
}

func updateResearcher(w http.ResponseWriter, r *http.Request) {
	id, err := parseIDFromPath(r.URL.Path, "/api/researchers/")
	if err != nil {
		sendError(w, http.StatusBadRequest, "invalid researcher id")
		return
	}

	var researcher Researcher
	if err := decodeJSONBody(r, &researcher); err != nil {
		sendError(w, http.StatusBadRequest, "invalid JSON body")
		return
	}

	result, err := DB.Exec(`
		UPDATE Researcher
		SET FullName = ?, Email = ?, Role = ?, DepartmentID = ?
		WHERE ResearcherID = ?
	`, researcher.FullName, researcher.Email, researcher.Role, researcher.DepartmentID, id)
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
		sendError(w, http.StatusNotFound, "researcher not found")
		return
	}

	researcher.ResearcherID = id
	sendJSON(w, http.StatusOK, researcher)
}

func deleteResearcher(w http.ResponseWriter, r *http.Request) {
	id, err := parseIDFromPath(r.URL.Path, "/api/researchers/")
	if err != nil {
		sendError(w, http.StatusBadRequest, "invalid researcher id")
		return
	}

	result, err := DB.Exec(`
		DELETE FROM Researcher
		WHERE ResearcherID = ?
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
		sendError(w, http.StatusNotFound, "researcher not found")
		return
	}

	sendJSON(w, http.StatusOK, map[string]string{"message": "researcher deleted successfully"})
}
