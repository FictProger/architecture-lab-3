package dormitories

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"strconv"
)

type Student struct {
	DormId    int64  `json:"id"`
	Specialty string `json:"specialty"`
}

type Dormitory struct {
	Id                   int64            `json:"id"`
	StudentsSpecialtiesQ *json.RawMessage `json:"studentsCount"`
}

type Store struct {
	Db *sql.DB
}

func NewDormitories(db *sql.DB) *Store {
	return &Store{Db: db}
}

func (s *Store) RecomendDormitory(specity string) (Dormitory, error) {
	if len(specity) < 0 {
		return Dormitory{}, fmt.Errorf("specialty is not provided")
	}

	querySelect := `SELECT d.id, s.specialty, COUNT(s.specialty)
    FROM students s
    JOIN dormitories d ON d.id = s.dormitory_id
    WHERE d.id = (SELECT d.id
        FROM students s
        JOIN dormitories d ON d.id = s.dormitory_id
        WHERE specialty = $1
        GROUP BY d.id, s.specialty
        ORDER BY COUNT(s.specialty)
        LIMIT 1)
    GROUP BY d.id, s.specialty`
	rows, err := s.Db.Query(querySelect, specity)
	if err != nil {
		return Dormitory{}, err
	}
	defer rows.Close()

	var dorm Dormitory
	students := make(map[string]string)
	var spec string
	var count int64
	for rows.Next() {
		if err := rows.Scan(&dorm.Id, &spec, &count); err != nil {
			return Dormitory{}, err
		}
		students[spec] = strconv.FormatInt(count, 10)
	}
	s0, _ := json.Marshal(students)
	s1 := json.RawMessage(s0)
	dorm.StudentsSpecialtiesQ = &s1
	return dorm, nil
}

func (s *Store) AddStudent(dormId int64, specialty string) error {
	if len(specialty) < 0 {
		return fmt.Errorf("specialty is not provided")
	}

	queryInsert := `INSERT INTO students (dormitory_id, specialty)
        VALUES ($1, $2)`
	_, err := s.Db.Exec(queryInsert, dormId, specialty)
	return err
}
