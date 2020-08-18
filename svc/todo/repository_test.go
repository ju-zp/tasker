package todo

import (
	"database/sql"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/go-test/deep"
	"github.com/jinzhu/gorm"
	"github.com/ju-zp/tasker/svc/services/tasker/models"
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

func (s *Suite) Test_repository_Create() {
	var (
		content = "test-todo"
		taskId = int64(2)
	)

	mockedRow := sqlmock.NewRows([]string{"id"}).
		AddRow("1")

	s.mock.ExpectBegin()
	s.mock.ExpectQuery(regexp.QuoteMeta(`INSERT INTO "todos" ("done","task_id","todo") VALUES ($1,$2,$3) RETURNING "todos"."id"`)).
		WithArgs(false, taskId, content).
		WillReturnRows(mockedRow)
	s.mock.ExpectCommit()

	err := s.repository.Create(&content, taskId)

	require.NoError(s.T(), err)
}

func (s *Suite) Test_repository_Find() {
	var (
		id   = "1"
		content = "test-todo"
		taskId = int64(3)
		done = false
	)

	mockedRow := sqlmock.NewRows([]string{"id","todo","task_id","done"}).
		AddRow(id, content, taskId, done)

	s.mock.ExpectQuery(regexp.QuoteMeta(
		`SELECT * FROM "todos" WHERE (id = $1)`)).
		WithArgs(id).
		WillReturnRows(mockedRow)

	res, err := s.repository.Find(id)

	require.NoError(s.T(), err)
	require.Nil(s.T(), deep.Equal(&models.Todo{ID: 1, Todo: &content, TaskID: taskId, Done: &done}, res))
}

func (s *Suite) Test_FindByTaskId() {
	var (
		id = "1"
		content = "test-todo"
		taskId = int64(3)
		done = false
		id2 = "2"
		content2 = "test-todo2"
		id3 = "3"
		content3 = "test-todo3"
	)

	mockedRows := sqlmock.NewRows([]string{"id", "todo", "task_id", "done"}).
		AddRow(id, content, taskId, done).
		AddRow(id2, content2, taskId, done).
		AddRow(id3, content3, taskId, done)

	s.mock.ExpectQuery(regexp.QuoteMeta(
		`SELECT * FROM "todos" WHERE (task_id = $1)`)).
		WithArgs(taskId).
		WillReturnRows(mockedRows)

	res, err := s.repository.FindByTaskId(taskId)

	require.NoError(s.T(), err)
	require.Nil(s.T(), deep.Equal(&models.Todo{ID: 1, Todo: &content, TaskID: taskId, Done: &done}, res[0]))
	require.Nil(s.T(), deep.Equal(&models.Todo{ID: 2, Todo: &content2, TaskID: taskId, Done: &done}, res[1]))
	require.Nil(s.T(), deep.Equal(&models.Todo{ID: 3, Todo: &content3, TaskID: taskId, Done: &done}, res[2]))
}

func (s *Suite) Test_UpdateStatus() {
	var (
		id = "1"
		content = "test-todo"
		taskId = int64(3)
		done = false
	)

	mockedRow := sqlmock.NewRows([]string{"id","todo","task_id","done"}).
		AddRow(id, content, taskId, done)

	s.mock.ExpectQuery(regexp.QuoteMeta(
		`SELECT * FROM "todos" WHERE (id = $1)`)).
		WithArgs(id).
		WillReturnRows(mockedRow)

	s.mock.ExpectBegin()

	s.mock.ExpectExec(regexp.QuoteMeta(`UPDATE "todos" SET "done" = $1, "task_id" = $2, "todo" = $3 WHERE "todos"."id" = $4`)).
		WithArgs(true, taskId, content, int64(1)).
		WillReturnResult(sqlmock.NewResult(1, 1))

	s.mock.ExpectCommit()

	err := s.repository.UpdateStatus(id, true)

	require.NoError(s.T(), err)
}

func (s *Suite) Test_DeleteByTodo() {
	content := "test-todo"
	todo := &models.Todo{
		ID: 1,
		Todo: &content,
		TaskID: 2,
	}

	s.mock.ExpectBegin()
	s.mock.ExpectExec(regexp.QuoteMeta(`DELETE FROM "todos" WHERE "todos"."id" = $1`)).
		WithArgs(todo.ID).
		WillReturnResult(sqlmock.NewResult(1, 1))
	s.mock.ExpectCommit()

	s.repository.DeleteByTodo(todo)
}

func (s *Suite) Test_Delete() {
	var (
		id = "1"
		content = "test-todo"
		taskId = int64(3)
		done = false
	)

	mockedRow := sqlmock.NewRows([]string{"id","todo","task_id","done"}).
		AddRow(id, content, taskId, done)

	s.mock.ExpectQuery(regexp.QuoteMeta(
		`SELECT * FROM "todos" WHERE (id = $1)`)).
		WithArgs(id).
		WillReturnRows(mockedRow)

	s.mock.ExpectBegin()

	s.mock.ExpectExec(regexp.QuoteMeta(`DELETE FROM "todos" WHERE "todos"."id" = $1`)).
		WithArgs(int64(1)).
		WillReturnResult(sqlmock.NewResult(1, 1))

	s.mock.ExpectCommit()

	err := s.repository.Delete(id)

	require.NoError(s.T(), err)
}


