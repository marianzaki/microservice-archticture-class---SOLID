package service

import (
	"errors"
	"product-catalog-solid/product-catalog-solid/data"
)

// GetBooks returns all books.
func GetBooks() []data.Book {
	return data.Books
}

// GetBookByID returns a book by its ID.
func GetBookByID(id string) (data.Book, error) {
	for _, b := range data.Books {
		if b.ID == id {
			return b, nil
		}
	}
	return data.Book{}, errors.New("book not found")
}

// PostBook adds a new book.
func PostBook(newBook data.Book) {
	data.Books = append(data.Books, newBook)
}
