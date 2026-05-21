package project

import (
	"errors"
	"fmt"

	"github.com/google/uuid"
)

/*
Вам необходимо разработать простую систему управления проектами,
которая позволяет создавать проекты и задачи, обновлять их статусы и просматривать информацию о них.

Необходимо реализовать отдельный пакет project, где будут следующие возможности:

Создание проекта
	Реализуйте структуру Project, которая будет содержать уникальный идентификатор, название проекта и список задач.
	Напишите функцию New, которая будет создавать новый проект. Проект не может быть создан с пустым названием.

Добавление и обновление задач
	Реализуйте метод AddTask для структуры Project,
	который будет добавлять задачи в проект (если ID задачи повторяется, нужно выбросить ошибку).
	Реализуйте метод UpdateTask, который будет обновлять информацию о задаче в проекте.
	Если в проекте нет переданной задачи, нужно вернуть ошибку.

Фильтрация задач
	Реализуйте метод FilterTasksByStatus, который будет возвращать список задач по заданному статусу (активные или закрытые).

Вывод информации
	Реализуйте метод PrintInfo для структуры Project, который будет выводить информацию о проекте и всех его задачах.
*/

type Project struct {
	ID uuid.UUID
	Name string
	TaskIds []uuid.UUID
	Tasks map[uuid.UUID]Task
}

func New(ID uuid.UUID, name string) (*Project, error) {
	if name == "" {
		return nil, errors.New("emty name")
	}
	return &Project{
		ID: ID,
		Name: name,
		TaskIds: make([]uuid.UUID, 0),
		Tasks: make(map[uuid.UUID]Task),
	}, nil
}

func (p *Project) AddTask(task Task) error {
	if _, ok := p.Tasks[task.ID]; ok {
		return errors.New("task already exists")
	}
	p.TaskIds = append(p.TaskIds, task.ID)
	p.Tasks[task.ID] = task
	return nil
}

func (p *Project) UpdateTask(task Task) error {
	if _, ok := p.Tasks[task.ID]; !ok {
		return errors.New("task not exists")
	}
	p.Tasks[task.ID] = task
	return nil
}

func (p Project) FilterTasksByStatus(status Status) []Task {
	var result []Task
	for _, tid := range p.TaskIds {
		task := p.Tasks[tid]
		if task.Status == status {
			result = append(result, task)
		}
	}
	return result
}

func (p Project) PrintInfo() {
	fmt.Printf("Проект: %s\n", p.Name)
	for _, tid := range p.TaskIds {
		task := p.Tasks[tid]
		fmt.Printf("ID: %s, Задача: %s, Описание: %s, Статус: %s\n", task.ID, task.Title, task.Description, task.Status)
	}
}