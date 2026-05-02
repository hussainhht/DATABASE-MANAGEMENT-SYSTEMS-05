package handlers

import (
	"database/sql"
	"net/http"
)

var DB *sql.DB

func RegisterRoutes(db *sql.DB) {
	DB = db

	http.HandleFunc("/api/departments", departmentsHandler)
	http.HandleFunc("/api/departments/", departmentByIDHandler)
	http.HandleFunc("/api/researchers", researchersHandler)
	http.HandleFunc("/api/researchers/", researcherByIDHandler)
	http.HandleFunc("/api/venues", venuesHandler)
	http.HandleFunc("/api/venues/", venueByIDHandler)
	http.HandleFunc("/api/publications", publicationsHandler)
	http.HandleFunc("/api/publications/", publicationByIDHandler)
	http.HandleFunc("/api/keywords", keywordsHandler)
	http.HandleFunc("/api/keywords/", keywordByIDHandler)
	http.HandleFunc("/api/authorship", authorshipHandler)
	http.HandleFunc("/api/authorship/", authorshipByIDsHandler)
	http.HandleFunc("/api/publication-keywords", publicationKeywordsHandler)
	http.HandleFunc("/api/publication-keywords/", publicationKeywordByIDsHandler)

	http.HandleFunc("/api/reports/publications-with-authors", publicationsWithAuthorsHandler)
	http.HandleFunc("/api/reports/publications-by-department", publicationsByDepartmentHandler)
	http.HandleFunc("/api/reports/publications-by-researcher", publicationsByResearcherHandler)
	http.HandleFunc("/api/reports/publications-by-year", publicationsByYearHandler)
	http.HandleFunc("/api/reports/researchers-more-than-one-publication", researchersMoreThanOnePublicationHandler)
	http.HandleFunc("/api/reports/publications-by-keyword", publicationsByKeywordHandler)
	http.HandleFunc("/api/reports/publications-by-type", publicationsByTypeHandler)
	http.HandleFunc("/api/reports/publications-by-venue", publicationsByVenueHandler)
	http.HandleFunc("/api/reports/latest-publications", latestPublicationsHandler)
	http.HandleFunc("/api/reports/search-publications", searchPublicationsHandler)
}
