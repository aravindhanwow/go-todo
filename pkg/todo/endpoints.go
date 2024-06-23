package todo

import (
	"context"
	"github.com/go-kit/kit/endpoint"
)

type Endpoints struct {
	GetAllTodos    endpoint.Endpoint
	GetTodoByID    endpoint.Endpoint
	CreateTodo     endpoint.Endpoint
	UpdateTodo     endpoint.Endpoint
	DeleteTodoByID endpoint.Endpoint
}

func NewEndpoints(s Service) Endpoints {
	return Endpoints{
		GetAllTodos:    makeGetAllTodosEndpoint(s),
		GetTodoByID:    makeGetTodoByIDEndpoint(s),
		CreateTodo:     makeCreateTodoEndpoint(s),
		UpdateTodo:     makeUpdateTodoEndpoint(s),
		DeleteTodoByID: makeDeleteTodoByIDEndpoint(s),
	}
}

func makeGetAllTodosEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		return s.GetAllTodos()
	}
}

func makeGetTodoByIDEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(getTodoByIDRequest)
		return s.GetTodoByID(req.ID)
	}
}

func makeCreateTodoEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(createTodoRequest)
		id, err := s.CreateTodo(req.Todo)
		return createTodoResponse{ID: id, Err: err}, nil
	}
}

func makeUpdateTodoEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(updateTodoRequest)
		err := s.UpdateTodo(req.Todo)
		return updateTodoResponse{Err: err}, nil
	}
}

func makeDeleteTodoByIDEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(deleteTodoByIDRequest)
		err := s.DeleteTodoByID(req.ID)
		return deleteTodoByIDResponse{Err: err}, nil
	}
}

type getTodoByIDRequest struct {
	ID int
}

type createTodoRequest struct {
	Todo Todo
}

type createTodoResponse struct {
	ID  int   `json:"id"`
	Err error `json:"error,omitempty"`
}

type updateTodoRequest struct {
	Todo Todo
}

type updateTodoResponse struct {
	Err error `json:"error,omitempty"`
}

type deleteTodoByIDRequest struct {
	ID int
}

type deleteTodoByIDResponse struct {
	Err error `json:"error,omitempty"`
}
