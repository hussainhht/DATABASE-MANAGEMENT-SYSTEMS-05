# Research Publication Tracker - Project Implementation

## System Architecture

```
┌──────────────────────────────┐
│   Browser Pages              │
│   (HTML / CSS / JS)          │
└────────────────┬─────────────┘
                 │ fetch()
┌────────────────▼─────────────┐
│   Go Backend API             │
│   (net/http handlers)        │
└────────────────┬─────────────┘
                 │ SQL
┌────────────────▼─────────────┐
│   MySQL Database             │
│   (ResearchPublicationTracker)
└──────────────────────────────┘
```

## Core Entities

- **Departments** - Academic departments/colleges
- **Researchers** - Faculty and researchers
- **Venues** - Journals, conferences, publishers
- **Publications** - Research articles and papers
- **Keywords** - Research topic tags
- **Authorship** - Author-publication relationships
- **Publication Keywords** - Publication-keyword associations

## Application Pages

1. **Dashboard** - Overview with statistics
2. **Departments** - Department management
3. **Researchers** - Researcher profiles and details
4. **Publications** - Publication browsing and search
5. **Reports** - Analytics and reporting
6. **Additional Pages** - Keywords, Authorship, Venue details

## UI Components

Each page includes:

- **Page Title** - Section heading
- **Search Bar** - Full-text search functionality
- **Add New Button** - Create new records
- **Data Table** - Record listing with pagination
- **Action Buttons**:
  - Edit button - Modify records
  - Delete button - Remove records
  - View Details button - Show full information

## Database Operations (CRUD)

| Operation  | Description              | SQL Command |
| ---------- | ------------------------ | ----------- |
| **C**reate | Insert new records       | INSERT      |
| **R**ead   | Select and retrieve data | SELECT      |
| **U**pdate | Modify existing records  | UPDATE      |
| **D**elete | Remove records           | DELETE      |

## Sample Data Fields

### Researcher

- Full Name
- Email
- Role (Professor, Assistant Professor, Lecturer, etc.)
- Department

### Publication

- Title
- Year
- Type (Journal Article, Conference Paper, Book Chapter, etc.)
- DOI
- Abstract
- Venue/Publisher

## Project Directory Structure

```
project/
│
├── backend/
│   ├── main.go                 # Server entry point
│   ├── db.go                   # Database connection
│   ├── models.go               # Data structures
│   ├── go.mod                  # Go module file
│   └── handlers/               # API endpoint handlers
│       ├── routes.go           # Route registration
│       ├── departments.go      # Department API
│       ├── researchers.go      # Researcher API
│       ├── venues.go           # Venue API
│       ├── publications.go     # Publication API
│       ├── keywords.go         # Keyword API
│       ├── authorship.go       # Authorship API
│       ├── publication_keywords.go  # Keyword mapping API
│       ├── reports.go          # Report generation
│       ├── models.go           # Handler data types
│       └── helpers.go          # Utility functions
│
├── frontend/
│   ├── index.html              # Dashboard/Home
│   ├── departments.html        # Departments page
│   ├── researchers.html        # Researchers page
│   ├── publications.html       # Publications page
│   ├── reports.html            # Reports page
│   ├── css/
│   │   └── style.css           # Global styles
│   └── js/
│       ├── api.js              # API client utilities
│       ├── departments.js      # Department page logic
│       ├── researchers.js      # Researcher page logic
│       ├── reports.js          # Report page logic
│
└── db/
    ├── docs/
    │   ├── README.md           # Project overview
    │   ├── Relational-Schema.md # Entity schemas
    │   ├── Constraints.md      # DB constraints
    │   ├── SQL-Queries-and-Results.md  # Query examples
    │   └── System-Requirements.md      # Requirements
    └── SQL/
        ├── 01_create_database_and_tables.sql
        ├── 02_insert_sample_data.sql
        └── 03_project_queries.sql
```

## Technology Stack

| Layer        | Technology                        |
| ------------ | --------------------------------- |
| **Frontend** | HTML5, CSS3, JavaScript (Vanilla) |
| **Backend**  | Go 1.24.2                         |
| **Server**   | Go net/http                       |
| **Database** | MySQL 5.7+                        |
| **Driver**   | github.com/go-sql-driver/mysql    |

## Quick Start

### 1. Database Setup

```bash
# Create database and tables
mysql -u root < db/SQL/01_create_database_and_tables.sql

# Insert sample data
mysql -u root < db/SQL/02_insert_sample_data.sql
```

### 2. Backend Setup

```bash
cd backend
go get github.com/go-sql-driver/mysql
```

### 3. Run Server

```bash
cd backend
go run .
```

Server starts on `http://localhost:8080`

### 4. Access Application

Open browser to `http://localhost:8080`

## API Endpoints Overview

### Core Resources

- `/api/departments` - Department CRUD
- `/api/researchers` - Researcher CRUD
- `/api/publications` - Publication CRUD
- `/api/venues` - Venue CRUD
- `/api/keywords` - Keyword CRUD
- `/api/authorship` - Authorship relationships
- `/api/publication-keywords` - Publication-keyword mappings

### Reports & Analytics

- `/api/reports/publications-with-authors`
- `/api/reports/publications-by-department`
- `/api/reports/publications-by-researcher`
- `/api/reports/publications-by-year`
- `/api/reports/publications-by-keyword`
- `/api/reports/publications-by-venue`
- `/api/reports/search-publications`

## Database Schema Highlights

- **Primary Keys**: Auto-incrementing integers for all tables
- **Foreign Keys**: Enforce referential integrity
- **Unique Constraints**: Email, DOI, KeywordText
- **Junction Tables**: Authorship and PublicationKeyword for many-to-many relationships
- **Normalization**: Third Normal Form (3NF)

## Key Features

✅ Full CRUD operations for all entities  
✅ Complex relationship management  
✅ Advanced reporting and analytics  
✅ Full-text search capability  
✅ RESTful API design  
✅ Responsive web interface  
✅ Type-safe Go backend  
✅ Database constraint enforcement
