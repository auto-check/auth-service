package repository

import (
	"auth/model"
	sqlmock "github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"regexp"
	"testing"
)

func TestGetByGcn(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	rows := sqlmock.NewRows([]string{"id", "gcn", "name", "email"}).
		AddRow(1, "2318", "조호원", "hwc9169@gmail.com")

	query := regexp.QuoteMeta(`SELECT id, gcn, name, email FROM student WHERE gcn = ?`)
	studentGcn := "2318"
	mock.ExpectQuery(query).WithArgs("2318").WillReturnRows(rows)

	sr := NewStudentRepository(db)

	student, err := sr.GetByGcn(studentGcn)
	assert.NoError(t, err)
	assert.NotNil(t, student)
}

func TestStore(t *testing.T) {
	s := &model.Student{
		Gcn: "2318",
		Name: "조호원",
		Email: "hwc9169@gmail.com",
	}

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	query := regexp.QuoteMeta(`INSERT INTO student SET gcn=?, name=?, email=?`)
	prep := mock.ExpectPrepare(query)
	prep.ExpectExec().WithArgs(s.Gcn, s.Name, s.Email).WillReturnResult(sqlmock.NewResult(1, 1))

	sr := NewStudentRepository(db)

	err = sr.Store(s)
	assert.NoError(t, err)
	assert.Equal(t, int64(1), s.ID)
}
