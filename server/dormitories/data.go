package dormitories

import (
	"database/sql"
	"encoding/json"
	"fmt"
     "log"
     "strconv"
)

type Student struct {
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

func (s *Store) ListDormitories() ([]*Dormitory, error) {
	querySelect := `SELECT d.id, s.specialty, COUNT(s.specialty)
        FROM students s
        JOIN dormitories d ON d.id = s.dormitory_id
        GROUP BY d.id, s.specialty`
        rows, err := s.Db.Query(querySelect)
        if err != nil {
            return nil, err
        }

        defer rows.Close()

        var res []*Dormitory
        var dorm *Dormitory
        students := make(map[string]string)
        var spec string
        var count int64
        for rows.Next() {
            var curId int64
            if err := rows.Scan(&curId, &spec, &count); err != nil {
                return nil, err
            }

            if dorm == nil {
                dorm = new(Dormitory)
                dorm.Id = curId
            } else {
                if dorm.Id != curId {
                    s0, _ := json.Marshal(students)
                    s1 := json.RawMessage(s0)
                    dorm.StudentsSpecialtiesQ = &s1
                    res = append(res, dorm)

                    dorm = new(Dormitory)
                    dorm.Id = curId
                    students = make(map[string]string)
                }
            }

            students[spec] = strconv.FormatInt(count, 10)
        }
        if dorm != nil {
            s0, _ := json.Marshal(students)
            s1 := json.RawMessage(s0)
            dorm.StudentsSpecialtiesQ = &s1
            res = append(res, dorm)
        }
        if res == nil {
            res = make([]*Dormitory, 0)
        }
        return res, nil
}

func (s *Store) AddStudent(specialty string) error {
	if len(specialty) < 0 {
            return fmt.Errorf("specialty is not provided")
        }

        querySelect := `SELECT d.id
        FROM students s
        JOIN dormitories d ON d.id = s.dormitory_id
        WHERE specialty = $1
        GROUP BY d.id, s.specialty
        ORDER BY COUNT(s.specialty)
        LIMIT 1`
        rows, err := s.Db.Query(querySelect, specialty)
        if err != nil {
            return err
        }
        defer rows.Close()

        if !rows.Next() {
            log.Println("aaaa")
            return err
        }

        var id int64
        if err := rows.Scan(&id); err != nil {
            return err
        }
        log.Println(id)

        queryInsert := `INSERT INTO students (dormitory_id, specialty)
        VALUES ($1, $2)`
        _, err = s.Db.Exec(queryInsert, id, specialty)
        return err
}
