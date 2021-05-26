package controller

import (
	authpb "auth/protocol-buffer/golang/auth"
	"auth/usecase"
	"github.com/labstack/echo"
)

type StudentController struct {
	authpb.AuthServer
	usecase.StudentUsecase
}

func NewStudentController(e *echo.Echo, us usecase.StudentUsecase) {
	handler := &StudentController{
		StudentUsecase: us,
	}

	e.POST("/login", handler.Login)
}

func (s *StudentController) Login(c echo.Context) error {

	return nil
}