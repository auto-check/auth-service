package repository

import "github.com/auto-check/common-module/model"

type StudentRepository interface {
	Store(*model.Student) error
	GetByGcn(gcn string) (*model.Student, error)
}