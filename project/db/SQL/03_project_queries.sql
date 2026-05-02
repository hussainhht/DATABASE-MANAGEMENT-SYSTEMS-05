USE ResearchPublicationTracker;

-- Query 1: Publications with authors, departments, and venues
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

-- Query 2: Publications newer than the average publication year
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

-- Query 3: Number of publications per department
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

-- Query 4: Latest publications ordered by year and title
SELECT
    PublicationID,
    Title,
    PublicationYear,
    PublicationType
FROM Publication
ORDER BY PublicationYear DESC, Title ASC
LIMIT 5;

-- Query 5: Researchers who have not authored any publication
-- Query 5: Researchers who authored selected publications
SELECT
    r.ResearcherID,
    r.FullName,
    d.DepartmentName
FROM Researcher r, Department d
WHERE r.DepartmentID = d.DepartmentID
  AND r.ResearcherID IN (
      SELECT ResearcherID
      FROM Authorship
      WHERE PublicationID IN (1, 2, 3)
  )
ORDER BY r.FullName;
-- Query 6: Publications related to energy using LIKE
SELECT
    PublicationID,
    Title,
    PublicationYear,
    PublicationType
FROM Publication
WHERE Title LIKE '%Energy%'
   OR Abstract LIKE '%energy%'
ORDER BY PublicationYear DESC, Title;

-- Query 7: Overall publication year statistics
SELECT
    COUNT(*) AS TotalPublications,
    MIN(PublicationYear) AS EarliestYear,
    MAX(PublicationYear) AS LatestYear,
    ROUND(AVG(PublicationYear), 1) AS AverageYear
FROM Publication;

-- Query 8: Researchers with more than two publications
SELECT
    r.ResearcherID,
    r.FullName,
    COUNT(a.PublicationID) AS PublicationCount
FROM Researcher r, Authorship a
WHERE r.ResearcherID = a.ResearcherID
GROUP BY r.ResearcherID, r.FullName
HAVING COUNT(a.PublicationID) > 2
ORDER BY PublicationCount DESC, r.FullName;

-- Query 9: Publication age calculated from the project year 2026
SELECT
    PublicationID,
    Title,
    PublicationYear,
    (2026 - PublicationYear) AS PublicationAgeYears
FROM Publication
ORDER BY PublicationAgeYears DESC, PublicationID
LIMIT 5;

-- Query 10: Complete publication report using most database relations
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
