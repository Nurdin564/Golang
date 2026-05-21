package project

import (
	"errors"

	"github.com/google/uuid"
)

/*
Создание задач
	Реализуйте структуру Task, которая будет содержать уникальный идентификатор, заголовок, описание и статус задачи (активная или закрытая).
	Напишите функцию NewTask, которая будет создавать новую задачу.
	Задача не может быть создана с пустым заголовком или описанием.
	Задача всегда создается с активным статусом.

Закрытие задач
	Реализуйте метод Close для структуры Task, который будет изменять статус задачи на закрытый,
	если статус уже закрытый, то необходимо вернуть ошибку.
	Обеспечьте возможность обновления описания задачи только для активных задач с помощью метода UpdateDescription.
	Не забудьте проверить то, что проверяли в функции-конструкторе.
*/

type Status string 

const (
	StatusActive = "ACTIVE"
	StatusClosed = "CLOSED"
)

type Task struct {
	ID uuid.UUID
	Title string
	Description string
	Status Status
}

func NewTask(ID uuid.UUID, title, description string) (*Task, error) {
	if title == "" {
		return nil, errors.New("empty title")
	}
	if description == "" {
		return nil, errors.New("empty description")
	}
	return &Task{
		ID: ID,
		Title: title,
		Description: description,
		Status: StatusActive,
	}, nil
}

func (t *Task) UpdateDescription(description string) error {
	if description == "" {
		return errors.New("empty description")
	}
	if t.Status != StatusActive {
		return errors.New("task status is not active")
	}
	t.Description = description
	return nil
}

func (t *Task) Close() error {
	if t.Status == StatusClosed {
		return errors.New("Statis already closed")
	}
	t.Status = StatusClosed
	return nil
}