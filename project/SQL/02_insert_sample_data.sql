USE ResearchPublicationTracker;

-- =========================
-- Insert Departments
-- =========================
INSERT INTO Department (DepartmentID, DepartmentName, College)
VALUES
(1, 'Computer Science', 'College of Information Technology'),
(2, 'Information Systems', 'College of Information Technology'),
(3, 'Computer Engineering', 'College of Engineering'),
(4, 'Electrical Engineering', 'College of Engineering'),
(5, 'Mechanical Engineering', 'College of Engineering'),
(6, 'Civil Engineering', 'College of Engineering'),
(7, 'Mathematics', 'College of Science'),
(8, 'Physics', 'College of Science'),
(9, 'Chemistry', 'College of Science'),
(10, 'Biology', 'College of Science'),
(11, 'Business Administration', 'College of Business Administration'),
(12, 'Accounting', 'College of Business Administration'),
(13, 'Finance', 'College of Business Administration'),
(14, 'Economics', 'College of Business Administration'),
(15, 'Architecture', 'College of Engineering'),
(16, 'Software Engineering', 'College of Information Technology'),
(17, 'Cybersecurity', 'College of Information Technology'),
(18, 'Data Science', 'College of Information Technology'),
(19, 'Renewable Energy Engineering', 'College of Engineering'),
(20, 'Environmental Science', 'College of Science');

-- =========================
-- Insert Researchers
-- Role must be: Faculty, Student, External Researcher
-- =========================
INSERT INTO Researcher (ResearcherID, FullName, Email, Role, DepartmentID)
VALUES
(1, 'Ahmed Ali Hassan', 'ahmed.hassan@uob.edu.bh', 'Faculty', 1),
(2, 'Fatima Salman Mohamed', 'fatima.salman@uob.edu.bh', 'Faculty', 2),
(3, 'Hussain Jassim Alawi', 'hussain.alawi@uob.edu.bh', 'Student', 1),
(4, 'Maryam Abdulrahman Noor', 'maryam.noor@uob.edu.bh', 'Student', 18),
(5, 'Ali Mahdi Abbas', 'ali.abbas@uob.edu.bh', 'Faculty', 4),
(6, 'Zainab Khalil Ahmed', 'zainab.ahmed@uob.edu.bh', 'Faculty', 17),
(7, 'Mohamed Yousif Isa', 'mohamed.isa@uob.edu.bh', 'Student', 16),
(8, 'Noor Hassan Khalifa', 'noor.khalifa@uob.edu.bh', 'Faculty', 3),
(9, 'Sara Faisal Mahmood', 'sara.mahmood@uob.edu.bh', 'Student', 7),
(10, 'Khalid Abdulla Nasser', 'khalid.nasser@uob.edu.bh', 'Faculty', 11),
(11, 'Layla Ebrahim Saleh', 'layla.saleh@uob.edu.bh', 'External Researcher', 20),
(12, 'Yusuf Hamad Rashid', 'yusuf.rashid@uob.edu.bh', 'Faculty', 19),
(13, 'Aisha Adel Mansoor', 'aisha.mansoor@uob.edu.bh', 'Student', 2),
(14, 'Sayed Mustafa Almoosawi', 'mustafa.almoosawi@uob.edu.bh', 'Faculty', 5),
(15, 'Mona Sami Darwish', 'mona.darwish@uob.edu.bh', 'Faculty', 13),
(16, 'Abdulla Redha Jawad', 'abdulla.jawad@uob.edu.bh', 'Student', 8),
(17, 'Reem Nabeel Faraj', 'reem.faraj@uob.edu.bh', 'External Researcher', 10),
(18, 'Hasan Fadel Amin', 'hasan.amin@uob.edu.bh', 'Faculty', 6),
(19, 'Ruqaya Yaqoob Saeed', 'ruqaya.saeed@uob.edu.bh', 'Student', 9),
(20, 'Omar Adel Kareem', 'omar.kareem@uob.edu.bh', 'Faculty', 12);

