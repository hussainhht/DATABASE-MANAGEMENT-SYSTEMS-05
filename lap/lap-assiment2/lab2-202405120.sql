-- ITCS285 Database Management Systems - Lab #2
-- Student ID: 202405120
-- Student Name: hussain ali h. ali
-- These queries assume the Lab #1 bank tables and data already exist.

-- 1. Find Account ID, first name, last name, bank name, branch address,
--    and account balance of each client.

SELECT
    a.AccountID,
    c.FirstName,
    c.LastName,
    b.BankName,
    br.BranchAddress,
    a.Balance
FROM Account a
JOIN Client c ON a.ClientID = c.ClientID
JOIN Branch br ON a.BranchID = br.BranchID
JOIN Bank b ON br.BankID = b.BankID;

-- 2. Find clients who have accounts at CITI bank but not at HSBC bank.

SELECT DISTINCT
    c.ClientID,
    c.FirstName,
    c.LastName,
    c.Phone,
    c.Address,
    c.Profession
FROM Client c
WHERE EXISTS (
    SELECT 1
    FROM Account a
    JOIN Branch br ON a.BranchID = br.BranchID
    JOIN Bank b ON br.BankID = b.BankID
    WHERE a.ClientID = c.ClientID
      AND b.BankName = 'CITI'
)
AND NOT EXISTS (
    SELECT 1
    FROM Account a
    JOIN Branch br ON a.BranchID = br.BranchID
    JOIN Bank b ON br.BankID = b.BankID
    WHERE a.ClientID = c.ClientID
      AND b.BankName = 'HSBC'
);

-- 3. Find most paid Professor at NBB bank.
-- Add a branch for NBB

-- First, we need to insert the NBB bank and its branch if they don't already exist.
INSERT INTO Branch (BranchID, BankID, BranchName, BranchAddress)
VALUES ('N100', 30002, 'NBB MAIN', 'Manama');

-- Add an account for a Professor at NBB
INSERT INTO Account (AccountID, BranchID, ClientID, Balance)
VALUES ('AC106', 'N100', 10000, 70000);

SELECT
    c.ClientID,
    c.FirstName,
    c.LastName,
    a.Balance
FROM Client c
JOIN Account a ON c.ClientID = a.ClientID
JOIN Branch br ON a.BranchID = br.BranchID
JOIN Bank b ON br.BankID = b.BankID
WHERE c.Profession = 'Professor'
  AND b.BankName = 'NBB'
  AND a.Balance = (
      SELECT MAX(a2.Balance)
      FROM Client c2
      JOIN Account a2 ON c2.ClientID = a2.ClientID
      JOIN Branch br2 ON a2.BranchID = br2.BranchID
      JOIN Bank b2 ON br2.BankID = b2.BankID
      WHERE c2.Profession = 'Professor'
        AND b2.BankName = 'NBB'
  );

-- 4. Find number of students accounts at CITI bank.

SELECT
    COUNT(a.AccountID) AS StudentAccountCount
FROM Client c
JOIN Account a ON c.ClientID = a.ClientID
JOIN Branch br ON a.BranchID = br.BranchID
JOIN Bank b ON br.BankID = b.BankID
WHERE c.Profession = 'Student'
  AND b.BankName = 'CITI';

-- 5. Find sum of balances in each branch.

SELECT
    br.BranchID,
    br.BranchName,
    COALESCE(SUM(a.Balance), 0) AS TotalBalance
FROM Branch br
LEFT JOIN Account a ON br.BranchID = a.BranchID
GROUP BY
    br.BranchID,
    br.BranchName;

