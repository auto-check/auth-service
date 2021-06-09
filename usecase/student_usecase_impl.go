package usecase

import (
	"auth/repository"
	"github.com/auto-check/common-module/jwt"
	"github.com/auto-check/common-module/model"
)

type studentUsecase struct {
	sr repository.StudentRepository
}

func NewStudentUsecase(sr repository.StudentRepository) StudentUsecase{
	return &studentUsecase{
		sr: sr,
	}
}

func (su *studentUsecase) LoginAuth(s *model.Student) (string, string, error) {
	s, err := su.sr.GetByGcn(s.Gcn)
	if err != nil {
		return "", "", err
	}

	at, rt, err := jwt.GenerateToken(s.ID)
	if err != nil {
		return "", "", err
	}

	return at, rt, nil
}