-- ITCS285 Database Management Systems - Lab #3
-- Student ID: 202405120
-- Student Name: hussain ali h. ali
-- MySQL queries for the Lab #1 bank database.

USE itcs285_lab1;

-- 1. For each client, find his accounts with their branches.

SELECT
    c.ClientID,
    c.FirstName,
    c.LastName,
    a.AccountID,
    a.Balance,
    br.BranchID,
    br.BranchName,
    br.BranchAddress,
    b.BankName
FROM Client c
LEFT JOIN Account a ON c.ClientID = a.ClientID
LEFT JOIN Branch br ON a.BranchID = br.BranchID
LEFT JOIN Bank b ON br.BankID = b.BankID
ORDER BY
    c.ClientID,
    a.AccountID;

-- 2. Find clients who have only BBK accounts.

SELECT
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
      AND b.BankName = 'BBK'
)
AND NOT EXISTS (
    SELECT 1
    FROM Account a
    JOIN Branch br ON a.BranchID = br.BranchID
    JOIN Bank b ON br.BankID = b.BankID
    WHERE a.ClientID = c.ClientID
      AND b.BankName <> 'BBK'
);

-- 3. Find clients who have at least two accounts in different banks.

SELECT
    c.ClientID,
    c.FirstName,
    c.LastName,
    c.Phone,
    c.Address,
    c.Profession,
    COUNT(DISTINCT a.AccountID) AS AccountCount,
    COUNT(DISTINCT b.BankID) AS BankCount
FROM Client c
JOIN Account a ON c.ClientID = a.ClientID
JOIN Branch br ON a.BranchID = br.BranchID
JOIN Bank b ON br.BankID = b.BankID
GROUP BY
    c.ClientID,
    c.FirstName,
    c.LastName,
    c.Phone,
    c.Address,
    c.Profession
HAVING COUNT(DISTINCT a.AccountID) >= 2
   AND COUNT(DISTINCT b.BankID) >= 2;

-- 4. For each account, find details of clients with the highest salary.
--    The bank schema has no Salary column, so Balance is used as the amount.

SELECT
    a.AccountID,
    a.BranchID,
    a.Balance,
    c.ClientID,
    c.FirstName,
    c.LastName,
    c.Phone,
    c.Address,
    c.Profession
FROM Account a
JOIN Client c ON a.ClientID = c.ClientID
WHERE a.Balance = (
    SELECT MAX(Balance)
    FROM Account
);

-- 5. Find details of clients who do not have BBK nor CITI accounts.

SELECT
    c.ClientID,
    c.FirstName,
    c.LastName,
    c.Phone,
    c.Address,
    c.Profession
FROM Client c
WHERE NOT EXISTS (
    SELECT 1
    FROM Account a
    JOIN Branch br ON a.BranchID = br.BranchID
    JOIN Bank b ON br.BankID = b.BankID
    WHERE a.ClientID = c.ClientID
      AND b.BankName IN ('BBK', 'CITI')
);
