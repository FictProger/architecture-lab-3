DROP TABLE IF EXISTS dormitories CASCADE;
CREATE TABLE dormitories (
    id SERIAL PRIMARY KEY
);

DROP TABLE IF EXISTS students;
CREATE TABLE students (
    id SERIAL PRIMARY KEY,
    dormitory_id INT REFERENCES dormitories NOT NULL,
    specialty VARCHAR(50) NOT NULL
);

INSERT INTO dormitories DEFAULT VALUES;
INSERT INTO dormitories DEFAULT VALUES;

INSERT INTO students (dormitory_id, specialty)
VALUES (1, 'biology'),
    (1, 'biology'),
    (1, 'biology'),
    (1, 'computerScience'),
    (1, 'computerScience'),
    (1, 'computerScience'),
    (1, 'computerScience'),
    (1, 'computerScience'),
    (1, 'literature'),
    (1, 'literature'),
    (2, 'biology'),
    (2, 'computerScience'),
    (2, 'computerScience'),
    (2, 'literature'),
    (2, 'literature'),
    (2, 'literature'),
    (2, 'literature');

SELECT * FROM dormitories;
SELECT * FROM students;
