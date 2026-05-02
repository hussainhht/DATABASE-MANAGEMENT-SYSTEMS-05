# SQL Queries and Results

## Query 1: Publications with Authors, Departments, and Venues

Purpose:
This query lists selected publications with their authors, author order, department, and publication venue. It demonstrates a multiple-table join.

SQL Query:
```sql
SELECT
    p.PublicationID,
    p.Title,
    r.FullName AS AuthorName,
    a.AuthorOrder,
    d.DepartmentName,
    v.VenueName
FROM Publication p, Authorship a, Researcher r, Department d, Venue v
WHERE p.PublicationID = a.PublicationID
  AND a.ResearcherID = r.ResearcherID
  AND r.DepartmentID = d.DepartmentID
  AND p.VenueID = v.VenueID
  AND p.PublicationID IN (1, 2, 3)
ORDER BY p.PublicationID, a.AuthorOrder;
```

Expected Result:

| PublicationID | Title | AuthorName | AuthorOrder | DepartmentName | VenueName |
| --- | --- | --- | --- | --- | --- |
| 1 | Artificial Intelligence Applications in Smart Education | Ahmed Ali Hassan | 1 | Computer Science | International Journal of Artificial Intelligence Research |
| 1 | Artificial Intelligence Applications in Smart Education | Hussain Jassim Alawi | 2 | Computer Science | International Journal of Artificial Intelligence Research |
| 2 | Cybersecurity Awareness Among University Students | Zainab Khalil Ahmed | 1 | Cybersecurity | GCC Conference on Computing and Technology |
| 2 | Cybersecurity Awareness Among University Students | Aisha Adel Mansoor | 2 | Information Systems | GCC Conference on Computing and Technology |
| 3 | Machine Learning for Research Publication Classification | Maryam Abdulrahman Noor | 1 | Data Science | Middle East Data Science Conference |
| 3 | Machine Learning for Research Publication Classification | Ahmed Ali Hassan | 2 | Computer Science | Middle East Data Science Conference |

## Query 2: Publications Newer Than the Average Publication Year

Purpose:
This query finds publications that were published after the average publication year in the database. It demonstrates a subquery.

SQL Query:
```sql
SELECT
    PublicationID,
    Title,
    PublicationYear,
    PublicationType
FROM Publication
WHERE PublicationYear > (
    SELECT AVG(PublicationYear)
    FROM Publication
)
ORDER BY PublicationYear DESC, PublicationID
LIMIT 5;
```

Expected Result:

| PublicationID | Title | PublicationYear | PublicationType |
| --- | --- | --- | --- |
| 3 | Machine Learning for Research Publication Classification | 2025 | Journal |
| 9 | Software Engineering Practices in Student Projects | 2025 | Journal |
| 15 | Civil Infrastructure Risk Assessment Using Data Analytics | 2025 | Conference |
| 18 | Energy Efficiency in Smart University Buildings | 2025 | Conference |
| 1 | Artificial Intelligence Applications in Smart Education | 2024 | Journal |

## Query 3: Number of Publications per Department

Purpose:
This query counts the number of distinct publications connected to each department. It demonstrates GROUP BY.

SQL Query:
```sql
SELECT
    d.DepartmentName,
    COUNT(DISTINCT p.PublicationID) AS PublicationCount
FROM Department d, Researcher r, Authorship a, Publication p
WHERE d.DepartmentID = r.DepartmentID
  AND r.ResearcherID = a.ResearcherID
  AND a.PublicationID = p.PublicationID
GROUP BY d.DepartmentID, d.DepartmentName
ORDER BY PublicationCount DESC, d.DepartmentName
LIMIT 5;
```

Expected Result:

| DepartmentName | PublicationCount |
| --- | --- |
| Information Systems | 5 |
| Computer Science | 4 |
| Electrical Engineering | 3 |
| Environmental Science | 3 |
| Renewable Energy Engineering | 3 |

## Query 4: Latest Publications Ordered by Year and Title

Purpose:
This query displays the newest publications first and sorts publications from the same year alphabetically by title. It demonstrates ORDER BY.

SQL Query:
```sql
SELECT
    PublicationID,
    Title,
    PublicationYear,
    PublicationType
FROM Publication
ORDER BY PublicationYear DESC, Title ASC
LIMIT 5;
```

Expected Result:

| PublicationID | Title | PublicationYear | PublicationType |
| --- | --- | --- | --- |
| 15 | Civil Infrastructure Risk Assessment Using Data Analytics | 2025 | Conference |
| 18 | Energy Efficiency in Smart University Buildings | 2025 | Conference |
| 3 | Machine Learning for Research Publication Classification | 2025 | Journal |
| 9 | Software Engineering Practices in Student Projects | 2025 | Journal |
| 1 | Artificial Intelligence Applications in Smart Education | 2024 | Journal |

## Query 5: Researchers Who Have Not Authored Any Publication

Purpose:
This query finds researchers who do not appear in the Authorship table. It demonstrates NOT IN.

SQL Query:
```sql
SELECT
    r.ResearcherID,
    r.FullName,
    d.DepartmentName
FROM Researcher r, Department d
WHERE r.DepartmentID = d.DepartmentID
  AND r.ResearcherID NOT IN (
      SELECT ResearcherID
      FROM Authorship
  )
ORDER BY r.FullName;
```

Expected Result:

| ResearcherID | FullName | DepartmentName |
| --- | --- | --- |
| 20 | Omar Adel Kareem | Accounting |

## Query 6: Publications Related to Energy

Purpose:
This query searches for publications that contain the word "Energy" in the title or abstract. It demonstrates LIKE.

