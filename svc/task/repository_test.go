package task

import (
	"database/sql"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/go-test/deep"
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
	todo     *models.Todo
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

func (s *Suite)Test_Create() {
	var (
		id = "1"
		title = "test-task"
		projectId =int64(5)
	)

	mockedRow :=  sqlmock.NewRows([]string{"id"}).
		AddRow(id)

	s.mock.ExpectBegin()
	s.mock.ExpectQuery(regexp.QuoteMeta(`INSERT INTO "tasks" ("done","project_id","title") VALUES ($1,$2,$3) RETURNING "tasks"."id"`)).
		WithArgs(false, projectId, title).
		WillReturnRows(mockedRow)
	s.mock.ExpectCommit()

	err := s.repository.Create(&title, projectId)

	require.NoError(s.T(), err)
}

func (s *Suite)Test_Find() {
	var (
		id = "1"
		title = "test-task"
		projectId = int64(5)
		done = false
	)

	mockedRow :=  sqlmock.NewRows([]string{"id", "title", "done", "project_id"}).
		AddRow(id, title, done, projectId)

	s.mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "tasks" WHERE (id = $1)`)).
		WithArgs(id).
		WillReturnRows(mockedRow)

	res, err := s.repository.Find(id)

	require.NoError(s.T(), err)
	require.Nil(s.T(), deep.Equal(&models.Task{ID: 1, Title: &title, ProjectID: projectId, Done: done}, res))

}
