package usecase

import "auth/model"

type StudentUsecase interface {
	LoginAuth(*model.Student) (string, string, error)
}