package worker

import (
	"fmt"
	"yourgo/model"
	"yourgo/repository"
)

type repo interface {
	ByID(ID int) (model.Book, error)
	Add(b model.Book) error
}

type Worker struct {
	repo repo
}

func New(repo repo) (*Worker, error) {
	return &Worker{
		repo: repository.NewBookInMemoryRepository(),
	}, nil
}

func (w Worker) CreateBooks() error {
	b1, err := model.NewBook(5, "War and peace", "Lev Tolstoy")
	if err != nil {
		return fmt.Errorf("create new book: %w", err)
	}
	if err := w.repo.Add(b1); err != nil {
		return fmt.Errorf("add book with id %d in repository: %w", b1.ID, err)
	}

	b2, err := model.NewBook(10, "Сrime and punishment", "Phedor Dostaevsky")
	if err != nil {
		return fmt.Errorf("create new book: %w", err)
	}
	if err := w.repo.Add(b2); err != nil {
		return fmt.Errorf("add book with id %d in repository: %w", b2.ID, err)
	}

	return nil
}

func (w Worker) PrintBook(ID int) error {
	b, err := w.repo.ByID(ID)
	if err != nil {
		return fmt.Errorf("get book by id %d from repository: %w", ID, err)
	}
	fmt.Printf("%+v\n", b)
	return nil
}