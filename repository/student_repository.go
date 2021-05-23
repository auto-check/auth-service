package repository

import "auth/model"

type StudentRepository interface {
	Store(*model.Student) error
}