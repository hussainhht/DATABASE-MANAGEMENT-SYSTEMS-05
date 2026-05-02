package handlers

import "net/http"

func publicationsWithAuthorsHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		rows, err := DB.Query(`
			SELECT
				p.PublicationID,
				p.Title,
				p.PublicationYear,
				p.PublicationType,
				v.VenueName,
				r.FullName,
				a.AuthorOrder
			FROM Publication p, Venue v, Authorship a, Researcher r
			WHERE p.VenueID = v.VenueID
			  AND p.PublicationID = a.PublicationID
			  AND a.ResearcherID = r.ResearcherID
			ORDER BY p.PublicationID, a.AuthorOrder
		`)
		if err != nil {
			sendError(w, http.StatusInternalServerError, err.Error())
			return
		}
		defer rows.Close()

		reportRows := []PublicationWithAuthorReport{}
		for rows.Next() {
			var reportRow PublicationWithAuthorReport
			err := rows.Scan(
				&reportRow.PublicationID,
				&reportRow.Title,
				&reportRow.PublicationYear,
				&reportRow.PublicationType,
				&reportRow.VenueName,
				&reportRow.ResearcherName,
				&reportRow.AuthorOrder,
			)
			if err != nil {
				sendError(w, http.StatusInternalServerError, err.Error())
				return
			}

			reportRows = append(reportRows, reportRow)
		}

		if err := rows.Err(); err != nil {
			sendError(w, http.StatusInternalServerError, err.Error())
			return
		}

		sendJSON(w, http.StatusOK, reportRows)
	default:
		methodNotAllowed(w)
	}
}

func publicationsByDepartmentHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		rows, err := DB.Query(`
			SELECT
				d.DepartmentID,
				d.DepartmentName,
				COUNT(DISTINCT p.PublicationID) AS PublicationCount
			FROM Department d, Researcher r, Authorship a, Publication p
			WHERE d.DepartmentID = r.DepartmentID
			  AND r.ResearcherID = a.ResearcherID
			  AND a.PublicationID = p.PublicationID
			GROUP BY d.DepartmentID, d.DepartmentName
			ORDER BY PublicationCount DESC, d.DepartmentName
		`)
		if err != nil {
			sendError(w, http.StatusInternalServerError, err.Error())
			return
		}
		defer rows.Close()

		reportRows := []DepartmentPublicationReport{}
		for rows.Next() {
			var reportRow DepartmentPublicationReport
			err := rows.Scan(&reportRow.DepartmentID, &reportRow.DepartmentName, &reportRow.PublicationCount)
			if err != nil {
				sendError(w, http.StatusInternalServerError, err.Error())
				return
			}

			reportRows = append(reportRows, reportRow)
		}

		if err := rows.Err(); err != nil {
			sendError(w, http.StatusInternalServerError, err.Error())
			return
		}

		sendJSON(w, http.StatusOK, reportRows)
	default:
		methodNotAllowed(w)
	}
}

func publicationsByResearcherHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		rows, err := DB.Query(`
			SELECT
				r.ResearcherID,
				r.FullName,
				COUNT(a.PublicationID) AS PublicationCount
			FROM Researcher r, Authorship a
			WHERE r.ResearcherID = a.ResearcherID
			GROUP BY r.ResearcherID, r.FullName
			ORDER BY PublicationCount DESC, r.FullName
		`)
		if err != nil {
			sendError(w, http.StatusInternalServerError, err.Error())
			return
		}
		defer rows.Close()

		reportRows := []ResearcherPublicationReport{}
		for rows.Next() {
			var reportRow ResearcherPublicationReport
			err := rows.Scan(&reportRow.ResearcherID, &reportRow.FullName, &reportRow.PublicationCount)
			if err != nil {
				sendError(w, http.StatusInternalServerError, err.Error())
				return
			}

			reportRows = append(reportRows, reportRow)
		}

		if err := rows.Err(); err != nil {
			sendError(w, http.StatusInternalServerError, err.Error())
			return
		}

		sendJSON(w, http.StatusOK, reportRows)
	default:
		methodNotAllowed(w)
	}
}

func publicationsByYearHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		rows, err := DB.Query(`
			SELECT PublicationYear, COUNT(*) AS PublicationCount
			FROM Publication
			GROUP BY PublicationYear
			ORDER BY PublicationYear DESC
		`)
		if err != nil {
			sendError(w, http.StatusInternalServerError, err.Error())
			return
		}
		defer rows.Close()

		reportRows := []YearPublicationReport{}
		for rows.Next() {
			var reportRow YearPublicationReport
			err := rows.Scan(&reportRow.PublicationYear, &reportRow.PublicationCount)
			if err != nil {
				sendError(w, http.StatusInternalServerError, err.Error())
				return
			}

			reportRows = append(reportRows, reportRow)
		}

		if err := rows.Err(); err != nil {
			sendError(w, http.StatusInternalServerError, err.Error())
			return
		}

		sendJSON(w, http.StatusOK, reportRows)
	default:
		methodNotAllowed(w)
	}
}

func researchersMoreThanOnePublicationHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		rows, err := DB.Query(`
			SELECT
				r.ResearcherID,
				r.FullName,
				COUNT(a.PublicationID) AS PublicationCount
			FROM Researcher r, Authorship a
			WHERE r.ResearcherID = a.ResearcherID
			GROUP BY r.ResearcherID, r.FullName
			HAVING COUNT(a.PublicationID) > 1
			ORDER BY PublicationCount DESC, r.FullName
		`)
		if err != nil {
			sendError(w, http.StatusInternalServerError, err.Error())
			return
		}
		defer rows.Close()

		reportRows := []ResearcherPublicationReport{}
		for rows.Next() {
			var reportRow ResearcherPublicationReport
			err := rows.Scan(&reportRow.ResearcherID, &reportRow.FullName, &reportRow.PublicationCount)
			if err != nil {
				sendError(w, http.StatusInternalServerError, err.Error())
				return
			}

			reportRows = append(reportRows, reportRow)
		}

		if err := rows.Err(); err != nil {
			sendError(w, http.StatusInternalServerError, err.Error())
			return
		}

		sendJSON(w, http.StatusOK, reportRows)
	default:
		methodNotAllowed(w)
	}
}

func publicationsByKeywordHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		keyword := r.URL.Query().Get("keyword")
		if keyword == "" {
			sendError(w, http.StatusBadRequest, "keyword query parameter is required")
			return
		}

		rows, err := DB.Query(`
			SELECT DISTINCT
				p.PublicationID,
				p.Title,
				k.KeywordText
			FROM Publication p, PublicationKeyword pk, Keyword k
			WHERE p.PublicationID = pk.PublicationID
			  AND pk.KeywordID = k.KeywordID
			  AND k.KeywordText LIKE ?
			ORDER BY p.PublicationID
		`, "%"+keyword+"%")
		if err != nil {
			sendError(w, http.StatusInternalServerError, err.Error())
			return
		}
		defer rows.Close()

		reportRows := []PublicationKeywordReport{}
		for rows.Next() {
			var reportRow PublicationKeywordReport
			err := rows.Scan(&reportRow.PublicationID, &reportRow.Title, &reportRow.KeywordText)
			if err != nil {
				sendError(w, http.StatusInternalServerError, err.Error())
				return
			}

			reportRows = append(reportRows, reportRow)
		}

		if err := rows.Err(); err != nil {
			sendError(w, http.StatusInternalServerError, err.Error())
			return
		}

		sendJSON(w, http.StatusOK, reportRows)
	default:
		methodNotAllowed(w)
	}
}

func publicationsByTypeHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		publicationType := r.URL.Query().Get("type")
		if publicationType == "" {
			sendError(w, http.StatusBadRequest, "type query parameter is required")
			return
		}

		rows, err := DB.Query(`
			SELECT PublicationID, Title, PublicationYear, PublicationType
			FROM Publication
			WHERE PublicationType = ?
			ORDER BY PublicationYear DESC, Title
		`, publicationType)
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
	default:
		methodNotAllowed(w)
	}
}

func publicationsByVenueHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		rows, err := DB.Query(`
			SELECT
				p.PublicationID,
				p.Title,
				v.VenueName,
				v.VenueType,
				COALESCE(v.Publisher, '') AS Publisher
			FROM Publication p, Venue v
			WHERE p.VenueID = v.VenueID
			ORDER BY v.VenueName, p.Title
		`)
		if err != nil {
			sendError(w, http.StatusInternalServerError, err.Error())
			return
		}
		defer rows.Close()

		reportRows := []PublicationVenueReport{}
		for rows.Next() {
			var reportRow PublicationVenueReport
			err := rows.Scan(
				&reportRow.PublicationID,
				&reportRow.Title,
				&reportRow.VenueName,
				&reportRow.VenueType,
				&reportRow.Publisher,
			)
			if err != nil {
				sendError(w, http.StatusInternalServerError, err.Error())
				return
			}

			reportRows = append(reportRows, reportRow)
		}

		if err := rows.Err(); err != nil {
			sendError(w, http.StatusInternalServerError, err.Error())
			return
		}

		sendJSON(w, http.StatusOK, reportRows)
	default:
		methodNotAllowed(w)
	}
}

func latestPublicationsHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		rows, err := DB.Query(`
			SELECT PublicationID, Title, PublicationYear, PublicationType
			FROM Publication
			ORDER BY PublicationYear DESC, PublicationID DESC
			LIMIT 5
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
	default:
		methodNotAllowed(w)
	}
}

func searchPublicationsHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		title := r.URL.Query().Get("title")
		if title == "" {
			sendError(w, http.StatusBadRequest, "title query parameter is required")
			return
		}

		rows, err := DB.Query(`
			SELECT PublicationID, Title, PublicationYear, PublicationType
			FROM Publication
			WHERE Title LIKE ?
			ORDER BY PublicationYear DESC, Title
		`, "%"+title+"%")
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
	default:
		methodNotAllowed(w)
	}
}