SQL Query:
```sql
SELECT
    PublicationID,
    Title,
    PublicationYear,
    PublicationType
FROM Publication
WHERE Title LIKE '%Energy%'
   OR Abstract LIKE '%energy%'
ORDER BY PublicationYear DESC, Title;
```

Expected Result:

| PublicationID | Title | PublicationYear | PublicationType |
| --- | --- | --- | --- |
| 18 | Energy Efficiency in Smart University Buildings | 2025 | Conference |
| 6 | Renewable Energy Adoption in Bahrain Universities | 2023 | Conference |

## Query 7: Overall Publication Year Statistics

Purpose:
This query summarizes the publication years in the database. It demonstrates aggregate functions such as COUNT, MIN, MAX, and AVG.

SQL Query:
```sql
SELECT
    COUNT(*) AS TotalPublications,
    MIN(PublicationYear) AS EarliestYear,
    MAX(PublicationYear) AS LatestYear,
    ROUND(AVG(PublicationYear), 1) AS AverageYear
FROM Publication;
```

Expected Result:

| TotalPublications | EarliestYear | LatestYear | AverageYear |
| --- | --- | --- | --- |
| 20 | 2021 | 2025 | 2023.3 |

## Query 8: Researchers with More Than Two Publications

Purpose:
This query lists researchers who authored more than two publications. It demonstrates HAVING.

SQL Query:
```sql
SELECT
    r.ResearcherID,
    r.FullName,
    COUNT(a.PublicationID) AS PublicationCount
FROM Researcher r, Authorship a
WHERE r.ResearcherID = a.ResearcherID
GROUP BY r.ResearcherID, r.FullName
HAVING COUNT(a.PublicationID) > 2
ORDER BY PublicationCount DESC, r.FullName;
```

Expected Result:

| ResearcherID | FullName | PublicationCount |
| --- | --- | --- |
| 1 | Ahmed Ali Hassan | 4 |
| 2 | Fatima Salman Mohamed | 4 |
| 5 | Ali Mahdi Abbas | 3 |
| 11 | Layla Ebrahim Saleh | 3 |
| 7 | Mohamed Yousif Isa | 3 |
| 12 | Yusuf Hamad Rashid | 3 |

## Query 9: Publication Age in Years

Purpose:
This query calculates the age of each publication using the project year 2026. It demonstrates a calculated field.

SQL Query:
```sql
SELECT
    PublicationID,
    Title,
    PublicationYear,
    (2026 - PublicationYear) AS PublicationAgeYears
FROM Publication
ORDER BY PublicationAgeYears DESC, PublicationID
LIMIT 5;
```

Expected Result:

| PublicationID | Title | PublicationYear | PublicationAgeYears |
| --- | --- | --- | --- |
| 8 | Data Mining Techniques for Student Performance Analysis | 2021 | 5 |
| 14 | Mechanical Design Improvements for Manufacturing Efficiency | 2021 | 5 |
| 5 | Blockchain-Based Verification of Research Publications | 2022 | 4 |
| 11 | Statistical Modeling in Applied Physics Experiments | 2022 | 4 |
| 16 | Advanced Relational Database Systems for Research Management | 2022 | 4 |

## Query 10: Complete Publication Report

Purpose:
This query produces a detailed publication report showing publication, venue, authors, departments, and keywords. It involves most of the database relations.

SQL Query:
```sql
SELECT
    p.PublicationID,
    p.Title,
    p.PublicationYear,
    v.VenueName,
    GROUP_CONCAT(DISTINCT r.FullName ORDER BY a.AuthorOrder SEPARATOR ', ') AS Authors,
    GROUP_CONCAT(DISTINCT d.DepartmentName ORDER BY d.DepartmentName SEPARATOR ', ') AS Departments,
    GROUP_CONCAT(DISTINCT k.KeywordText ORDER BY k.KeywordText SEPARATOR ', ') AS Keywords
FROM Department d, Researcher r, Authorship a, Publication p, Venue v, PublicationKeyword pk, Keyword k
WHERE d.DepartmentID = r.DepartmentID
  AND r.ResearcherID = a.ResearcherID
  AND a.PublicationID = p.PublicationID
  AND p.VenueID = v.VenueID
  AND p.PublicationID = pk.PublicationID
  AND pk.KeywordID = k.KeywordID
GROUP BY p.PublicationID, p.Title, p.PublicationYear, v.VenueName
ORDER BY p.PublicationID
LIMIT 5;
```

Expected Result:

| PublicationID | Title | PublicationYear | VenueName | Authors | Departments | Keywords |
| --- | --- | --- | --- | --- | --- | --- |
| 1 | Artificial Intelligence Applications in Smart Education | 2024 | International Journal of Artificial Intelligence Research | Ahmed Ali Hassan, Hussain Jassim Alawi | Computer Science | Artificial Intelligence, Digital Learning, Machine Learning |
| 2 | Cybersecurity Awareness Among University Students | 2023 | GCC Conference on Computing and Technology | Zainab Khalil Ahmed, Aisha Adel Mansoor | Cybersecurity, Information Systems | Cybersecurity, Digital Learning |
| 3 | Machine Learning for Research Publication Classification | 2025 | Middle East Data Science Conference | Maryam Abdulrahman Noor, Ahmed Ali Hassan | Computer Science, Data Science | Artificial Intelligence, Machine Learning, Research Tracking |
| 4 | Database Design for Academic Research Tracking | 2024 | University of Bahrain Research Reports | Fatima Salman Mohamed, Mohamed Yousif Isa | Information Systems, Software Engineering | Database Systems, Research Tracking |
| 5 | Blockchain-Based Verification of Research Publications | 2022 | Journal of Cybersecurity and Digital Forensics | Zainab Khalil Ahmed, Layla Ebrahim Saleh | Cybersecurity, Environmental Science | Blockchain, Cybersecurity |
