-- Student ID: 202405120
-- Student Name: hussain ali h. ali

CREATE DATABASE ResearchPublicationTracker;
USE ResearchPublicationTracker;

-- Department table
CREATE TABLE Department (
    DepartmentID INT AUTO_INCREMENT PRIMARY KEY,
    DepartmentName VARCHAR(100) NOT NULL,
    College VARCHAR(100) NOT NULL
);

-- Researcher table
CREATE TABLE Researcher (
    ResearcherID INT AUTO_INCREMENT PRIMARY KEY,
    FullName VARCHAR(150) NOT NULL,
    Email VARCHAR(150) NOT NULL UNIQUE,
    Role VARCHAR(50) NOT NULL,
    DepartmentID INT NOT NULL,
    FOREIGN KEY (DepartmentID) REFERENCES Department(DepartmentID),
    CHECK (Role IN ('Faculty', 'Student', 'External Researcher'))
);

-- Venue table
CREATE TABLE Venue (
    VenueID INT AUTO_INCREMENT PRIMARY KEY,
    VenueName VARCHAR(150) NOT NULL,
    VenueType VARCHAR(50) NOT NULL,
    Publisher VARCHAR(150),
    CHECK (VenueType IN ('Journal', 'Conference', 'Book', 'Report', 'Other'))
);

-- Publication table
CREATE TABLE Publication (
    PublicationID INT AUTO_INCREMENT PRIMARY KEY,
    Title VARCHAR(255) NOT NULL,
    PublicationYear INT NOT NULL,
    PublicationType VARCHAR(50) NOT NULL,
    DOI VARCHAR(100) UNIQUE,
    Abstract TEXT,
    VenueID INT NOT NULL,
    FOREIGN KEY (VenueID) REFERENCES Venue(VenueID),
    CHECK (PublicationYear BETWEEN 1900 AND 2026),
    CHECK (PublicationType IN ('Journal', 'Conference', 'Book Chapter', 'Report'))
);

-- Keyword table
CREATE TABLE Keyword (
    KeywordID INT AUTO_INCREMENT PRIMARY KEY,
    KeywordText VARCHAR(100) NOT NULL UNIQUE
);

-- Authorship table
CREATE TABLE Authorship (
    ResearcherID INT NOT NULL,
    PublicationID INT NOT NULL,
    AuthorOrder INT NOT NULL,
    PRIMARY KEY (ResearcherID, PublicationID),
    FOREIGN KEY (ResearcherID) REFERENCES Researcher(ResearcherID),
    FOREIGN KEY (PublicationID) REFERENCES Publication(PublicationID),
    CHECK (AuthorOrder > 0)
);

-- PublicationKeyword table
CREATE TABLE PublicationKeyword (
    PublicationID INT NOT NULL,
    KeywordID INT NOT NULL,
    PRIMARY KEY (PublicationID, KeywordID),
    FOREIGN KEY (PublicationID) REFERENCES Publication(PublicationID),
    FOREIGN KEY (KeywordID) REFERENCES Keyword(KeywordID)
);