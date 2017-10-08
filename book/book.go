package book

import "time"

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

type Review struct {
	rating      int
	comment     string
	name        string
	timeCreated time.Time
}

type BookRepository interface {
	findByISBN(isbn string) (Book, error)
}

type ReviewRepostiory interface {
	findByEAN(ean string) ([]Review, error)
}

// END OF VENDOR CODE

// You need to implement the following interface using TDD

type BookService interface {
	// Both functions eturns JSON no matter the response from the service
	getBook(isbn string) string
	getBookWithReviews(isbn string) string
	}

// Your code here!


