package todo

import (
	"database/sql"
)

type Repository interface {
	GetAll() ([]Todo, error)
	GetByID(id int) (Todo, error)
	Create(todo Todo) (int, error)
	Update(todo Todo) error
	Delete(id int) error
}

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &repository{db}
}

func (r *repository) GetAll() ([]Todo, error) {
	var todos []Todo
	rows, err := r.db.Query("SELECT * FROM todos")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var todo Todo
		if err := rows.Scan(&todo.ID, &todo.Title, &todo.Completed); err != nil {
			return nil, err
		}
		todos = append(todos, todo)
	}

	return todos, nil
}

func (r *repository) GetByID(id int) (Todo, error) {
	var todo Todo
	err := r.db.QueryRow("SELECT * FROM todos WHERE id = ?", id).Scan(&todo.ID, &todo.Title, &todo.Completed)
	if err != nil {
		return todo, err
	}
	return todo, nil
}

func (r *repository) Create(todo Todo) (int, error) {
	result, err := r.db.Exec("INSERT INTO todos (title, completed) VALUES (?, ?)", todo.Title, todo.Completed)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

func (r *repository) Update(todo Todo) error {
	_, err := r.db.Exec("UPDATE todos SET title = ?, completed = ? WHERE id = ?", todo.Title, todo.Completed, todo.ID)
	return err
}

func (r *repository) Delete(id int) error {
	_, err := r.db.Exec("DELETE FROM todos WHERE id = ?", id)
	return err
}
