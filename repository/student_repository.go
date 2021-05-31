package repository

import "auth/model"

type StudentRepository interface {
	Store(*model.Student) error
	GetByGcn(gcn string) (*model.Student, error)
}