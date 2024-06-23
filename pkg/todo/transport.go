package todo

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
)

func RegisterHTTPHandlers(r *mux.Router, endpoints Endpoints) {
	r.Methods("GET").Path("/todos").Handler(httptransport.NewServer(
		endpoints.GetAllTodos,
		decodeGetAllTodosRequest,
		encodeResponse,
	))

	r.Methods("GET").Path("/todos/{id}").Handler(httptransport.NewServer(
		endpoints.GetTodoByID,
		decodeGetTodoByIDRequest,
		encodeResponse,
	))

	r.Methods("POST").Path("/todos").Handler(httptransport.NewServer(
		endpoints.CreateTodo,
		decodeCreateTodoRequest,
		encodeResponse,
	))

	r.Methods("PUT").Path("/todos/{id}").Handler(httptransport.NewServer(
		endpoints.UpdateTodo,
		decodeUpdateTodoRequest,
		encodeResponse,
	))

	r.Methods("DELETE").Path("/todos/{id}").Handler(httptransport.NewServer(
		endpoints.DeleteTodoByID,
		decodeDeleteTodoByIDRequest,
		encodeResponse,
	))
}

func decodeGetAllTodosRequest(_ context.Context, r *http.Request) (interface{}, error) {
	return nil, nil
}

func decodeGetTodoByIDRequest(_ context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		return nil, err
	}
	return getTodoByIDRequest{ID: id}, nil
}

func decodeCreateTodoRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req createTodoRequest
	if err := json.NewDecoder(r.Body).Decode(&req.Todo); err != nil {
		return nil, err
	}
	return req, nil
}

func decodeUpdateTodoRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req updateTodoRequest
	if err := json.NewDecoder(r.Body).Decode(&req.Todo); err != nil {
		return nil, err
	}
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		return nil, err
	}
	req.Todo.ID = id
	return req, nil
}

func decodeDeleteTodoByIDRequest(_ context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		return nil, err
	}
	return deleteTodoByIDRequest{ID: id}, nil
}

func encodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}
