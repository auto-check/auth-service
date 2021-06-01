package mocks

import (
	"auth/model"
	"github.com/stretchr/testify/mock"
)

package mocks

import (
"auth/model"
"github.com/stretchr/testify/mock"
)

type StudentUsecase struct {
	mock.Mock
}

func (_m *StudentUsecase) LoginAuth(s *model.Student) (string, string, error) {
	ret := _m.Called(s)

	var r0 string
	if rf, ok := ret.Get(0).(func(student *model.Student) string); ok {
		r0 = rf(s)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 string
	if rf, ok := ret.Get(0).(func(student *model.Student) string); ok {
		r1 = rf(s)
	} else {
		r1 = ret.Get(0).(string)
	}

	var r2 error
	if rf, ok := ret.Get(0).(func(student *model.Student) error); ok {
		r2 = rf(s)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}