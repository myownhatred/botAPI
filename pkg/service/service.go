package service

import "github.com/myownhatred/botAPI/pkg/repository"

type Authorization interface {
}

type TodoList interface {
}

type TodoItem interface {
}

type Service struct {
	Authorization
	TodoItem
	TodoList
}

func NewService(rep *repository.Repository) *Service {
	return &Service{}
}
