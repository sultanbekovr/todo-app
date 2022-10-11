package service

type Authorization struct {
}

type TodoList struct {
}

type TodoItem struct {
}

type Service struct {
	Authorization
	TodoList
	TodoItem
}

func newService() *Service {
	return &Service{}
}
