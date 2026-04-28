# System Requirements

## Project Title
Research Publication Tracker

## System Overview
The Research Publication Tracker is a database system designed to store and manage information about research publications produced by university researchers. The system helps track researchers, departments, publications, publication venues, keywords, and authorship details.

The main purpose of the system is to organize research output and make it easy to search, retrieve, and generate reports about publications by researcher, department, year, type, venue, or keyword.

## System Objectives
- Store complete information about research publications.
- Track researchers and their departments.
- Track the authors of each publication.
- Store the order of authors in each publication.
- Store publication venues such as journals and conferences.
- Store keywords related to publications.
- Support useful SQL queries and reports.

## Functional Requirements

1. The system shall store information about each researcher.
2. The system shall store researcher ID, full name, email, role, and department.
3. The system shall store information about each department.
4. The system shall store department ID, department name, and college.
5. The system shall store information about each publication.
6. The system shall store publication ID, title, publication year, publication type, DOI, and abstract.
7. The system shall allow one researcher to author many publications.
8. The system shall allow one publication to have many researchers as authors.
9. The system shall store the author order for each publication.
10. The system shall store information about publication venues.
11. The system shall store venue ID, venue name, venue type, and publisher.
12. The system shall allow each publication to be linked to one venue.
13. The system shall store keywords related to publications.
14. The system shall allow one publication to have many keywords.
15. The system shall allow one keyword to be used by many publications.
16. The system shall support searching publications by title, researcher, department, year, type, venue, or keyword.
17. The system shall support generating reports such as publications per department, publications per researcher, and publications per year.

## Data Requirements

The system will store data about the following main entities:

### Researcher
Stores information about researchers who write publications.

Attributes:
- ResearcherID
- FullName
- Email
- Role
- DepartmentID

### Department
Stores information about university departments.

Attributes:
- DepartmentID
- DepartmentName
- College

### Publication
Stores information about research publications.

Attributes:
- PublicationID
- Title
- PublicationYear
- PublicationType
- DOI
- Abstract
- VenueID

### Venue
Stores information about where the publication was published.

Attributes:
- VenueID
- VenueName
- VenueType
- Publisher

### Keyword
Stores keywords used to describe publications.

Attributes:
- KeywordID
- KeywordText

### Authorship
Represents the many-to-many relationship between researchers and publications.

Attributes:
- ResearcherID
- PublicationID
- AuthorOrder

### PublicationKeyword
Represents the many-to-many relationship between publications and keywords.

Attributes:
- PublicationID
- KeywordID

## Business Rules

1. Each researcher must belong to one department.
2. Each department can have many researchers.
3. Each publication must have at least one researcher as an author.
4. Each researcher can author many publications.
5. Each publication can have many authors.
6. Each publication can be linked to one venue.
7. Each venue can contain many publications.
8. Each publication can have many keywords.
9. Each keyword can be linked to many publications.
10. The author order must be stored for each researcher in a publication.
11. Researcher email should be unique.
12. DOI should be unique if it is available.
13. Publication year should be a valid year.
14. Publication type should be limited to values such as Journal, Conference, Book Chapter, or Report.

## Expected Queries

The database should support queries such as:

1. List all publications written by a specific researcher.
2. List all publications published in a specific year.
3. List all publications from a specific department.
4. List all publications published in a specific venue.
5. List all publications with a specific keyword.
6. Count the number of publications for each researcher.
7. Count the number of publications for each department.
8. Display researchers who have more than one publication.
9. Display publications with their authors in author order.
10. Display publications by type, such as Journal or Conference.

## Assumptions

- A researcher can be a faculty member, student, or external researcher.
- A publication can have more than one author.
- A researcher may participate in many publications.
- A venue can be a journal, conference, or other publication source.
- Some publications may not have a DOI, but if a DOI exists, it must be unique.
- Keywords are used to make searching easier.