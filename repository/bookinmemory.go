package repository

import (
	"fmt"
	"yourgo/model"
)

type BookInMemoryPerository struct {
	books map[int]model.Book
}

func NewBookInMemoryRepository() *BookInMemoryPerository {
	return &BookInMemoryPerository{
		books: make(map[int]model.Book),
	}
}

func (r BookInMemoryPerository) ByID(ID int) (model.Book, error) {
	book, exists := r.books[ID]
	if !exists {
		return model.Book{}, fmt.Errorf("book with id: %d not found", ID)
	}
	return book, nil
}

func (r BookInMemoryPerository) Add(b model.Book) error {
	if _, exists := r.books[b.ID]; exists {
		return fmt.Errorf("book with id %d already exists", b.ID)
	}
	r.books[b.ID] = b
	return nil
}

func (r BookInMemoryPerository) Delete(ID int) error {
	if _, exists := r.books[ID]; !exists {
		return fmt.Errorf("book with id %d not found", ID)
	}
	delete(r.books, ID)
	return nil
}