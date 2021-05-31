package usecase

import (
	"auth/model"
	"auth/repository/mocks"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLoginAuth(t *testing.T) {
	mockStudentRepo := new(mocks.StudentRepository)
	mockStudent := &model.Student{
		Gcn: "2318",
		Name: "조호원",
		Email: "hwc9169@gmail.com",
	}

	t.Run("success", func(t *testing.T){
		mockStudentRepo.On("GetByGcn", "2318").
			Return(mockStudent, nil).Once()

		su := NewStudentUsecase(mockStudentRepo)
		at, rt, err := su.LoginAuth(mockStudent)

		assert.NoError(t, err)
		assert.NotEmpty(t, rt)
		assert.NotEmpty(t, at)
	})

}
