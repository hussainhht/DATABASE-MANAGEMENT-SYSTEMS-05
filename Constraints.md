# Constraints

This section describes the main constraints used in the Research Publication Tracker database.

---

## 1. Primary Key Constraints

Primary keys are used to uniquely identify each record in a table.

- DepartmentID is the primary key in Department.
- ResearcherID is the primary key in Researcher.
- VenueID is the primary key in Venue.
- PublicationID is the primary key in Publication.
- KeywordID is the primary key in Keyword.
- The combination of ResearcherID and PublicationID is the primary key in Authorship.
- The combination of PublicationID and KeywordID is the primary key in PublicationKeyword.

---

## 2. Foreign Key Constraints

Foreign keys are used to connect related tables together.

- Researcher.DepartmentID references Department.DepartmentID.
- Publication.VenueID references Venue.VenueID.
- Authorship.ResearcherID references Researcher.ResearcherID.
- Authorship.PublicationID references Publication.PublicationID.
- PublicationKeyword.PublicationID references Publication.PublicationID.
- PublicationKeyword.KeywordID references Keyword.KeywordID.

---

## 3. Unique Constraints

Unique constraints prevent duplicate values in important fields.

- Researcher.Email must be unique.
- Publication.DOI must be unique if it is available.
- Keyword.KeywordText should be unique.

---

## 4. Not Null Constraints

Not null constraints make sure important fields cannot be empty.

- DepartmentID cannot be null.
- DepartmentName cannot be null.
- ResearcherID cannot be null.
- FullName cannot be null.
- Email cannot be null.
- PublicationID cannot be null.
- Title cannot be null.
- PublicationYear cannot be null.
- PublicationType cannot be null.
- VenueID cannot be null.
- KeywordID cannot be null.
- KeywordText cannot be null.
- AuthorOrder cannot be null.

---

## 5. Check Constraints

Check constraints limit the allowed values in some fields.

- PublicationYear must be a valid year.
- PublicationType must be one of the following values:
  - Journal
  - Conference
  - Book Chapter
  - Report

- VenueType must be one of the following values:
  - Journal
  - Conference
  - Book
  - Other

- AuthorOrder must be greater than 0.

---

## 6. Referential Integrity Constraints

Referential integrity ensures that foreign key values must refer to existing records in the referenced table.

Examples:

- A researcher cannot be assigned to a department that does not exist.
- A publication cannot be linked to a venue that does not exist.
- An authorship record cannot be created unless both the researcher and publication already exist.
- A publication keyword record cannot be created unless both the publication and keyword already exist.

---

## 7. Business Rule Constraints

The database should follow these business rules:

- Each researcher must belong to one department.
- Each department can have many researchers.
- Each publication must have at least one researcher as an author.
- Each publication can have many researchers as authors.
- Each researcher can author many publications.
- Each publication can be linked to one venue.
- Each venue can contain many publications.
- Each publication can have many keywords.
- Each keyword can be linked to many publications.
- The author order must be stored for each researcher in a publication.

