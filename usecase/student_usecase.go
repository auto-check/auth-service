package usecase

import "github.com/auto-check/common-module/model"

type StudentUsecase interface {
	LoginAuth(*model.Student) (string, string, error)
}