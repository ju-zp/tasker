package project

import (
	"database/sql"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/go-test/deep"
	"github.com/jinzhu/gorm"
	"github.com/ju-zp/tasker/svc/models"
	"github.com/stretchr/testify/assert"
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

func (s *Suite) Test_Create() {
	var (
		title = "test-name"
		description = "here"
	)

	mockedRow := sqlmock.NewRows([]string{"id"}).AddRow("1")

	s.mock.ExpectBegin()
	s.mock.ExpectQuery(regexp.QuoteMeta(`INSERT INTO "projects" ("description","title") VALUES ($1,$2) RETURNING "projects"."id"`)).WithArgs(description,title).WillReturnRows(mockedRow)
	s.mock.ExpectCommit()

	err := s.repository.Create(&title, &description)

	require.NoError(s.T(), err)
}

func (s *Suite) Test_Find() {
	var (
		id   = "1"
		title = "test-name"
		description = "test-description"
	)

	s.mock.ExpectQuery(regexp.QuoteMeta(
		`SELECT * FROM "projects" WHERE (id = $1)`)).
		WithArgs(id).
		WillReturnRows(sqlmock.NewRows([]string{"id", "title", "description"}).
			AddRow(id, title, description))

	res, err := s.repository.Find(id)

	require.NoError(s.T(), err)
	require.Nil(s.T(), deep.Equal(&models.Project{ID: 1, Title: &title, Description: &description}, res))
}

func (s *Suite) Test_FindAll() {
	var (
		id1   = "1"
		title1 = "test-name"
		description1 = "test-description"
		id2   = "2"
		title2 = "test-name1"
		description2 = "test-description1"
	)

	s.mock.ExpectQuery(regexp.QuoteMeta(
		`SELECT * FROM "projects"`)).
		WillReturnRows(sqlmock.NewRows([]string{"id", "title", "description"}).
			AddRow(id1, title1, description1).
			AddRow(id2, title2, description2))

	res, err := s.repository.FindAll()

	require.NoError(s.T(), err)
	require.Nil(s.T(), deep.Equal(&models.Project{ID: 1, Title: &title1, Description: &description1}, res[0]))
	require.Nil(s.T(), deep.Equal(&models.Project{ID: 2, Title: &title2, Description: &description2}, res[1]))
}

func (s *Suite) Test_FindByID() {
	var (
		projectId   = "1"
		title = "test-name"
		description = "test-description"
		taskId = "2"
		done = false
		todoId = "1"
		content = "test-todo"
	)

	s.mock.ExpectQuery(regexp.QuoteMeta(
		`SELECT * FROM "projects" WHERE (id = $1)`)).
		WithArgs(projectId).
		WillReturnRows(sqlmock.NewRows([]string{"id", "title", "description"}).
			AddRow(projectId, title, description))

	mockedRows := sqlmock.NewRows([]string{"id", "title", "done", "project_id"}).
		AddRow(taskId, title, done, projectId)

	s.mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "tasks" WHERE (project_id = $1)`)).
		WithArgs(int64(1)).
		WillReturnRows(mockedRows)

	mockedRows = sqlmock.NewRows([]string{"id", "todo", "task_id", "done"}).
		AddRow(todoId, content, taskId, done)

	s.mock.ExpectQuery(regexp.QuoteMeta(
		`SELECT * FROM "todos" WHERE (task_id = $1)`)).
		WithArgs(int64(2)).
		WillReturnRows(mockedRows)

	project, taskTodos, err := s.repository.FindByID(projectId)

	require.NoError(s.T(), err)

	require.Nil(s.T(), deep.Equal(&models.Project{ID: 1, Title: &title, Description: &description}, project))

	assert.Len(s.T(), taskTodos, 1)
	require.Nil(s.T(), deep.Equal(&models.Task{ID: 2, Title: &title, ProjectID: 1, Done: done}, taskTodos[0].Task))

	assert.Len(s.T(), taskTodos[0].Todos, 1)
	require.Nil(s.T(), deep.Equal(&models.Todo{ID: 1, Todo: &content, TaskID: 2, Done: &done}, taskTodos[0].Todos[0]))
}

