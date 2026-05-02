package main

type Department struct {
	DepartmentID   int    `json:"department_id"`
	DepartmentName string `json:"department_name"`
	College        string `json:"college"`
}

type Researcher struct {
	ResearcherID   int    `json:"researcher_id"`
	FullName       string `json:"full_name"`
	Email          string `json:"email"`
	Role           string `json:"role"`
	DepartmentID   int    `json:"department_id"`
	DepartmentName string `json:"department_name,omitempty"`
}

type Venue struct {
	VenueID   int    `json:"venue_id"`
	VenueName string `json:"venue_name"`
	VenueType string `json:"venue_type"`
	Publisher string `json:"publisher"`
}

type Publication struct {
	PublicationID   int    `json:"publication_id"`
	Title           string `json:"title"`
	PublicationYear int    `json:"publication_year"`
	PublicationType string `json:"publication_type"`
	DOI             string `json:"doi"`
	Abstract        string `json:"abstract"`
	VenueID         int    `json:"venue_id"`
	VenueName       string `json:"venue_name,omitempty"`
}

type Keyword struct {
	KeywordID   int    `json:"keyword_id"`
	KeywordText string `json:"keyword_text"`
}

type Authorship struct {
	ResearcherID     int    `json:"researcher_id"`
	PublicationID    int    `json:"publication_id"`
	AuthorOrder      int    `json:"author_order"`
	ResearcherName   string `json:"researcher_name,omitempty"`
	PublicationTitle string `json:"publication_title,omitempty"`
}

type PublicationKeyword struct {
	PublicationID int    `json:"publication_id"`
	KeywordID     int    `json:"keyword_id"`
	Title         string `json:"title,omitempty"`
	KeywordText   string `json:"keyword_text,omitempty"`
}

type PublicationWithAuthorReport struct {
	PublicationID   int    `json:"publication_id"`
	Title           string `json:"title"`
	PublicationYear int    `json:"publication_year"`
	PublicationType string `json:"publication_type"`
	VenueName       string `json:"venue_name"`
	ResearcherName  string `json:"researcher_name"`
	AuthorOrder     int    `json:"author_order"`
}

type DepartmentPublicationReport struct {
	DepartmentID     int    `json:"department_id"`
	DepartmentName   string `json:"department_name"`
	PublicationCount int    `json:"publication_count"`
}

type ResearcherPublicationReport struct {
	ResearcherID     int    `json:"researcher_id"`
	FullName         string `json:"full_name"`
	PublicationCount int    `json:"publication_count"`
}

type YearPublicationReport struct {
	PublicationYear  int `json:"publication_year"`
	PublicationCount int `json:"publication_count"`
}

type PublicationKeywordReport struct {
	PublicationID int    `json:"publication_id"`
	Title         string `json:"title"`
	KeywordText   string `json:"keyword_text"`
}

type PublicationVenueReport struct {
	PublicationID int    `json:"publication_id"`
	Title         string `json:"title"`
	VenueName     string `json:"venue_name"`
	VenueType     string `json:"venue_type"`
	Publisher     string `json:"publisher"`
}
