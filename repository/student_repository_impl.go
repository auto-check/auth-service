package repository

import (
	"auth/model"
	"database/sql"
	"fmt"
	log "github.com/sirupsen/logrus"
	"runtime/debug"
)

type studentRepository struct {
	DB *sql.DB
}

func NewStudentRepository(db *sql.DB) StudentRepository {
	return &studentRepository{
		DB: db,
	}
}

func (sr *studentRepository) fetch(query string, args ...interface{}) ([]model.Student, error){
	rows, err := sr.DB.Query(query, args...)
	if err != nil {
		log.Error(fmt.Sprintf("Error %s\n%s", err, debug.Stack()))
		return nil, err
	}

	defer func() {
		err = rows.Close()
		if err != nil {
			log.Error(fmt.Sprintf("Error %s\n%s", err, debug.Stack()))
		}
	}()

	result := make([]model.Student, 0)
	for rows.Next() {
		s := model.Student{}
		err = rows.Scan(
			&s.ID,
			&s.Name,
			&s.Gcn,
			&s.Email,
			)
		if err != nil {
			log.Error(fmt.Sprintf("Error %s\n%s", err, debug.Stack()))
			return nil, err
		}
		result = append(result, s)
	}

	return result, nil
}

func (sr *studentRepository) Store(s *model.Student) error {
	query := `INSERT INTO student SET gcn=?, name=?, email=?`
	stmt, err := sr.DB.Prepare(query)
	if err != nil {
		log.Errorf(fmt.Sprintf("Error %s\n%s", err, debug.Stack()))
		return err
	}

	res, err := stmt.Exec(s.Gcn, s.Name, s.Email)
	if err != nil {
		log.Errorf(fmt.Sprintf("Error %s\n%s", err, debug.Stack()))
		return err
	}

	lastID, err := res.LastInsertId()
	if err != nil {
		log.Errorf(fmt.Sprintf("Error %s\n%s", err, debug.Stack()))
		return err
	}
	s.ID = lastID

	return nil
}

func (sr *studentRepository) GetByGcn(gcn string) (*model.Student, error){
	query := `SELECT id, gcn, name, email FROM student WHERE gcn = ?`

	res, err := sr.fetch(query, gcn)
	if err != nil {
		return nil, err
	}

	if len(res) < 1 {
		return nil, fmt.Errorf("No response data for student gcn %s", gcn)
	}

	return &res[0], nil
}






