package repository

import "database/sql"

type studentRepository struct {
	DB *sql.DB
}

func NewStudentRepository(db *sql.DB) StudentRepository {
	return studentRepository{
		DB: db,
	}
}