-- =========================
-- Insert Venues
-- VenueType must be: Journal, Conference, Book, Report, Other
-- =========================
INSERT INTO Venue (VenueID, VenueName, VenueType, Publisher)
VALUES
(1, 'International Journal of Artificial Intelligence Research', 'Journal', 'Springer'),
(2, 'GCC Conference on Computing and Technology', 'Conference', 'IEEE'),
(3, 'Journal of Cybersecurity and Digital Forensics', 'Journal', 'Elsevier'),
(4, 'Bahrain International Engineering Conference', 'Conference', 'University of Bahrain'),
(5, 'Renewable Energy Systems Journal', 'Journal', 'Taylor and Francis'),
(6, 'Middle East Data Science Conference', 'Conference', 'ACM'),
(7, 'University of Bahrain Research Reports', 'Report', 'University of Bahrain'),
(8, 'Journal of Information Systems Management', 'Journal', 'Emerald'),
(9, 'Smart Cities and IoT Symposium', 'Conference', 'IEEE'),
(10, 'International Journal of Software Engineering', 'Journal', 'Springer'),
(11, 'Applied Physics Research Journal', 'Journal', 'Elsevier'),
(12, 'Sustainable Environment Reports', 'Report', 'UNESCO'),
(13, 'Business Analytics and Finance Conference', 'Conference', 'ACM'),
(14, 'Journal of Mechanical Design and Manufacturing', 'Journal', 'ASME'),
(15, 'Civil Infrastructure Research Forum', 'Conference', 'ICE Publishing'),
(16, 'Book of Advanced Database Systems', 'Book', 'Pearson'),
(17, 'Journal of Educational Technology', 'Journal', 'Taylor and Francis'),
(18, 'International Conference on Renewable Energy', 'Conference', 'IEEE'),
(19, 'Health and Biology Research Journal', 'Journal', 'Nature Portfolio'),
(20, 'Chemistry and Materials Science Reports', 'Report', 'University of Bahrain');

-- =========================
-- Insert Publications
-- PublicationType must be: Journal, Conference, Book Chapter, Report
-- =========================
INSERT INTO Publication (PublicationID, Title, PublicationYear, PublicationType, DOI, Abstract, VenueID)
VALUES
(1, 'Artificial Intelligence Applications in Smart Education', 2024, 'Journal', '10.1000/rpt001', 'This study explores the use of artificial intelligence in smart education systems.', 1),
(2, 'Cybersecurity Awareness Among University Students', 2023, 'Conference', '10.1000/rpt002', 'This paper investigates cybersecurity awareness among university students.', 2),
(3, 'Machine Learning for Research Publication Classification', 2025, 'Journal', '10.1000/rpt003', 'This research applies machine learning techniques to classify academic publications.', 6),
(4, 'Database Design for Academic Research Tracking', 2024, 'Report', NULL, 'This report presents a relational database design for tracking academic research output.', 7),
(5, 'Blockchain-Based Verification of Research Publications', 2022, 'Journal', '10.1000/rpt005', 'This paper discusses blockchain methods for verifying research publications.', 3),
(6, 'Renewable Energy Adoption in Bahrain Universities', 2023, 'Conference', '10.1000/rpt006', 'This research studies renewable energy adoption in university campuses.', 18),
(7, 'Internet of Things Framework for Smart Laboratories', 2024, 'Conference', '10.1000/rpt007', 'This paper proposes an IoT framework for monitoring smart laboratories.', 9),
(8, 'Data Mining Techniques for Student Performance Analysis', 2021, 'Journal', '10.1000/rpt008', 'This study applies data mining methods to analyze student performance.', 8),
(9, 'Software Engineering Practices in Student Projects', 2025, 'Journal', '10.1000/rpt009', 'This research evaluates software engineering practices in university student projects.', 10),
(10, 'Optimization of Solar Panel Cleaning Systems', 2024, 'Conference', '10.1000/rpt010', 'This paper presents methods to improve solar panel cleaning systems.', 4),
(11, 'Statistical Modeling in Applied Physics Experiments', 2022, 'Journal', '10.1000/rpt011', 'This study uses statistical modeling to analyze physics experiment results.', 11),
(12, 'Environmental Impact of Plastic Waste Management', 2023, 'Report', NULL, 'This report examines environmental impacts related to plastic waste management.', 12),
(13, 'Business Intelligence for Financial Decision Making', 2024, 'Conference', '10.1000/rpt013', 'This paper studies the role of business intelligence in financial decision making.', 13),
(14, 'Mechanical Design Improvements for Manufacturing Efficiency', 2021, 'Journal', '10.1000/rpt014', 'This publication discusses mechanical design methods for improving manufacturing efficiency.', 14),
(15, 'Civil Infrastructure Risk Assessment Using Data Analytics', 2025, 'Conference', '10.1000/rpt015', 'This study applies data analytics to civil infrastructure risk assessment.', 15),
(16, 'Advanced Relational Database Systems for Research Management', 2022, 'Book Chapter', '10.1000/rpt016', 'This chapter explains relational database systems used for research management.', 16),
(17, 'Digital Learning Platforms and Student Engagement', 2024, 'Journal', '10.1000/rpt017', 'This research studies the relationship between digital learning platforms and student engagement.', 17),
(18, 'Energy Efficiency in Smart University Buildings', 2025, 'Conference', '10.1000/rpt018', 'This paper examines energy efficiency solutions for smart university buildings.', 18),
(19, 'Biological Data Analysis Using Machine Learning', 2023, 'Journal', '10.1000/rpt019', 'This study applies machine learning techniques to biological data analysis.', 19),
(20, 'Chemical Materials for Sustainable Engineering Applications', 2022, 'Report', NULL, 'This report discusses chemical materials used in sustainable engineering applications.', 20);

