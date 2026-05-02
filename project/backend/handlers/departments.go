package handlers

import (
	"database/sql"
	"net/http"
)

func departmentsHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		getDepartments(w, r)
	case http.MethodPost:
		createDepartment(w, r)
	default:
		methodNotAllowed(w)
	}
}

func departmentByIDHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		getDepartmentByID(w, r)
	case http.MethodPut:
		updateDepartment(w, r)
	case http.MethodDelete:
		deleteDepartment(w, r)
	default:
		methodNotAllowed(w)
	}
}

func getDepartments(w http.ResponseWriter, r *http.Request) {
	rows, err := DB.Query(`
		SELECT DepartmentID, DepartmentName, College
		FROM Department
		ORDER BY DepartmentID
	`)
	if err != nil {
		sendError(w, http.StatusInternalServerError, err.Error())
		return
	}
	defer rows.Close()

	departments := []Department{}
	for rows.Next() {
		var department Department
		err := rows.Scan(&department.DepartmentID, &department.DepartmentName, &department.College)
		if err != nil {
			sendError(w, http.StatusInternalServerError, err.Error())
			return
		}

		departments = append(departments, department)
	}

	if err := rows.Err(); err != nil {
		sendError(w, http.StatusInternalServerError, err.Error())
		return
	}

	sendJSON(w, http.StatusOK, departments)
}

func getDepartmentByID(w http.ResponseWriter, r *http.Request) {
	id, err := parseIDFromPath(r.URL.Path, "/api/departments/")
	if err != nil {
		sendError(w, http.StatusBadRequest, "invalid department id")
		return
	}

	var department Department
	err = DB.QueryRow(`
		SELECT DepartmentID, DepartmentName, College
		FROM Department
		WHERE DepartmentID = ?
	`, id).Scan(&department.DepartmentID, &department.DepartmentName, &department.College)
	if err == sql.ErrNoRows {
		sendError(w, http.StatusNotFound, "department not found")
		return
	}
	if err != nil {
		sendError(w, http.StatusInternalServerError, err.Error())
		return
	}

	sendJSON(w, http.StatusOK, department)
}

func createDepartment(w http.ResponseWriter, r *http.Request) {
	var department Department
	if err := decodeJSONBody(r, &department); err != nil {
		sendError(w, http.StatusBadRequest, "invalid JSON body")
		return
	}

	result, err := DB.Exec(`
		INSERT INTO Department (DepartmentName, College)
		VALUES (?, ?)
	`, department.DepartmentName, department.College)
	if err != nil {
		sendError(w, http.StatusInternalServerError, err.Error())
		return
	}

	id, err := result.LastInsertId()
	if err != nil {
		sendError(w, http.StatusInternalServerError, err.Error())
		return
	}

	department.DepartmentID = int(id)
	sendJSON(w, http.StatusCreated, department)
}

func updateDepartment(w http.ResponseWriter, r *http.Request) {
	id, err := parseIDFromPath(r.URL.Path, "/api/departments/")
	if err != nil {
		sendError(w, http.StatusBadRequest, "invalid department id")
		return
	}

	var department Department
	if err := decodeJSONBody(r, &department); err != nil {
		sendError(w, http.StatusBadRequest, "invalid JSON body")
		return
	}

	result, err := DB.Exec(`
		UPDATE Department
		SET DepartmentName = ?, College = ?
		WHERE DepartmentID = ?
	`, department.DepartmentName, department.College, id)
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
		sendError(w, http.StatusNotFound, "department not found")
		return
	}

	department.DepartmentID = id
	sendJSON(w, http.StatusOK, department)
}

func deleteDepartment(w http.ResponseWriter, r *http.Request) {
	id, err := parseIDFromPath(r.URL.Path, "/api/departments/")
	if err != nil {
		sendError(w, http.StatusBadRequest, "invalid department id")
		return
	}

	result, err := DB.Exec(`
		DELETE FROM Department
		WHERE DepartmentID = ?
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
		sendError(w, http.StatusNotFound, "department not found")
		return
	}

	sendJSON(w, http.StatusOK, map[string]string{"message": "department deleted successfully"})
}
