# Research Publication Tracker (RPT)

**Course:** ITCS285 - Database Management Systems  
**Institution:** University of Bahrain, College of Information Technology  
**Section:** 5  
**Submission Date:** 17 May 2026

---

## Table of Contents

- [Project Overview](#project-overview)
- [Features](#features)
- [System Architecture](#system-architecture)
- [Technology Stack](#technology-stack)
- [Database Schema](#database-schema)
- [Project Structure](#project-structure)
- [Installation & Setup](#installation--setup)
- [Running the Application](#running-the-application)
- [API Endpoints](#api-endpoints)
- [Available Reports](#available-reports)
- [Lab Assignments](#lab-assignments)

---

## Project Overview

The **Research Publication Tracker** is a comprehensive database management system designed to organize, track, and manage academic research publications at the university level. The system provides a structured approach to storing information about researchers, departments, publications, publication venues, and keywords.

### Problem Statement

Universities produce numerous research outputs annually, including journal articles, conference papers, book chapters, and reports. Without a structured database system, tracking publication authorship, departmental contributions, and research topics becomes difficult and inefficient.

### Solution

This project implements a relational database system that:

- Centralizes research publication data
- Tracks researcher information and departmental affiliations
- Manages publication details and metadata
- Supports complex queries and reporting
- Provides a user-friendly web interface for data access

---

## Features

✅ **Researcher Management**

- Store researcher information (name, email, role, department)
- Track researcher affiliations and contact details

✅ **Department Tracking**

- Organize researchers by department and college
- Generate department-level publication statistics

✅ **Publication Management**

- Store comprehensive publication details (title, year, type, DOI, abstract)
- Link publications to research venues (journals, conferences)
- Support multiple publication types (journal articles, conference papers, etc.)

✅ **Authorship Tracking**

- Support multiple authors per publication
- Maintain author order/sequence
- Generate author-publication reports

✅ **Keyword Management**

- Tag publications with relevant keywords
- Support many-to-many relationships between publications and keywords
- Facilitate keyword-based search and filtering

✅ **Advanced Reporting**

- Publications by department
- Publications by researcher
- Publications by year
- Publications by venue
- Publications by keyword
- Author contribution analysis

✅ **Web Interface**

- Dashboard with key statistics
- Browsable data views (Departments, Researchers, Publications)
- Detailed report generation
- Full-text search capabilities

---

## System Architecture

### Architecture Overview

```
┌─────────────────────────────────────────┐
│         Web Frontend (HTML/CSS/JS)      │
│  - Dashboard                             │
│  - Data Management Views                 │
│  - Report Generation                     │
└────────────────────┬────────────────────┘
                     │ HTTP REST API
┌────────────────────▼────────────────────┐
│        Backend API Server (Go)          │
│  - Request Routing                       │
│  - Business Logic                        │
│  - Database Operations                   │
└────────────────────┬────────────────────┘
                     │ MySQL Driver
┌────────────────────▼────────────────────┐
│   MySQL Relational Database             │
│  - Tables & Relationships                │
│  - Constraints & Validation              │
│  - Data Persistence                      │
└─────────────────────────────────────────┘
```

### Design Patterns

- **REST API**: Standard HTTP endpoints for CRUD operations
- **MVC-like Structure**: Separation of concerns between handlers and business logic
- **Data Models**: Type-safe Go structs for all entities

---

## Technology Stack

| Component           | Technology                        | Version  |
| ------------------- | --------------------------------- | -------- |
| **Backend**         | Go                                | 1.24.2   |
| **Database**        | MySQL                             | 5.7+     |
| **Frontend**        | HTML5, CSS3, JavaScript (Vanilla) | -        |
| **HTTP Server**     | Go net/http                       | Built-in |
| **Database Driver** | github.com/go-sql-driver/mysql    | v1.10.0  |

---

## Database Schema

### Entity-Relationship Model

```
┌──────────────────────────────────────────────────────────┐
│                      Department                          │
│ PK: DepartmentID                                         │
│     DepartmentName                                       │
│     College                                              │
└──────────────────────────┬───────────────────────────────┘
                           │ 1:N
                           │
┌──────────────────────────▼───────────────────────────────┐
│                      Researcher                          │
│ PK: ResearcherID                                         │
│     FullName                                             │
│     Email (UNIQUE)                                       │
│     Role                                                 │
│ FK: DepartmentID                                         │
└─────────┬──────────────────────────┬────────────────────┘
          │                          │
          │ N:M (AuthorShip)         │ N:M (PublicationKeyword)
          │                          │
┌─────────▼─────────────┐  ┌────────▼──────────────────────┐
│    Publication        │  │         Keyword              │
│ PK: PublicationID     │  │ PK: KeywordID               │
│     Title             │  │     KeywordText (UNIQUE)    │
│     PublicationYear   │  └─────────────────────────────┘
│     PublicationType   │
│     DOI (UNIQUE)      │
│     Abstract          │
│ FK: VenueID           │
└─────────┬─────────────┘
          │ N:1
          │
┌─────────▼─────────────┐
│       Venue           │
│ PK: VenueID           │
│     VenueName         │
│     VenueType         │
│     Publisher         │
└───────────────────────┘
```

### Tables

#### Department

- **DepartmentID** (PK): Unique identifier
- **DepartmentName**: Name of the department
- **College**: College/Faculty name

#### Researcher

- **ResearcherID** (PK): Unique identifier
- **FullName**: Researcher's full name
- **Email** (UNIQUE): Email address
- **Role**: Position/role (e.g., Professor, Assistant Professor)
- **DepartmentID** (FK): Reference to Department

#### Publication

- **PublicationID** (PK): Unique identifier
- **Title**: Publication title
- **PublicationYear**: Year of publication
- **PublicationType**: Type (journal article, conference paper, etc.)
- **DOI** (UNIQUE): Digital Object Identifier
- **Abstract**: Publication abstract
- **VenueID** (FK): Reference to Venue

#### Venue

- **VenueID** (PK): Unique identifier
- **VenueName**: Venue name (journal/conference name)
- **VenueType**: Type of venue (journal, conference, etc.)
- **Publisher**: Publisher name

#### Keyword

- **KeywordID** (PK): Unique identifier
- **KeywordText** (UNIQUE): Keyword text

#### Authorship (Junction Table)

- **ResearcherID** (PK, FK): Reference to Researcher
- **PublicationID** (PK, FK): Reference to Publication
- **AuthorOrder**: Position of author in publication

#### PublicationKeyword (Junction Table)

- **PublicationID** (PK, FK): Reference to Publication
- **KeywordID** (PK, FK): Reference to Keyword

---

## Project Structure

```
DATABASE MANAGEMENT SYSTEMS-05/
│
├── lap/                          # Lab Assignments
│   ├── lap-assiment1/
│   │   ├── lab1-202405120.sql   # Lab 1 SQL Script
│   │   └── scripts.csv           # Lab 1 Scripts CSV
│   └── lap-assiment2/
│       └── lab2-202405120.sql   # Lab 2 SQL Script
│
├── project/                      # Main Project
│   ├── readme.md                 # Project documentation
│   │
│   ├── backend/                  # Go Backend API
│   │   ├── main.go              # Application entry point
│   │   ├── db.go                # Database connection setup
│   │   ├── models.go            # Data structure definitions
│   │   ├── go.mod               # Go module dependencies
│   │   └── handlers/            # HTTP request handlers
│   │       ├── routes.go        # Route registration
│   │       ├── departments.go   # Department endpoints
│   │       ├── researchers.go   # Researcher endpoints
│   │       ├── publications.go  # Publication endpoints
│   │       ├── venues.go        # Venue endpoints
│   │       ├── keywords.go      # Keyword endpoints
│   │       ├── authorship.go    # Authorship endpoints
│   │       ├── publication_keywords.go  # Keyword-publication junction
│   │       ├── reports.go       # Report generation
│   │       ├── helpers.go       # Utility functions
│   │       └── models.go        # Handler-specific types
│   │
│   ├── frontend/                 # Web Interface
│   │   ├── index.html           # Dashboard page
│   │   ├── departments.html     # Departments view
│   │   ├── researchers.html     # Researchers view
│   │   ├── publications.html    # Publications view
│   │   ├── reports.html         # Reports view
│   │   ├── css/
│   │   │   └── style.css        # Styling
│   │   └── js/
│   │       ├── api.js           # API client utilities
│   │       ├── departments.js   # Department page logic
│   │       ├── researchers.js   # Researcher page logic
│   │       ├── reports.js       # Report page logic
│   │
│   └── db/                       # Database Setup
│       ├── docs/                # Documentation
│       │   ├── README.md        # Project overview
│       │   ├── Relational-Schema.md     # Schema definition
│       │   ├── Constraints.md   # Constraint specifications
│       │   ├── SQL-Queries-and-Results.md  # Sample queries
│       │   └── System-Requirements.md      # Requirements doc
│       └── SQL/
│           ├── 01_create_database_and_tables.sql  # Table creation
│           ├── 02_insert_sample_data.sql          # Sample data
│           └── 03_project_queries.sql             # SQL queries
│
└── README.md                     # This file
```

---

## Installation & Setup

### Prerequisites

- **MySQL Server** (5.7 or higher)
- **Go** (1.24.2 or higher)
- **Web Browser** (for frontend)

### Step 1: Database Setup

1. Open MySQL command line or MySQL Workbench
2. Run the database setup script:

```sql
-- Create database and tables
SOURCE project/db/SQL/01_create_database_and_tables.sql;

-- Insert sample data
SOURCE project/db/SQL/02_insert_sample_data.sql;
```

Or execute the scripts using:

```bash
mysql -u root < project/db/SQL/01_create_database_and_tables.sql
mysql -u root < project/db/SQL/02_insert_sample_data.sql
```

### Step 2: Backend Setup

Navigate to the backend directory:

```bash
cd project/backend
```

Download dependencies:

```bash
go get github.com/go-sql-driver/mysql
```

### Step 3: Configuration

Update database credentials in `project/backend/db.go` if needed:

```go
username := "root"      // Your MySQL username
password := ""          // Your MySQL password
host := "127.0.0.1"     // Database host
port := "3306"          // MySQL port
dbName := "ResearchPublicationTracker"  // Database name
```

---

## Running the Application

### Starting the Backend Server

From the `project/backend` directory:

```bash
go run .
```

The server will start on `http://localhost:8080`

Output:

```
Connected to MySQL successfully
Server running on http://localhost:8080
```

### Accessing the Application

Open your web browser and navigate to:

```
http://localhost:8080
```

### Available Pages

- **Dashboard** (`/`) - Overview with key statistics
- **Departments** (`/departments.html`) - Department listing and management
- **Researchers** (`/researchers.html`) - Researcher information
- **Publications** (`/publications.html`) - Publication browsing
- **Reports** (`/reports.html`) - Advanced analytics and reporting

---

## API Endpoints

### Core Entity Endpoints

#### Departments

- `GET /api/departments` - List all departments
- `GET /api/departments/{id}` - Get specific department

#### Researchers

- `GET /api/researchers` - List all researchers
- `GET /api/researchers/{id}` - Get specific researcher

#### Publications

- `GET /api/publications` - List all publications
- `GET /api/publications/{id}` - Get specific publication

#### Venues

- `GET /api/venues` - List all venues
- `GET /api/venues/{id}` - Get specific venue

#### Keywords

- `GET /api/keywords` - List all keywords
- `GET /api/keywords/{id}` - Get specific keyword

#### Authorship

- `GET /api/authorship` - List all authorships
- `GET /api/authorship/{researcherId},{publicationId}` - Get specific authorship

#### Publication Keywords

- `GET /api/publication-keywords` - List all publication-keyword relationships
- `GET /api/publication-keywords/{publicationId},{keywordId}` - Get specific relationship

### Report Endpoints

- `GET /api/reports/publications-with-authors` - Publications with author details
- `GET /api/reports/publications-by-department` - Count by department
- `GET /api/reports/publications-by-researcher` - Count by researcher
- `GET /api/reports/publications-by-year` - Count by publication year
- `GET /api/reports/researchers-more-than-one-publication` - Prolific researchers
- `GET /api/reports/publications-by-keyword` - Count by keyword
- `GET /api/reports/publications-by-type` - Count by publication type
- `GET /api/reports/publications-by-venue` - Publications organized by venue
- `GET /api/reports/latest-publications` - Recent publications
- `GET /api/reports/search-publications?query={searchTerm}` - Full-text search

---

## Available Reports

The system provides comprehensive reporting capabilities:

### Statistical Reports

1. **Publications by Department** - Total publication count per department
2. **Publications by Researcher** - Publication count for each researcher
3. **Publications by Year** - Timeline of publications
4. **Publications by Type** - Distribution of publication types
5. **Publications by Venue** - Articles in different journals/conferences

### Analytical Reports

1. **Publications with Authors** - Detailed publication list with author information
2. **Prolific Researchers** - Researchers with multiple publications
3. **Latest Publications** - Most recent publications
4. **Publications by Keyword** - Articles tagged with specific keywords

### Search Capabilities

- Full-text publication search by title and abstract
- Filter by author, department, year, venue, or keyword

---

## Lab Assignments

### Lab 1 - Database Design & SQL Fundamentals

**File:** `lab/lap-assiment1/lab1-202405120.sql`

Topics covered:

- Entity identification and relationship modeling
- Primary and foreign key design
- Constraint specification
- Basic SQL queries

**Output:** `lab/lap-assiment1/scripts.csv`

### Lab 2 - Advanced SQL & Reporting

**File:** `lab/lap-assiment2/lab2-202405120.sql`

Topics covered:

- Complex SQL queries
- Aggregate functions
- JOIN operations
- Report generation
- Data analysis

---

## Data Model Highlights

### Key Design Decisions

1. **Surrogate Keys**: All tables use auto-incrementing integer primary keys for efficiency
2. **Unique Constraints**: Email and DOI fields are unique to prevent duplicates
3. **Referential Integrity**: Foreign keys enforce relationships between entities
4. **Many-to-Many Relationships**: Junction tables (Authorship, PublicationKeyword) handle complex relationships
5. **Author Order Tracking**: AuthorOrder field in Authorship table maintains publication author sequence

### Normalization

The database design adheres to **Third Normal Form (3NF)**:

- No transitive dependencies
- All non-key attributes are fully dependent on the primary key
- All determinant attributes are candidate keys

---

## Key Features

### Researcher Management

- Store researcher profile including name, email, and role
- Track departmental affiliations
- Support multiple departments (many-to-many through appropriate design)

### Publication Tracking

- Comprehensive publication metadata (title, year, type, DOI, abstract)
- Link publications to specific venues (journals, conferences)
- Categorize by publication type

### Author Attribution

- Track multiple authors per publication
- Maintain author sequence/order
- Generate author contribution statistics

### Keyword Management

- Tag publications with relevant keywords
- Support many-to-many publication-keyword relationships
- Facilitate research topic organization

### Reporting & Analytics

- Publication statistics by various dimensions
- Author contribution analysis
- Departmental publication metrics
- Temporal publication trends

---

## System Constraints

### Integrity Constraints

**Primary Key Constraints:**

- Each table has a unique identifier
- All primary keys are non-null and auto-incrementing

**Foreign Key Constraints:**

- Researcher.DepartmentID → Department.DepartmentID
- Publication.VenueID → Venue.VenueID
- Authorship.ResearcherID → Researcher.ResearcherID
- Authorship.PublicationID → Publication.PublicationID
- PublicationKeyword.PublicationID → Publication.PublicationID
- PublicationKeyword.KeywordID → Keyword.KeywordID

**Unique Constraints:**

- Researcher.Email (only one researcher per email)
- Publication.DOI (each DOI is unique)
- Keyword.KeywordText (no duplicate keywords)

**Domain Constraints:**

- PublicationYear: Positive integer
- AuthorOrder: Positive integer (sequence number)
- All text fields: Non-null and reasonable length limits

---

## Database Queries

The system supports the following important queries:

```sql
-- Publications with author details
SELECT p.PublicationID, p.Title, r.FullName, a.AuthorOrder
FROM Publication p
JOIN Authorship a ON p.PublicationID = a.PublicationID
JOIN Researcher r ON a.ResearcherID = r.ResearcherID
ORDER BY p.PublicationID, a.AuthorOrder;

-- Publications per department
SELECT d.DepartmentName, COUNT(p.PublicationID) as PublicationCount
FROM Department d
LEFT JOIN Researcher r ON d.DepartmentID = r.DepartmentID
LEFT JOIN Authorship a ON r.ResearcherID = a.ResearcherID
LEFT JOIN Publication p ON a.PublicationID = p.PublicationID
GROUP BY d.DepartmentID, d.DepartmentName;

-- Researchers with more than one publication
SELECT r.ResearcherID, r.FullName, COUNT(a.PublicationID) as PublicationCount
FROM Researcher r
LEFT JOIN Authorship a ON r.ResearcherID = a.ResearcherID
GROUP BY r.ResearcherID, r.FullName
HAVING COUNT(a.PublicationID) > 1;

-- Publications by keyword
SELECT k.KeywordText, COUNT(p.PublicationID) as PublicationCount
FROM Keyword k
LEFT JOIN PublicationKeyword pk ON k.KeywordID = pk.KeywordID
LEFT JOIN Publication p ON pk.PublicationID = p.PublicationID
GROUP BY k.KeywordID, k.KeywordText
ORDER BY PublicationCount DESC;
```

---

## Troubleshooting

### Database Connection Error

**Error:** `Database connection error: unknown driver "mysql"`

**Solution:** Ensure MySQL driver is installed:

```bash
go get github.com/go-sql-driver/mysql
```

### Port Already in Use

**Error:** `address already in use`

**Solution:** Change the port in `main.go` or set the PORT environment variable:

```bash
set PORT=8081  # Windows
go run .
```

### Database Not Found

**Error:** `Error 1049: Unknown database 'ResearchPublicationTracker'`

**Solution:** Ensure database setup scripts have been executed:

```bash
mysql -u root < project/db/SQL/01_create_database_and_tables.sql
```

---

## Contributing Guidelines

When modifying the system:

1. **Database Changes**: Update SQL files in `project/db/SQL/`
2. **API Changes**: Modify handlers in `project/backend/handlers/`
3. **Frontend Changes**: Update HTML/CSS/JS in `project/frontend/`
4. **Documentation**: Update relevant markdown files in `project/db/docs/`

---

## Course Information

**Course:** ITCS285 - Database Management Systems  
**Institution:** University of Bahrain  
**College:** College of Information Technology  
**Section:** 05  
**DBMS:** MySQL

### Learning Outcomes

Students will understand and apply:

- Relational database design principles
- Entity-relationship modeling
- Normalization techniques
- SQL programming
- Query optimization
- Database administration
- Application integration with databases

---

## References

- [MySQL Documentation](https://dev.mysql.com/doc/)
- [Go Database Package](https://golang.org/pkg/database/sql/)
- [SQL Standard Reference](https://en.wikipedia.org/wiki/SQL)
- University of Bahrain ITCS285 Course Materials

---

## License

This is an educational project for ITCS285 - Database Management Systems course at the University of Bahrain.

---

**Last Updated:** May 2, 2026  
**Version:** 1.0  
**Status:** Complete
