package repository

import (
	todo "github.com/istomin10593/reminder_todo"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user todo.User) (int, error)
	GetUser(username, password string) (todo.User, error)
}

type TodoList interface {
	Create(userId int, list todo.TodoList) (int, error)
	GetAll(userId int) ([]todo.TodoList, error)
	GetById(userId, listId int) (todo.TodoList, error)
	Delete(userId int, listid int) error
	Update(userId, listId int, input todo.UpdateListInput) error
}

type TodoItem interface {
	Create(listId int, item todo.TodoItem) (int, error)
}

type Repository struct {
	Authorization
	TodoList
	TodoItem
}

func Newepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		TodoList:      NewTodoListPostgres(db),
	}
}
