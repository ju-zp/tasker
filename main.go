package main

import (
	"flag"
	"github.com/ju-zp/tasker/svc/project"
	"github.com/ju-zp/tasker/svc/task"
	"github.com/ju-zp/tasker/svc/todo"
	"log"
	"os"
	"strconv"

	"github.com/ju-zp/tasker/svc/database"

	"github.com/ju-zp/tasker/svc/handlers/projecthandlers"
	"github.com/ju-zp/tasker/svc/handlers/taskhandlers"
	"github.com/ju-zp/tasker/svc/handlers/todohandlers"

	"github.com/go-openapi/loads"
	"github.com/joho/godotenv"
	"github.com/ju-zp/tasker/svc/handlers/pinghandlers"
	"github.com/ju-zp/tasker/svc/services/tasker/restapi"
	"github.com/ju-zp/tasker/svc/services/tasker/restapi/operations"
)

func init() {
	// loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

func main() {

	db := database.InitDB()

	defer db.Close()

	port := os.Getenv("PORT")
	i, err := strconv.Atoi(port)

	var portFlag = flag.Int(":", i, "Port to run this service on")

	// load embedded swagger file
	swaggerSpec, err := loads.Analyzed(restapi.SwaggerJSON, "")
	if err != nil {
		log.Fatalln(err)
	}

	// create new service API

	api := operations.NewTaskerAPI(swaggerSpec)
	server := restapi.NewServer(api)
	defer server.Shutdown()

	projectCtx := &projecthandlers.Context{
		DB: db,
		Repository: project.CreateRepository(db),
	}
	todoCtx := &todohandlers.Context{
		Repository: todo.CreateRepository(db),
	}
	taskCtx := &taskhandlers.Context{
		DB: db,
		Repository: task.CreateRepository(db),
	}
	// parse flags
	flag.Parse()
	// set the port this service will be run on
	server.Port = *portFlag

	// TODO: Set Handle
	api.GetPingHandler = operations.GetPingHandlerFunc(pinghandlers.GetPing)

	// todo handlers
	api.SetTodoStatusHandler = operations.SetTodoStatusHandlerFunc(todoCtx.SetTodoStatus)
	api.DeleteTodoHandler = operations.DeleteTodoHandlerFunc(todoCtx.DeleteTodo)

	// task handlers
	api.GetTasksHandler = operations.GetTasksHandlerFunc(taskCtx.GetTasks)
	api.CreateTaskHandler = operations.CreateTaskHandlerFunc(taskCtx.CreateTask)
	api.GetTaskTodosHandler = operations.GetTaskTodosHandlerFunc(taskCtx.GetTaskTodos)
	api.CreateTaskTodoHandler = operations.CreateTaskTodoHandlerFunc(taskCtx.CreateTaskTodo)
	api.DeleteTaskHandler = operations.DeleteTaskHandlerFunc(taskCtx.DeleteTask)

	api.GetProjectsHandler = operations.GetProjectsHandlerFunc(projectCtx.GetProjects)
	api.CreateProjectHandler = operations.CreateProjectHandlerFunc(projectCtx.CreateProject)
	api.GetProjectHandler = operations.GetProjectHandlerFunc(projectCtx.GetProject)
	api.DeleteProjectHandler = operations.DeleteProjectHandlerFunc(projectCtx.DeleteProject)

	// serve API
	if err := server.Serve(); err != nil {
		log.Fatalln(err)
	}

}
