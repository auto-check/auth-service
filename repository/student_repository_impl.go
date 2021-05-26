package repository

import (
	"auth/model"
	"database/sql"
	log "github.com/sirupsen/logrus"
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
		log.Error(err)
		return nil, err
	}

	defer func() {
		errRow := rows.Close()
		if errRow != nil {
			log.Error(errRow)
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
			log.Error(err)
			return nil, err
		}
		result = append(result, s)
	}

	return result, nil
}

func (sr *studentRepository) Store(student *model.Student) error {
	return nil
}

func (sr *studentRepository) IsExistByGcn(gcn string) (bool, error){
	query := `SELECT id, gcn, name, email FROM student WHERE gcn = ?`

	res, err := sr.fetch(query, gcn)
	if err != nil {
		return false, err
	}

	if len(res) < 1 {
		return false, nil
	}

	return true, nil
}






