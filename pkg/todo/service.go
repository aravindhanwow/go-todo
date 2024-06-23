package todo

type Service interface {
	GetAllTodos() ([]Todo, error)
	GetTodoByID(id int) (Todo, error)
	CreateTodo(todo Todo) (int, error)
	UpdateTodo(todo Todo) error
	DeleteTodoByID(id int) error
}

type service struct {
	repository Repository
}

func NewService(repo Repository) Service {
	return &service{repo}
}

func (s *service) GetAllTodos() ([]Todo, error) {
	return s.repository.GetAll()
}

func (s *service) GetTodoByID(id int) (Todo, error) {
	return s.repository.GetByID(id)
}

func (s *service) CreateTodo(todo Todo) (int, error) {
	return s.repository.Create(todo)
}

func (s *service) UpdateTodo(todo Todo) error {
	return s.repository.Update(todo)
}

func (s *service) DeleteTodoByID(id int) error {
	return s.repository.Delete(id)
}
