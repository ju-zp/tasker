package project

import (
	"database/sql"
	"fmt"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jinzhu/gorm"
	"github.com/ju-zp/tasker/svc/models"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"regexp"
	"testing"
)

type Suite struct {
	suite.Suite
	DB   *gorm.DB
	mock sqlmock.Sqlmock

	repository *Repository
	project     *models.Project
}

func (s *Suite) SetupSuite() {
	var (
		db  *sql.DB
		err error
	)

	db, s.mock, err = sqlmock.New()
	require.NoError(s.T(), err)

	s.DB, err = gorm.Open("postgres", db)
	require.NoError(s.T(), err)

	s.DB.LogMode(true)

	s.repository = CreateRepository(s.DB)
}

func (s *Suite) AfterTest(_, _ string) {
	require.NoError(s.T(), s.mock.ExpectationsWereMet())
}

func TestInit(t *testing.T) {
	suite.Run(t, new(Suite))
}
//
//func (s *Suite) Test_repository_Get() {
//	var (
//		id   = "1"
//		title = "test-name"
//	)
//
//	s.mock.ExpectQuery(regexp.QuoteMeta(
//		`SELECT * FROM "project" WHERE (id = $1)`)).
//		WithArgs(id).
//		WillReturnRows(sqlmock.NewRows([]string{"id", "title"}).
//			AddRow(id, title))
//
//	res, err := s.repository.Get(id)
//
//		fmt.Println(err)
//
//	require.NoError(s.T(), err)
//	require.Nil(s.T(), deep.Equal(&models.Project{ID: 1, Title: &title}, res))
//}

func (s *Suite) Test_repository_Create() {
	var (
		title = "test-name"
		description = "here"
	)

	mockedRow := sqlmock.NewRows([]string{"id"}).AddRow("1")

	s.mock.ExpectBegin()
	s.mock.ExpectQuery(regexp.QuoteMeta(`INSERT INTO "projects" ("description","title") VALUES ($1,$2) RETURNING "projects"."id"`)).WithArgs(description,title).WillReturnRows(mockedRow)
	s.mock.ExpectCommit()

	err := s.repository.Create(title)

	require.NoError(s.T(), err)
}