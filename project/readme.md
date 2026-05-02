Browser Pages
HTML / CSS / JS

        ↓ fetch()

Go Backend API
net/http handlers

        ↓ SQL

MySQL Database
ResearchPublicationTracker


Departments
Researchers
Venues
Publications
Keywords
Authorship
Publication Keywords
Reports / Queries


1. Dashboard
2. Departments
3. Researchers
4. Venues
5. Publications
6. Keywords
7. Authorship
8. Publication Keywords
9. Reports / SQL Queries
10. About Project



Page Title
Search bar
Add New button
Table showing records
Edit button
Delete button
View Details button


C = Create = Insert
R = Read = Select
U = Update = Modify
D = Delete

Full Name
Email
Role
Department



research-publication-tracker/
│
├── backend/
│   ├── main.go
│   ├── db.go
│   ├── models.go
│   ├── handlers/
│   │   ├── departments.go
│   │   ├── researchers.go
│   │   ├── venues.go
│   │   ├── publications.go
│   │   ├── keywords.go
│   │   ├── authorship.go
│   │   └── reports.go
│   └── go.mod
│
├── frontend/
│   ├── index.html
│   ├── researchers.html
│   ├── departments.html
│   ├── publications.html
│   ├── reports.html
│   ├── css/
│   │   └── style.css
│   └── js/
│       ├── api.js
│       ├── researchers.js
│       ├── departments.js
│       └── reports.js
│
└── database/
    ├── create_tables.sql
    ├── insert_data.sql
    └── queries.sql