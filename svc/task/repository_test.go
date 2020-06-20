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

func (s *Suite )Test_Create() {
	var (
		id = "1"
		title = "test-task"
		projectId = int64(5)
		done = false
	)

	mockedRow :=  sqlmock.NewRows([]string{"id"}).
		AddRow(id)

	s.mock.ExpectBegin()
	s.mock.ExpectQuery(regexp.QuoteMeta(`INSERT INTO "tasks" ("done","project_id","title") VALUES ($1,$2,$3) RETURNING "tasks"."id"`)).
		WithArgs(false, projectId, title).
		WillReturnRows(mockedRow)
	s.mock.ExpectCommit()

	res, err := s.repository.Create(&title, projectId)

	require.NoError(s.T(), err)
	require.Nil(s.T(), deep.Equal(&models.Task{ID: 1, Title: &title, ProjectID: projectId, Done: done}, res))
}

func (s *Suite) Test_Find() {
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

func (s *Suite) Test_FindByProjectId() {
	var (
		id = "1"
		id2 = "2"
		id3 = "3"
		title = "test-task"
		title2 = "test-task2"
		title3 = "test-task3"
		projectId = int64(5)
		done = false
	)

	mockedRows := sqlmock.NewRows([]string{"id", "title", "done", "project_id"}).
		AddRow(id, title, done, projectId).
		AddRow(id2, title2, done, projectId).
		AddRow(id3, title3, done, projectId)

	s.mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "tasks" WHERE (project_id = $1)`)).
		WithArgs(projectId).
		WillReturnRows(mockedRows)

	res, err := s.repository.FindByProjectId(projectId)

	require.NoError(s.T(), err)
	require.Nil(s.T(), deep.Equal(&models.Task{ID: 1, Title: &title, ProjectID: projectId, Done: done}, res[0]))
	require.Nil(s.T(), deep.Equal(&models.Task{ID: 2, Title: &title2, ProjectID: projectId, Done: done}, res[1]))
	require.Nil(s.T(), deep.Equal(&models.Task{ID: 3, Title: &title3, ProjectID: projectId, Done: done}, res[2]))
}

func (s *Suite) Test_DeleteByTask() {
	title := "test-title"
	task := &models.Task{
		Done:      false,
		ID:        1,
		ProjectID: 2,
		Title:     &title,
	}

	mockedRows := sqlmock.NewRows([]string{"id", "todo", "task_id", "done"}).
		AddRow(2, "test-todo", task.ID, false)

	s.mock.ExpectQuery(regexp.QuoteMeta(
		`SELECT * FROM "todos" WHERE (task_id = $1)`)).
		WithArgs(task.ID).
		WillReturnRows(mockedRows)

	s.mock.ExpectBegin()
	s.mock.ExpectExec(regexp.QuoteMeta(`DELETE FROM "todos" WHERE "todos"."id" = $1`)).
		WithArgs(2).
		WillReturnResult(sqlmock.NewResult(1, 1))
	s.mock.ExpectCommit()

	s.mock.ExpectBegin()
	s.mock.ExpectExec(regexp.QuoteMeta(`DELETE FROM "tasks" WHERE "tasks"."id" = $1`)).
		WithArgs(task.ID).
		WillReturnResult(sqlmock.NewResult(1, 1))
	s.mock.ExpectCommit()

	err := s.repository.DeleteByTask(task)
	require.NoError(s.T(), err)
}

func (s *Suite) Test_Delete() {
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
	s.mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "todos" WHERE (task_id = $1)`)).
		WithArgs(int64(1)).
		WillReturnRows(sqlmock.NewRows([]string{"id", "todo", "done", "task_id"}))
	
	s.mock.ExpectBegin()
	s.mock.ExpectExec(regexp.QuoteMeta(`DELETE FROM "tasks" WHERE "tasks"."id" = $1`)).
		WithArgs(int64(1)).
		WillReturnResult(sqlmock.NewResult(1, 1))
	s.mock.ExpectCommit()

	err := s.repository.Delete(id)
	require.NoError(s.T(), err)
}

