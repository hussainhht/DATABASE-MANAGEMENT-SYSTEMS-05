# Relational Schema

## Department
Department(DepartmentID, DepartmentName, College)

Primary Key:
- DepartmentID

---

## Researcher
Researcher(ResearcherID, FullName, Email, Role, DepartmentID)

Primary Key:
- ResearcherID

Foreign Key:
- DepartmentID references Department(DepartmentID)

Unique:
- Email

---

## Venue
Venue(VenueID, VenueName, VenueType, Publisher)

Primary Key:
- VenueID

---

## Publication
Publication(PublicationID, Title, PublicationYear, PublicationType, DOI, Abstract, VenueID)

Primary Key:
- PublicationID

Foreign Key:
- VenueID references Venue(VenueID)

Unique:
- DOI

---

## Keyword
Keyword(KeywordID, KeywordText)

Primary Key:
- KeywordID

Unique:
- KeywordText

---

## Authorship
Authorship(ResearcherID, PublicationID, AuthorOrder)

Primary Key:
- ResearcherID, PublicationID

Foreign Keys:
- ResearcherID references Researcher(ResearcherID)
- PublicationID references Publication(PublicationID)

---

## PublicationKeyword
PublicationKeyword(PublicationID, KeywordID)

Primary Key:
- PublicationID, KeywordID

Foreign Keys:
- PublicationID references Publication(PublicationID)
- KeywordID references Keyword(KeywordID)