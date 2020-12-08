package dormitories

import (
	"database/sql"
	"encoding/json"
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
	res := []*Dormitory{}
	return res, nil
}

func (s *Store) AddStudent(specialty string) error {

	return err
}
