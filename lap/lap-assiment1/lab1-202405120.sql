-- ITCS285 Database Management Systems - Lab #1
-- Name: hussain ali h. ali
-- Student ID: 202405120

-- 1. Create all tables with their related constraints.
CREATE DATABASE IF NOT EXISTS itcs285_lab1;
USE itcs285_lab1;

DROP TABLE IF EXISTS Account;
DROP TABLE IF EXISTS Branch;
DROP TABLE IF EXISTS Bank;
DROP TABLE IF EXISTS Client;

CREATE TABLE Client (
    ClientID INT PRIMARY KEY,
    FirstName VARCHAR(30) NOT NULL,
    LastName VARCHAR(30) NOT NULL,
    Phone VARCHAR(15) NOT NULL,
    Address VARCHAR(100) NOT NULL,
    Profession VARCHAR(30) NOT NULL
);

CREATE TABLE Bank (
    BankID INT PRIMARY KEY,
    BankName VARCHAR(50) NOT NULL
);

CREATE TABLE Branch (
    BranchID VARCHAR(10) PRIMARY KEY,
    BankID INT NOT NULL,
    BranchName VARCHAR(50) NOT NULL,
    BranchAddress VARCHAR(100) NOT NULL,
    CONSTRAINT fk_branch_bank
        FOREIGN KEY (BankID) REFERENCES Bank(BankID)
);

CREATE TABLE Account (
    AccountID VARCHAR(10) PRIMARY KEY,
    BranchID VARCHAR(10) NOT NULL,
    ClientID INT NOT NULL,
    Balance DECIMAL(12, 2) NOT NULL,
    CONSTRAINT fk_account_branch
        FOREIGN KEY (BranchID) REFERENCES Branch(BranchID),
    CONSTRAINT fk_account_client
        FOREIGN KEY (ClientID) REFERENCES Client(ClientID),
    CONSTRAINT chk_account_balance
        CHECK (Balance >= 0)
);

-- 2. Insert data into the different tables.

INSERT INTO Client (ClientID, FirstName, LastName, Phone, Address, Profession)
VALUES
    (10000, 'Ahmad', 'Abd Allah', '2134657', 'Manama', 'Professor'),
    (10001, 'Ghazi', 'Al-Amri', '2345676', 'Manama', 'Student'),
    (10002, 'Mohamed', 'Al-Hamed', '5467698', 'Manama', 'Trader'),
    (10003, 'Ali', 'Hussain', '6578790', 'Manama', 'Teller'),
    (10004, 'Amer', 'Al-Salah', '6578798', 'Manama', 'Professor'),
    (10005, 'Salah', 'Al-Kacem', '6579879', 'Manama', 'Engineer');

INSERT INTO Bank (BankID, BankName)
VALUES
    (30001, 'BBK'),
    (30002, 'NBB'),
    (30003, 'CITI');

INSERT INTO Branch (BranchID, BankID, BranchName, BranchAddress)
VALUES
    ('A100', 30003, 'CITYCENTER', 'Manama'),
    ('A101', 30003, 'HAMAD TOWN', 'HAMALAH'),
    ('R100', 30001, 'ZALLAQ', 'ZALLAQ'),
    ('R101', 30001, 'SEEF', 'Manama');

INSERT INTO Account (AccountID, BranchID, ClientID, Balance)
VALUES
    ('AC100', 'A100', 10000, 20000),
    ('AC101', 'A100', 10001, 120000),
    ('AC102', 'A101', 10002, 25300),
    ('AC103', 'R100', 10003, 26500),
    ('AC104', 'R100', 10004, 45600),
    ('AC105', 'R101', 10004, 56000);

-- 3. Modify table Bank to add a new attribute BankCode of type string
--    with exactly 8 characters.

ALTER TABLE Bank
ADD BankCode CHAR(8);

-- 4. Find clients with first names starting with A and having at least
--    3 characters.

SELECT *
FROM Client
WHERE FirstName LIKE 'A__%';
