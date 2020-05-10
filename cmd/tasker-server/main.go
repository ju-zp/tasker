package main

import (
	"flag"
	"log"
	"os"
	"strconv"

	"github.com/ju-zp/tasker/svc/models"

	"github.com/ju-zp/tasker/svc/todohandlers"

	"github.com/go-openapi/loads"
	"github.com/joho/godotenv"
	"github.com/ju-zp/tasker/svc/pinghandlers"
	"github.com/ju-zp/tasker/svc/restapi"
	"github.com/ju-zp/tasker/svc/restapi/operations"
)

func init() {
	// loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

func main() {

	db := models.InitDB()

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

	// parse flags
	flag.Parse()
	// set the port this service will be run on
	server.Port = *portFlag

	// TODO: Set Handle
	api.GetPingHandler = operations.GetPingHandlerFunc(pinghandlers.GetPing)

	api.GetTodosHandler = operations.GetTodosHandlerFunc(todohandlers.GetTodos)
	api.CreateTodoHandler = operations.CreateTodoHandlerFunc(todohandlers.CreateTodo)

	// serve API
	if err := server.Serve(); err != nil {
		log.Fatalln(err)
	}

}
