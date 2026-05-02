# Research Publication Tracker

## University of Bahrain

**College:** College of Information Technology  
**Course:** ITCS285 - Database Management Systems  
**Section:** 5  
**Instructor:** Dr. [Instructor Name]  
**Submission Date:** 17 May 2026

---

## DBMS Used: MySQL

## Prepared By

| Student | Name | ID |
| --- | --- | --- |
| Student 1 | Hussain Ali H. Ali | 202405120 |
| Student 2 | [Student Name] | [Student ID] |
| Student 3 | [Student Name] | [Student ID] |
| Student 4 | [Student Name] | [Student ID] |

---

## Project Overview

The **Research Publication Tracker** is a database management system project designed to organize and manage university research publications. The system focuses on storing information about researchers, departments, publications, publication venues, keywords, and authorship details.

Universities produce many research outputs every year, including journal articles, conference papers, book chapters, and reports. Without a structured database, it can become difficult to track who wrote each publication, which department produced it, where it was published, and what topics it covers. This project solves that problem by designing a relational database that stores publication data in a clear, organized, and searchable way.

The main goal of this project is to apply database design concepts learned in ITCS285, including entity identification, relationships, constraints, normalization, relational schema design, SQL table creation, sample data insertion, and useful SQL query writing.

---

## Project Objectives

This project aims to:

- Design a database system for tracking academic research publications.
- Store researcher information such as name, email, role, and department.
- Store department information such as department name and college.
- Store publication information such as title, year, type, DOI, abstract, and venue.
- Track publication venues such as journals, conferences, and publishers.
- Support many-to-many relationships between researchers and publications.
- Store the order of authors for each publication.
- Support many-to-many relationships between publications and keywords.
- Apply primary keys, foreign keys, unique constraints, and domain constraints.
- Write SQL statements to create tables and insert sample data.
- Write SQL queries that produce useful reports about publications.

---

## System Scope

The Research Publication Tracker will manage the following main data areas:

| Area | Description |
| --- | --- |
| Researchers | Stores information about faculty members, students, and external researchers. |
| Departments | Stores the academic department and college connected to each researcher. |
| Publications | Stores research output details such as title, year, type, DOI, and abstract. |
| Venues | Stores information about where publications are published, such as journals or conferences. |
| Keywords | Stores topic keywords that help classify and search publications. |
| Authorship | Connects researchers to publications and records author order. |
| Publication Keywords | Connects publications to their related keywords. |

---

## Expected Database Features

The database should support the following operations:

1. Add, store, and manage researcher records.
2. Add, store, and manage department records.
3. Add, store, and manage publication records.
4. Link each researcher to a department.
5. Link each publication to one publication venue.
6. Link multiple researchers to the same publication.
7. Link multiple publications to the same researcher.
8. Store the order of authors in each publication.
9. Link multiple keywords to the same publication.
10. Search and filter publications by researcher, department, year, type, venue, or keyword.
11. Generate reports such as publications per researcher, publications per department, and publications per year.

---

## Main Entities

The proposed database includes the following entities:

### Researcher

Stores information about researchers who contribute to publications.

Example attributes:

- ResearcherID
- FullName
- Email
- Role
- DepartmentID

### Department

Stores information about university departments.

Example attributes:

- DepartmentID
- DepartmentName
- College

### Publication

Stores information about research publications.

Example attributes:

- PublicationID
- Title
- PublicationYear
- PublicationType
- DOI
- Abstract
- VenueID

### Venue

Stores information about the publication source.

Example attributes:

- VenueID
- VenueName
- VenueType
- Publisher

### Keyword

Stores keywords used to describe and search publications.

Example attributes:

- KeywordID
- KeywordText

### Authorship

Represents the many-to-many relationship between researchers and publications.

Example attributes:

- ResearcherID
- PublicationID
- AuthorOrder

### PublicationKeyword

Represents the many-to-many relationship between publications and keywords.

Example attributes:

- PublicationID
- KeywordID

---

## Project Deliverables

The final project report should include the following sections:

1. Cover Page
2. Introduction
3. System Requirements
4. ER/EER Diagram
5. Relational Schema
6. Constraints
7. DBMS Used
8. SQL Create Table Statements
9. Inserted Data Samples
10. SQL Queries and Results
11. Contribution Sheet
12. Conclusion

---

## Repository Files

| File | Purpose |
| --- | --- |
| `README.md` | Main project overview and report outline. |
| `System-Requirements.md` | Detailed system requirements, business rules, assumptions, and expected queries. |
| `ITCS285-Project-S2-2025-2026.pdf` | Project instruction document for the course. |

---

## DBMS Used

The DBMS will be selected based on the project requirements and course instructions. Possible options include:

- MySQL
- Oracle Database
- Microsoft SQL Server
- PostgreSQL

The final selected DBMS should be clearly mentioned in the project report.

---

## Conclusion

The Research Publication Tracker is a practical database project that demonstrates how database systems can be used to organize academic research information. By designing a structured relational database, the system will make it easier to store, search, and report on university research publications.

This project also provides practice in important database management concepts, including entity relationship modeling, relational schema design, constraints, SQL table creation, data insertion, and query writing.




## steps
- Cover Page
- Introduction
- System Requirements
- ER/EER Diagram
- Relational Schema
- Constraints
- DBMS Used: MySQL
- SQL Create Tables
- Inserted Data Samples
- SQL Queries + Results
- Contribution Sheet
- Conclusion