-- =========================
-- Insert Keywords
-- =========================
INSERT INTO Keyword (KeywordID, KeywordText)
VALUES
(1, 'Artificial Intelligence'),
(2, 'Cybersecurity'),
(3, 'Machine Learning'),
(4, 'Database Systems'),
(5, 'Research Tracking'),
(6, 'Blockchain'),
(7, 'Renewable Energy'),
(8, 'Internet of Things'),
(9, 'Data Mining'),
(10, 'Software Engineering'),
(11, 'Solar Energy'),
(12, 'Applied Physics'),
(13, 'Environmental Sustainability'),
(14, 'Plastic Waste'),
(15, 'Business Intelligence'),
(16, 'Financial Analytics'),
(17, 'Mechanical Design'),
(18, 'Civil Infrastructure'),
(19, 'Digital Learning'),
(20, 'Smart Buildings');

-- =========================
-- Insert Authorship Records
-- Many-to-many between Researcher and Publication
-- =========================
INSERT INTO Authorship (ResearcherID, PublicationID, AuthorOrder)
VALUES
(1, 1, 1),
(3, 1, 2),
(6, 2, 1),
(13, 2, 2),
(4, 3, 1),
(1, 3, 2),
(2, 4, 1),
(7, 4, 2),
(6, 5, 1),
(11, 5, 2),
(12, 6, 1),
(5, 6, 2),
(8, 7, 1),
(7, 7, 2),
(2, 8, 1),
(9, 8, 2),
(7, 9, 1),
(1, 9, 2),
(12, 10, 1),
(14, 10, 2),
(16, 11, 1),
(8, 11, 2),
(11, 12, 1),
(17, 12, 2),
(10, 13, 1),
(15, 13, 2),
(14, 14, 1),
(5, 14, 2),
(18, 15, 1),
(4, 15, 2),
(2, 16, 1),
(1, 16, 2),
(13, 17, 1),
(2, 17, 2),
(12, 18, 1),
(5, 18, 2),
(17, 19, 1),
(19, 19, 2),
(19, 20, 1),
(11, 20, 2);

-- =========================
-- Insert PublicationKeyword Records
-- Many-to-many between Publication and Keyword
-- =========================
INSERT INTO PublicationKeyword (PublicationID, KeywordID)
VALUES
(1, 1),
(1, 19),
(1, 3),
(2, 2),
(2, 19),
(3, 3),
(3, 5),
(3, 1),
(4, 4),
(4, 5),
(5, 6),
(5, 2),
(6, 7),
(6, 11),
(7, 8),
(7, 20),
(8, 9),
(8, 19),
(9, 10),
(9, 4),
(10, 11),
(10, 7),
(11, 12),
(11, 9),
(12, 13),
(12, 14),
(13, 15),
(13, 16),
(14, 17),
(14, 10),
(15, 18),
(15, 9),
(16, 4),
(16, 5),
(17, 19),
(17, 1),
(18, 20),
(18, 7),
(19, 3),
(19, 9),
(20, 13),
(20, 17);