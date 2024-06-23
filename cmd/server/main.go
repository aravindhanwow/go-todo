package server

import (
	"fmt"

	"net/http"
	"os"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/gorilla/mux"
	"go-todo-app/pkg/db"
	"go-todo-app/pkg/todo"
)

func Run() {
	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stderr)
		logger = level.NewFilter(logger, level.AllowDebug())
		logger = log.With(logger, "ts", log.DefaultTimestampUTC, "caller", log.DefaultCaller)
	}

	dbs, err := db.NewMysqlConnect().NewDBConnection()
	if err != nil {
		fmt.Println(err)
	}

	repository := todo.NewRepository(dbs)
	service := todo.NewService(repository)
	endpoints := todo.NewEndpoints(service)

	r := mux.NewRouter()
	todo.RegisterHTTPHandlers(r, endpoints)

	http.Handle("/", r)
	port := ":8765"
	fmt.Println("Server is running on port", port)
	fmt.Println(http.ListenAndServe(port, nil))
}
