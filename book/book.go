package book

import (
	"encoding/json"
	"fmt"
)

// START OF VENDOR CODE

// Imagine this code is controlled by a vendor that produces the real
// implementation of the BookRepository and ReviewRepostiory interfaces.
// You must test and create your implementation without being able to
// adapt this section of code.

type Book struct {
	isbn          string
	ean           string
	title         string
	author        string
	price         float32
	yearPublished int
}

func NewBook(isbn, ean, title, author string, price float32, yearPublished int) Book {
	return Book{
		isbn:          isbn,
		ean:           ean,
		title:         title,
		author:        author,
		price:         price,
		yearPublished: yearPublished,
	}
}

type BookRepository interface {
	findByISBN(isbn string) (Book, error)
}

// END OF VENDOR CODE

// IMPLEMENTATION

type BookService struct {
	BookRepository BookRepository
}

type BookJSON struct {
	Isbn          string  `json:"isbn"`
	Ean           string  `json:"ean"`
	Title         string  `json:"title"`
	Author        string  `json:"author"`
	Price         float32 `json:"price"`
	YearPublished int     `json:"year_published"`
}

type ErrorJSON struct {
	Status int    `json:"status"`
	Code   int    `json:"code"`
	Title  string `json:"title"`
	Detail string `json:"detail"`
}

func (bs BookService) getBook(isbn string) string {
	book, err := bs.BookRepository.findByISBN(isbn)

	if err != nil {

		var errorJSON ErrorJSON
		if err.Error() == "NO BOOK" {
			errorJSON = ErrorJSON{
				Status: 404,
				Code:   1,
				Title:  "Book not found",
				Detail: "Could not find book with ISBN: " + isbn,
			}
		} else {
			errorJSON = ErrorJSON{
				Status: 500,
				Code:   2,
				Title:  "Repository unavailable",
				Detail: "Service unavailable, try again later",
			}
		}

		jsonString, _ := json.Marshal(errorJSON)

		return string(jsonString)
	}

	bookJSON := BookJSON{
		Isbn:          book.isbn,
		Ean:           book.ean,
		Title:         book.title,
		Author:        book.author,
		Price:         book.price,
		YearPublished: book.yearPublished,
	}

	fmt.Printf("bookJSON: %v", bookJSON)

	jsonString, _ := json.Marshal(bookJSON)

	return string(jsonString)
}
