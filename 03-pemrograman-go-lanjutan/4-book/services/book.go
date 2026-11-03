package services

import (
	"book/models"
	"errors"
	"fmt"
	"slices"
	"strings"
	"uuid"
)

var storage []models.Book = []models.Book{}

type BookService struct{}

func (bs *BookService) GetAll() []models.Book {
	return slices.Clone(storage)
}

func (bs *BookService) FindByID(id string) (models.Book, error) {
	idx, isFound := slices.BinarySearchFunc(storage, models.Book{ID: id}, func(a, b models.Book) int {
		return strings.Compare(a.ID, b.ID)
	})

	if !isFound {
		return models.Book{}, errors.New("book not found")
	}

	return storage[idx], nil
}

func (bs *BookService) Add(input models.BookInput) {
	record := models.Book{
		ID:                uuid.New().String(),
		Title:             input.Title,
		Author:            input.Author,
		Category:          input.Category,
		Publisher:         input.Publisher,
		YearOfPublication: input.YearOfPublication,
	}

	storage = append(storage, record)

	slices.SortFunc(storage, func(a, b models.Book) int {
		return strings.Compare(a.ID, b.ID)
	})

	fmt.Println("Book inserted!")
}

func (bs *BookService) Update(input models.BookInput, id string) error {
	idx, isFound := slices.BinarySearchFunc(storage, models.Book{ID: id}, func(a, b models.Book) int {
		return strings.Compare(a.ID, b.ID)
	})

	if !isFound {
		return errors.New("Book not found")
	}

	storage[idx] = models.Book{
		ID:                storage[idx].ID,
		Title:             input.Title,
		Author:            input.Author,
		Category:          input.Category,
		Publisher:         input.Publisher,
		YearOfPublication: input.YearOfPublication,
	}

	fmt.Println("Book updated!")

	return nil
}

func (bs *BookService) Delete(id string) error {
	_, isFound := slices.BinarySearchFunc(storage, models.Book{ID: id}, func(a, b models.Book) int {
		return strings.Compare(a.ID, b.ID)
	})

	if !isFound {
		return errors.New("Book not found")
	}

	storage = slices.DeleteFunc(storage, func(item models.Book) bool {
		return item.ID == id
	})

	fmt.Println("Book deleted!")

	return nil
}
